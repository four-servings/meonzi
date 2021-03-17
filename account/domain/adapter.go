package domain

import "context"

type (
	SocialAdapter interface {
		GetUser(ctx context.Context, token string) (ThirdUser, error)
	}

	// KakaoAdapter kakao service adapter
	KakaoAdapter SocialAdapter

	// GoogleAdapter google service adapter
	GoogleAdapter SocialAdapter

	// ThirdUser third party user data
	ThirdUser interface {
		ID() string
		AuthProvider() AuthProvider
	}
)
