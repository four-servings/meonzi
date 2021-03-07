package app

import (
	"context"
	"github/four-servings/meonzi/account/domain"
)

type (
	SocialAdapter interface {
		GetUser(ctx context.Context, token string) (ThirdUser, error)
	}

	// KakaoAdapter kakao service adapter
	KakaoAdapter interface {
		SocialAdapter
	}

	// GoogleAdapter google service adapter
	GoogleAdapter interface {
		SocialAdapter
	}

	// ThirdUser third party user data
	ThirdUser interface {
		ID() string
		AuthProvider() domain.AuthProvider
	}
)
