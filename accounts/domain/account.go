package domain

import "time"

type (
	//Account account interface
	Account interface {
		ID() string
		Name() string
		LastAccessedAt() time.Time
		CreatedAt() time.Time
		UpdatedAt() time.Time
		DeletedAt() *time.Time
	}

	// AccountImplement account model
	AccountImplement struct {
		id             string
		name           string
		lastAccessedAt time.Time
		createdAt      time.Time
		updatedAt      time.Time
		deletedAt      *time.Time
	}

	// AnemicAccount anemic account model
	AnemicAccount struct {
		ID             string
		Name           string
		LastAccessedAt time.Time
		CreatedAt      time.Time
		UpdatedAt      time.Time
		DeletedAt      *time.Time
	}
)

// NewAccount create account instance
func NewAccount(id, name string) Account {
	now := time.Now()
	return &AccountImplement{id, name, now, now, now, nil}
}

// ID account id
func (account *AccountImplement) ID() string {
	return account.id
}

// Name account name
func (account *AccountImplement) Name() string {
	return account.name
}

// LastAccessedAt account last access time
func (account *AccountImplement) LastAccessedAt() time.Time {
	return account.lastAccessedAt
}

// CreatedAt account created time
func (account *AccountImplement) CreatedAt() time.Time {
	return account.createdAt
}

// UpdatedAt account last updated time
func (account *AccountImplement) UpdatedAt() time.Time {
	return account.updatedAt
}

// DeletedAt account deleted time
func (account *AccountImplement) DeletedAt() *time.Time {
	return account.deletedAt
}

// ToRichModel create rich account model from anemic
func (anemic *AnemicAccount) ToRichModel() Account {
	id := anemic.ID
	name := anemic.Name
	lastAccessedAt := anemic.LastAccessedAt
	createdAt := anemic.CreatedAt
	updatedAt := anemic.UpdatedAt
	deletedAt := anemic.DeletedAt
	return &AccountImplement{id, name, lastAccessedAt, createdAt, updatedAt, deletedAt}
}
