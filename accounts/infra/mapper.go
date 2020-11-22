package infra

import "github/four-servings/meonzi/accounts/domain"

// EntityFromModel create entity from model
func EntityFromModel(model domain.Account) Account {
	return Account{model.ID(), model.Name(), model.LastAccessedAt(), model.CreatedAt(), model.UpdatedAt(), model.DeletedAt()}
}

// ModelFromEntity create model from entity
func ModelFromEntity(entity Account) domain.Account {
	anemic := domain.AnemicAccount{entity.ID, entity.Name, entity.LastAccessedAt, entity.CreatedAt, entity.UpdatedAt, entity.DeletedAt}
	return anemic.ToRichModel()
}
