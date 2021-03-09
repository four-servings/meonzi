package infra

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github/four-servings/meonzi/account/domain"
	"time"
)

type authServiceImpl struct {
	secretKey []byte
	refreshSecretKey []byte
	expires time.Duration
	refreshExpires time.Duration
}

func NewAuthService() domain.AuthService {
	return &authServiceImpl{

		//TODO config
		secretKey:        []byte("access_token_sign"),
		refreshSecretKey: []byte("refresh_token_sign"),
		expires: time.Hour * 7,
		refreshExpires: time.Hour * 24 * 7,
	}
}

func (a *authServiceImpl) GetToken(account domain.Account) (token domain.AccountToken) {
	now := time.Now()
	claims := a.getAccessClaims(account, now)
	refreshClaims := a.getRefreshClaims(account, now)
	token.AccessToken, _ = jwt.NewWithClaims(jwt.SigningMethodHS256, claims).
		SignedString(a.secretKey)
	token.RefreshToken, _ = jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims).
		SignedString(a.refreshSecretKey)
	return
}

func (a *authServiceImpl) GetAccountByAccessToken(token string) (claims *domain.AccountSignInClaims, err error) {
	claims, err = validateToken(token, a.refreshSecretKey)
	//TODO unauthorized error handling
	return
}

func (a *authServiceImpl) GetAccountByRefreshToken(token string) (claims *domain.AccountSignInClaims, err error) {
	claims, err = validateToken(token, a.refreshSecretKey)
	//TODO unauthorized error handling
	return
}

func (a *authServiceImpl) getAccessClaims(account domain.Account, now time.Time) domain.AccountSignInClaims {
	data := account.Data()
	return domain.AccountSignInClaims{
		StandardClaims: jwt.StandardClaims{
			//Audience:  "",
			ExpiresAt: now.Add(a.expires).Unix(),
			//Id:        "",
			IssuedAt:  now.Unix(),
			Issuer:    "self", // TODO refactor, more meaningful than "self"
			//NotBefore: 0
			//Subject:   "",
		},
		AccountId:      data.ID,
	}
}

func (a *authServiceImpl) getRefreshClaims(account domain.Account, now time.Time) domain.AccountSignInClaims {
	data := account.Data()
	return domain.AccountSignInClaims{
		StandardClaims: jwt.StandardClaims{
			//Audience:  "",
			ExpiresAt: now.Add(a.refreshExpires).Unix(),
			//Id:        "",
			IssuedAt:  now.Unix(),
			Issuer:    "self", // TODO refactor, more meaningful than "self"
			//NotBefore: 0
			//Subject:   "",
		},
		AccountId:      data.ID,
	}
}

func validateToken(token string, secretKey []byte) (*domain.AccountSignInClaims, error) {
	tk, err := jwt.ParseWithClaims(token, &domain.AccountSignInClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return secretKey, nil
	})
	if err != nil {
		return nil, err
	}

	return tk.Claims.(*domain.AccountSignInClaims), nil
}
