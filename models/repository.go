package models

import (
	"sync"
)

type Repo struct {
	mu     sync.RWMutex
	users  []User
	orders []Order
}

func NewRepository(users []User, orders []Order) *Repo {
	return &Repo{users: users, orders: orders}
}
