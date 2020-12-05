package domain

import (
	"testing"
	"time"
)

func TestNewAccount(t *testing.T) {
	account := NewAccount("id", "name")
	if account.ID() != "id" || account.Name() != "name" {
		t.Fail()
	}
}

func TestToRichModel(t *testing.T) {
	now := time.Now()
	account := AccountImplement{"id", "name", now, now, now, nil}
	enemic := AnemicAccount{"id", "name", now, now, now, nil}

	result := enemic.ToRichModel()

	if result.ID() != account.id {
		t.Fail()
	}
	if result.Name() != account.name {
		t.Fail()
	}
	if result.LastAccessedAt() != account.lastAccessedAt {
		t.Fail()
	}
	if result.CreatedAt() != account.createdAt {
		t.Fail()
	}
	if result.UpdatedAt() != account.updatedAt {
		t.Fail()
	}
	if result.DeletedAt() != account.deletedAt {
		t.Fail()
	}
}
