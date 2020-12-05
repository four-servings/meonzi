package infra

import (
	"github/four-servings/meonzi/accounts/domain"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type accountRepositoryImplement struct {
	db *gorm.DB
}

// NewRepository create repository instance
func NewRepository(db *gorm.DB) domain.AccountRepository {
	err := db.AutoMigrate(&Account{})
	if err != nil {
		panic(err)
	}
	return &accountRepositoryImplement{db}
}

// Save insert or update account date
func (r *accountRepositoryImplement) Save(account domain.Account) {
	if err := r.db.Save(EntityFromModel(account)).Error; err != nil {
		panic(err)
	}
}

// FindNewID find new id
func (r *accountRepositoryImplement) FindNewID() string {
	id, err := uuid.NewUUID()
	if err != nil {
		panic(err)
	}

	result := r.db.Where(&Account{ID: id.String()}).First(&Account{})
	if result.Error != nil && result.Error.Error() != "record not found" {
		panic(result.Error)
	}

	if result.RowsAffected != 0 {
		return r.FindNewID()
	}

	return id.String()
}

// FindByID find account by id
func (r *accountRepositoryImplement) FindByID(id string) domain.Account {
	account := Account{}
	result := r.db.Where(&Account{ID: id}).First(&account)
	if result.Error != nil {
		panic(result.Error)
	}
	return ModelFromEntity(account)
}
