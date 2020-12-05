package infra

import (
	"github/four-servings/meonzi/account/domain"
	"testing"
)

func TestEntityFromModel(t *testing.T) {
	model := domain.NewAccount("id", "name")

	createdAt := model.CreatedAt()
	updatedAt := model.UpdatedAt()
	deletedAt := model.DeletedAt()
	lastAccessedAt := model.LastAccessedAt()

	entity := Account{"id", "name", lastAccessedAt, createdAt, updatedAt, deletedAt}

	result := EntityFromModel(model)

	if result != entity {
		t.Fail()
	}
}

func TestModelFromEntity(t *testing.T) {
	model := domain.NewAccount("id", "name")

	lastAccessedAt := model.LastAccessedAt()
	createdAt := model.CreatedAt()
	updatedAt := model.UpdatedAt()
	deletedAt := model.DeletedAt()

	entity := Account{"id", "name", lastAccessedAt, createdAt, updatedAt, deletedAt}

	result := ModelFromEntity(entity)

	if result.ID() != model.ID() {
		t.Fail()
	}

	if result.Name() != model.Name() {
		t.Fail()
	}

	if result.LastAccessedAt() != model.LastAccessedAt() {
		t.Fail()
	}

	if result.CreatedAt() != model.CreatedAt() {
		t.Fail()
	}

	if result.UpdatedAt() != model.UpdatedAt() {
		t.Fail()
	}

	if result.DeletedAt() != model.DeletedAt() {
		t.Fail()
	}
}
