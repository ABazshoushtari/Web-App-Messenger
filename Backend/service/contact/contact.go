package contact

import "github.com/ABazshoushtari/Web-App-Messenger/repository"

type Contact struct {
	repos *repository.Repositories
}

func New(repos *repository.Repositories) *Contact {
	return &Contact{
		repos: repos,
	}
}
