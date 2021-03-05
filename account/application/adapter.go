package application

import (
	"context"
	"github/four-servings/meonzi/account/domain"
)

type (
	// KakaoAdapter kakao service adapter
	KakaoAdapter interface {
		GetUser(ctx context.Context, token string) (ThirdUser, error)
	}

	// GoogleAdapter google service adapter
	GoogleAdapter interface {
		GetUser(ctx context.Context, token string) (ThirdUser, error)
	}

	// ThirdUser third party user data
	ThirdUser interface {
		ID() string
		AuthProvider() domain.AuthProvider
	}
)
