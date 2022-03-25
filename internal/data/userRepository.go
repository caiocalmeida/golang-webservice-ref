package data

import (
	"github.com/caiocalmeida/go-webservice-ref/internal/domain"
	"github.com/google/uuid"
)

func GetUsers() []*domain.User {
	var users []*domain.User
	db.Find(&users)
	return users
}

func GetUserBy(id uuid.UUID) *domain.User {
	u := &domain.User{}
	result := db.First(u, id)

	if result.RowsAffected > 0 {
		return u
	}

	return nil
}

func AddUser(u *domain.User) *domain.User {
	db.Create(u)
	return u
}

func UpdateUser(u *domain.User) *domain.User {
	db.Save(u)
	return u
}

func DeleteUser(id uuid.UUID) bool {
	result := db.Delete(&domain.User{}, id)

	return result.RowsAffected > 0
}
