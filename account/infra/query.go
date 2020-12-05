package infra

import (
	"github/four-servings/meonzi/account/app/query"
	"log"

	"gorm.io/gorm"
)

type accountQueryImplement struct {
	db *gorm.DB
}

// NewQuery create query instance
func NewQuery(db *gorm.DB) query.AccountQuery {
	return &accountQueryImplement{db}
}

// FindByID find account by id
func (q *accountQueryImplement) FindByID(id string) query.Account {
	entity := Account{}
	if err := q.db.Where(&Account{ID: id}).First(&entity).Error; err != nil {
		panic(err)
	}
	return struct {
		ID   string
		Name string
	}{entity.ID, entity.Name}
}

// FindByName find account by name
func (q *accountQueryImplement) FindByName(name string) query.Accounts {
	entities := []Account{}
	if err := q.db.Where(&Account{Name: name}).Find(&entities).Error; err != nil {
		log.Println(err)
		panic(err)
	}
	result := query.Accounts{}
	for _, entity := range entities {
		result = append(result, struct{ Name string }{Name: entity.Name})
	}
	return result
}
