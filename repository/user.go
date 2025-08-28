package repository

import (
	"manajemen_gudang_be/entity"

	"gorm.io/gorm"
)

type UserRepository struct {
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{DB: db}
}

func (ur *UserRepository) CreateUser(user *entity.User) error {
	return ur.DB.Create(user).Error
}

func (ur *UserRepository) GetUserByUsername(username string) (*entity.User, error) {
	var user entity.User
	result := ur.DB.Where("username = ?", username).First(&user).Error
	return &user, result
}
