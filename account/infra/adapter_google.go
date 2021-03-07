package infra

import (
	"context"
	"encoding/json"
	"fmt"
	log "github.com/sirupsen/logrus"
	"github/four-servings/meonzi/account/app"
	"github/four-servings/meonzi/account/domain"
	"net/http"
	"net/url"
)

type (
	googleAdapter struct{}

	googleUserResponse struct {
		ID int64 `json:"sub"`
	}

	googleUser struct {
		id           string
	}

	googleError struct {
		ErrorDescription string `json:"error_description"`
	}
)

func (e googleError) Error() string {
	return e.ErrorDescription
}

func NewGoogleAdapter() app.GoogleAdapter {
	return &googleAdapter{}
}

func (*googleAdapter) GetUser(ctx context.Context, token string) (user app.ThirdUser, err error)  {
	val := make(url.Values)
	val.Add("id_token", token)

	req, err := http.NewRequestWithContext(ctx, http.MethodGet,
		fmt.Sprintf("%s%s?%s", endPoint, "/oauth2/v3/tokeninfo", val.Encode()), nil)
	if err != nil {
		log.WithError(err).Error("google api request: get token info")
		return
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.WithError(err).Error("google api response: get token info")
		return
	}

	defer resp.Body.Close()
	decoder := json.NewDecoder(resp.Body)

	var model googleUserResponse

	switch resp.StatusCode {
	case http.StatusOK:
		err = decoder.Decode(&model)
		user = googleUser{id: fmt.Sprint(model.ID)}
	default:
		var clientError googleError
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

func (g googleUser) ID() string {
	return g.id
}

func (g googleUser) AuthProvider() domain.AuthProvider {
	return domain.GoogleServiceProviderKey
}
