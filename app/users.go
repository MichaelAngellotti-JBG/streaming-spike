package app

import (
	"context"
	"log"

	firebase "firebase.google.com/go"
)

type UserServiceApi interface {
	Register(User) (*User, error)
}

type UserService struct {
	*firebase.App
}

type User struct {
	Email string `json:"email"`
	Id    int64  `json:"id"`
	Name  string `json:"name"`
}

// NewUserService sets up the dependencies we'll need. In our case we need our in memory DB
func NewUserService(ctx context.Context) *UserService {
	app, err := firebase.NewApp(ctx, nil)
	if err != nil {
		log.Fatalf("error initializing app: %v\n", err)
	}

	return &UserService{
		App: app,
	}
}

// Register a new user to our DB
func (service *UserService) Register(u User) (*User, error) {
	return nil, nil
}
