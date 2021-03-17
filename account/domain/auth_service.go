package domain

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
)


type AccountSignInClaims struct {
	jwt.StandardClaims
	AccountId uuid.UUID `json:"accountId"`
	//TODO if you need auth scope like acl, RBAC. so you will define auth scope and write code logic
	//Scopes []string
}

type AccountToken struct {
	AccessToken string
	RefreshToken string
}

type AuthService interface {
	GetToken(account Account) AccountToken
	GetAccountByAccessToken(token string) (claims *AccountSignInClaims, err error)
	GetAccountByRefreshToken(token string) (claims *AccountSignInClaims, err error)
}
