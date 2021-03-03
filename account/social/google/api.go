package google

import (
	"context"
	"encoding/json"
	"fmt"
	log "github.com/sirupsen/logrus"
	"net/http"
	"net/url"
)

const (
	endPoint = "https://www.googleapis.com"

)

type Auth interface {
	GetByIdToken(ctx context.Context, idToken string) (res TokenInfo, err error)
}

type authApi struct { }

func NewAuthClient() Auth {
	return &authApi{}
}

func (*authApi) GetByIdToken(ctx context.Context, idToken string) (res TokenInfo, err error) {
	val := make(url.Values)
	val.Add("id_token", idToken)

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