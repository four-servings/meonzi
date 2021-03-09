package domain

import "context"

type SocialService interface {
	GetUser(ctx context.Context, provider AuthProvider, token string) (ThirdUser, error)
}