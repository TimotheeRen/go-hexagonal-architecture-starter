package ports

import (
	"16/internal/core/domain"

	"gorm.io/gorm"
)

type UserRepository interface {
	CreateUser(db *gorm.DB, name string) error
	GetUser(db *gorm.DB, id uint) (*domain.User, error)
}
