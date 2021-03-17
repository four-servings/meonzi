package infra

import (
	"context"
	"errors"
	"github/four-servings/meonzi/account/domain"
)

type unknownAdapterImpl struct {}

var unknownAdapter domain.SocialAdapter = unknownAdapterImpl{}

func (unknownAdapterImpl) GetUser(_ context.Context, _ string) (domain.ThirdUser, error) {
	return nil, errors.New("weird auth provider")
}

type socialServiceImpl struct {
	domain.GoogleAdapter
	domain.KakaoAdapter
}

func NewSocialService(google domain.GoogleAdapter, kakao domain.KakaoAdapter) domain.SocialService {
	return &socialServiceImpl{google, kakao}
}

func (s *socialServiceImpl) GetUser(ctx context.Context, provider domain.AuthProvider, token string) (domain.ThirdUser, error) {
	return s.getSocialAdapter(provider).GetUser(ctx, token)
}

func (s *socialServiceImpl) getSocialAdapter(provider domain.AuthProvider) domain.SocialAdapter {
	switch provider {
	case domain.GoogleServiceProviderKey:
		return s.GoogleAdapter
	case domain.KakaoServiceProviderKey:
		return s.KakaoAdapter
	}

	return unknownAdapter
}