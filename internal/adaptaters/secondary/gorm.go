package secondary

import (
	"16/internal/core/domain"
	"16/internal/core/ports"

	"gorm.io/gorm"
)

type UserRepositoryImpl struct {}

func (u *UserRepositoryImpl) GetUser(db *gorm.DB, id uint) (*domain.User, error) {
	user := &domain.User{}
	result := db.First(user, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return user, nil
}

func (u *UserRepositoryImpl) CreateUser(db *gorm.DB, name string) error {
	user := &domain.User{
		Name: name,
	}
	result := db.Create(user)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

var _ ports.UserRepository = (*UserRepositoryImpl)(nil)
