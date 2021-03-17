package infra

import (
	"context"
	"encoding/json"
	"fmt"
	"github/four-servings/meonzi/account/domain"
	"net/http"

	log "github.com/sirupsen/logrus"
)

type (
	kakaoAdapter struct{}

	kakaoUserResponse struct {
		ID int64 `json:"id"`
	}

	kakaoUser struct {
		id           string
	}

	kakaoError struct {
		Code    int64  `json:"code"`
		Message string `json:"msg"`
	}
)

func (e kakaoError) Error() string {
	return fmt.Sprintf("code:%d/message:%s", e.Code, e.Message)
}

const (
	endPoint = "https://kapi.kakao.com"
)

func NewKakaoAdapter() domain.KakaoAdapter {
	return &kakaoAdapter{}
}

func (k *kakaoAdapter) GetUser(ctx context.Context, token string) (user domain.ThirdUser, err error) {
	url := fmt.Sprintf("%s%s", endPoint, "/v2/user/me?secure_resource=true")
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		log.WithError(err).Error("kakao api request: get user info")
		return
	}

	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", token))

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.WithError(err).Error("kakao api response: get user info")
		return
	}

	defer resp.Body.Close()
	decoder := json.NewDecoder(resp.Body)

	var model kakaoUserResponse

	switch resp.StatusCode {
	case http.StatusOK:
		err = decoder.Decode(&model)
		user = kakaoUser{id: fmt.Sprint(model.ID)}
	default:
		var clientError kakaoError
		err = decoder.Decode(&clientError)
		if err == nil {
			err = clientError
			return
		}
	}
	if err != nil {
		log.WithError(err).Error("response body json unmarshal failed")
	}

	return
}

func (k kakaoUser) ID() string {
	return k.id
}

func (k kakaoUser) AuthProvider() domain.AuthProvider {
	return domain.KakaoServiceProviderKey
}
