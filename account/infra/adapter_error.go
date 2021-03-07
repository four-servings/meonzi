package infra

import (
	"context"
	"errors"
	"github/four-servings/meonzi/account/app"
)

type unknownAdapter struct {}

var UnknownAdapter app.SocialAdapter = unknownAdapter{}

func (unknownAdapter) GetUser(_ context.Context, _ string) (app.ThirdUser, error) {
	return nil, errors.New("weird auth provider")
}