package infra

import "github/four-servings/meonzi/account/domain"

// EntityFromModel create entity from model
func EntityFromModel(model domain.Account) Account {
	return Account{model.ID(), model.Name(), model.LastAccessedAt(), model.CreatedAt(), model.UpdatedAt(), model.DeletedAt()}
}

// ModelFromEntity create model from entity
func ModelFromEntity(entity Account) domain.Account {
	anemic := domain.AnemicAccount{
		ID:             entity.ID,
		Name:           entity.Name,
		LastAccessedAt: entity.LastAccessedAt,
		CreatedAt:      entity.CreatedAt,
		UpdatedAt:      entity.UpdatedAt,
		DeletedAt:      entity.DeletedAt,
	}
	return anemic.ToRichModel()
}
