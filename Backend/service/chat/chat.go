package chat

import "github.com/ABazshoushtari/Web-App-Messenger/repository"

type Chat struct {
	repos *repository.Repositories
}

func New(repos *repository.Repositories) *Chat {
	return &Chat{
		repos: repos,
	}
}
