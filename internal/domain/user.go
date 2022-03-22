package domain

import "github.com/google/uuid"

type User struct {
	Id   uuid.UUID `json:"ID"`
	Name string    `json:"name"`
}
