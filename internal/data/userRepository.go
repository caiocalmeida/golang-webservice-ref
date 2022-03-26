package data

import (
	"github.com/caiocalmeida/go-webservice-ref/internal/domain"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserRepository interface {
	GetUsers() []*domain.User
	GetUserBy(id uuid.UUID) *domain.User
	AddUser(u *domain.User) *domain.User
	UpdateUser(u *domain.User) *domain.User
	DeleteUser(id uuid.UUID) bool
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db: db}
}

func (ur *userRepository) GetUsers() []*domain.User {
	var users []*domain.User
	ur.db.Find(&users)
	return users
}

func (ur *userRepository) GetUserBy(id uuid.UUID) *domain.User {
	u := &domain.User{}
	result := ur.db.First(u, id)

	if result.RowsAffected > 0 {
		return u
	}

	return nil
}

func (ur *userRepository) AddUser(u *domain.User) *domain.User {
	ur.db.Create(u)
	return u
}

func (ur *userRepository) UpdateUser(u *domain.User) *domain.User {
	ur.db.Save(u)
	return u
}

func (ur *userRepository) DeleteUser(id uuid.UUID) bool {
	result := ur.db.Delete(&domain.User{}, id)

	return result.RowsAffected > 0
}
