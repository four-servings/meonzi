package kakao

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	log "github.com/sirupsen/logrus"
)

const (
	endPoint = "https://kapi.kakao.com"
)

type User interface {
	GetMe(ctx context.Context, token string) (res UserInfo, err error)
}

type userApi struct{}

func NewUserClient() User {
	return &userApi{}
}

func (u *userApi) GetMe(ctx context.Context, token string) (res UserInfo, err error) {

	req, err := http.NewRequestWithContext(ctx, http.MethodGet,
		fmt.Sprintf("%s%s", endPoint, "/v2/user/me?secure_resource=true"), nil)
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

	switch resp.StatusCode {
	case http.StatusOK:
		err = decoder.Decode(&res)
	default:
		var clientError Error
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
