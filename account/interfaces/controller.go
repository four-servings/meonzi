package interfaces

import (
	"github/four-servings/meonzi/account/application"
)

type RegisterAccountDTO struct {
	Name     string `validate:"required"`
	Token    string
	Provider string
}

type Controller interface {
	RegisterAccount(dto RegisterAccountDTO)
}

type ControllerImpl struct {
	application.CommandBus
}

func (c *ControllerImpl) RegisterAccount(dto RegisterAccountDTO) {
	err := validateDto(dto)
	if err != nil {

	}
	c.CommandBus.Execute(application.RegisterAccountCommand{
		Name: dto.Name,
	})
}

func validateDto(dto RegisterAccountDTO) error {
	return nil
}