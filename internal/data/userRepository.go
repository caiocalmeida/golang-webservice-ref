package data

import (
	"github.com/caiocalmeida/go-webservice-ref/internal/domain"
	"github.com/google/uuid"
)

var users = []*domain.User{
	{Id: uuid.UUID{}, Name: "Name1"},
}

func GetUsers() []*domain.User {
	return users
}

func GetUserBy(id uuid.UUID) *domain.User {
	u, _ := findUserBy(id)

	return u
}

func AddUser(u *domain.User) *domain.User {
	users = append(users, u)

	return u
}

func UpdateUser(u *domain.User) *domain.User {
	if user, i := findUserBy(u.Id); user != nil {
		users[i] = u

		return u
	}

	return nil
}

func DeleteUser(id uuid.UUID) bool {
	if user, i := findUserBy(id); user != nil {
		users[i] = users[len(users)-1]
		users = users[:len(users)-1]

		return true
	}

	return false
}

func findUserBy(id uuid.UUID) (*domain.User, int) {
	for i := range users {
		if users[i].Id == id {
			return users[i], i
		}
	}

	return nil, -1
}
