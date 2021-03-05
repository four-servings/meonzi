package domain

import (
	"context"
	"github/four-servings/meonzi/util/pointer"
	"time"

	"github.com/google/uuid"
)

// NewAccount create new account object
func NewAccount(options NewAccountOptions) Account {
	return &accountImpl{
		id:           options.ID,
		name:         options.Name,
		authProvider: options.AuthProvider,
		socialID:     options.SocialID,
		createdAt:    time.Now(),
		updatedAt:    time.Now(),
		deletedAt:    nil,
	}
}

// ReconstituteAccount reconstitede account object
func ReconstituteAccount(options ReconstituteAccountOptions) Account {
	return &accountImpl{
		id:           options.ID,
		authProvider: options.AuthProvider,
		socialID:     options.SocialID,
		name:         options.Name,
		createdAt:    options.CreatedAt,
		updatedAt:    options.UpdatedAt,
		deletedAt:    options.DeletedAt,
	}
}

// NewAccountOptions account option for create new account
type NewAccountOptions struct {
	ID           string
	Name         string
	AuthProvider AuthProvider
	SocialID     string
}

// ReconstituteAccountOptions reconstitute option for reconstitute account
type ReconstituteAccountOptions struct {
	ID           string
	Name         string
	AuthProvider AuthProvider
	SocialID     string
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    *time.Time
}

// Account account domain object interface
type Account interface {
	Events() []interface{}
	Apply(interface{})
	Deregister()
}

// Deregister deregister account
func (a *accountImpl) Deregister() {
	a.deletedAt = pointer.Time(time.Now())
}

// Events get applied account events
func (a *accountImpl) Events() (events []interface{}) {
	events = append(events, a.events...)
	return
}

// Apply apply event to account
func (a *accountImpl) Apply(event interface{}) {
	a.events = append(a.events, event)
}

const (
	// KakaoServiceProviderKey key for AuthProvider value
	KakaoServiceProviderKey = AuthProvider("KAKAO")

	// GoogleServiceProviderKey key for AuthProvider value
	GoogleServiceProviderKey = AuthProvider("GOOGLE")
)

type accountImpl struct {
	id           string
	authProvider AuthProvider
	socialID     string
	name         string
	createdAt    time.Time
	updatedAt    time.Time
	deletedAt    *time.Time
	events       []interface{}
}

// AuthProvider third party service provider
type AuthProvider string

// AccountRepository account repository
type AccountRepository interface {
	Save(ctx context.Context, account Account) error
	FindByID(ctx context.Context, id uuid.UUID) (Account, error)
	FindNewID() (string, error)
	FindByProviderAndSocialID(ctx context.Context, provider AuthProvider, socialID string) (Account, error)
}
