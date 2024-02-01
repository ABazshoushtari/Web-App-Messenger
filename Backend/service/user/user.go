package user

import "github.com/ABazshoushtari/Web-App-Messenger/repository"

type User struct {
	repos *repository.Repositories
}

func New(repos *repository.Repositories) *User {
	return &User{
		repos: repos,
	}
}
