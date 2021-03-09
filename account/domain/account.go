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
		AccountData: AccountData{
			ID:           options.ID,
			Name:         options.Name,
			AuthProvider: options.AuthProvider,
			SocialID:     options.SocialID,
			CreatedAt:    time.Now(),
			UpdatedAt:    time.Now(),
			DeletedAt:    nil,
		},
	}
}

// ReconstituteAccount reconstitute account object
func ReconstituteAccount(options ReconstituteAccountOptions) Account {
	return &accountImpl{
		AccountData: AccountData{
			ID:           options.ID,
			Name:         options.Name,
			AuthProvider: options.AuthProvider,
			SocialID:     options.SocialID,
			CreatedAt:    options.CreatedAt,
			UpdatedAt:    options.UpdatedAt,
			DeletedAt:    options.DeletedAt,
		},
	}
}

// NewAccountOptions account option for create new account
type NewAccountOptions struct {
	ID           uuid.UUID
	Name         string
	AuthProvider AuthProvider
	SocialID     string
}

// ReconstituteAccountOptions reconstitute option for reconstitute account
type ReconstituteAccountOptions struct {
	ID           uuid.UUID
	Name         string
	AuthProvider AuthProvider
	SocialID     string
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    *time.Time
}

type AccountData struct {
	ID           uuid.UUID
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
	Data() AccountData
}

// Deregister deregister account
func (a *accountImpl) Deregister() {
	a.DeletedAt = pointer.Time(time.Now())
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

func (a *accountImpl) Data() AccountData {
	data := a.AccountData

	//deep copy
	if data.DeletedAt != nil {
		data.DeletedAt = pointer.Time(*data.DeletedAt)
	}

	return data
}

const (
	// KakaoServiceProviderKey key for AuthProvider value
	KakaoServiceProviderKey = AuthProvider("KAKAO")

	// GoogleServiceProviderKey key for AuthProvider value
	GoogleServiceProviderKey = AuthProvider("GOOGLE")
)

type accountImpl struct {
	AccountData
	events       []interface{}
}

// AuthProvider third party service provider
type AuthProvider string

// AccountRepository account repository
type AccountRepository interface {
	Save(ctx context.Context, account Account) (Account, error)
	FindByID(ctx context.Context, id uuid.UUID) (Account, error)
	FindNewID(ctx context.Context) (uuid.UUID, error)
	FindByProviderAndSocialID(ctx context.Context, provider AuthProvider, socialID string) (Account, error)
}
