package infra

import "time"

// Account account entity
type Account struct {
	ID             string     `gorm:"primary_key"`
	Name           string     `gorm:"not null"`
	LastAccessedAt time.Time  `gorm:"not null"`
	CreatedAt      time.Time  `gorm:"not null"`
	UpdatedAt      time.Time  `gorm:"not null"`
	DeletedAt      *time.Time `sql:"index"`
}
