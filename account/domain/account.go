package domain

import (
	"github.com/google/uuid"
	"github/four-servings/meonzi/ent/schema"
	"time"
)

type AccountAnemic struct {

}

type AccountOptions struct {
	SocialType schema.SocialType
	SocialId string
	Name string
}

type Account interface {

}

type accountImpl struct {
	id uuid.UUID
	socialType schema.SocialType
	socialId string
	name string
	lastAccessedAt time.Time
	createAt time.Time
	updateAt time.Time
	deleteAt *time.Time
}


func NewAccount(options AccountOptions) Account {
	return &accountImpl{
		 socialType: options.SocialType,
		 socialId: options.SocialId,
		 name: options.Name,
	}
}


