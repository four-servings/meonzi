package interfaces

import (
	"context"
	"errors"
	"github.com/go-playground/validator/v10"
	log "github.com/sirupsen/logrus"
	"github/four-servings/meonzi/account/app"
	"github/four-servings/meonzi/account/domain"
)

type RegisterAccountDTO struct {
	Name     string `validate:"required,min=2,max=8"`
	Token    string `validate:"required"`
	Provider string `validate:"required,eq=KAKAO|eq=GOOGLE"`
}

type SignInAccountDTO struct {
	Token string `validate:"required"`
	Provider string `validate:"required,eq=KAKAO|eq=GOOGLE"`
}

type Controller interface {
	RegisterAccount(dto RegisterAccountDTO) error
	SignIn(dto SignInAccountDTO) (domain.AccountToken, error)
}

type controllerImpl struct {
	app.CommandBus
	*validator.Validate
	domain.AccountRepository
	domain.AuthService
	domain.SocialService
}

func NewAccountController(bus app.CommandBus, validator *validator.Validate,
	repository domain.AccountRepository, authService domain.AuthService, socialService domain.SocialService) Controller {
	return &controllerImpl{bus, validator, repository, authService, socialService}
}

func (c *controllerImpl) RegisterAccount(dto RegisterAccountDTO) (err error) {
	err = c.Validate.Struct(dto)
	if err != nil {
		//TODO bad request
		return
	}
	return c.CommandBus.Execute(app.RegisterAccountCommand{
		Token: dto.Token,
		Name: dto.Name,
		Provider: domain.AuthProvider(dto.Provider),
	})
}


func (c *controllerImpl) SignIn(dto SignInAccountDTO) (token domain.AccountToken, err error) {
	err = c.Validate.Struct(dto)
	if err != nil {
		//TODO bad request
		return
	}

	 user, err := c.SocialService.GetUser(context.Background(), domain.AuthProvider(dto.Provider), dto.Token)
	 if err != nil {
		 log.WithError(err).Errorf("Can not fetch %s user", dto.Provider)
		 return
	 }

	account, _ := c.AccountRepository.FindByProviderAndSocialID(context.Background(), user.AuthProvider(), user.ID())
	if account == nil {
		err = errors.New("not found account")
		log.Error("Can not found account")
		return
	}

	token = c.AuthService.GetToken(account)
	return
}