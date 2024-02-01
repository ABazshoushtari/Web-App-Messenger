package contact

import (
	"context"
	"errors"
	"github.com/ABazshoushtari/Web-App-Messenger/domain"
	"github.com/ABazshoushtari/Web-App-Messenger/domain/payloads"
	"github.com/ABazshoushtari/Web-App-Messenger/internal/logger"
	"github.com/ABazshoushtari/Web-App-Messenger/repository"
)

type Contact struct {
	repos *repository.Repositories
}

func New(repos *repository.Repositories) *Contact {
	return &Contact{
		repos: repos,
	}
}
func (c *Contact) ShowContacts(ctx context.Context, userID uint64) (*payloads.ShowContactsResponse, error) {
	user := domain.GetUserDTO(ctx)
	if user == nil {
		return nil, errors.New("invalid jwt user")
	}
	if user.ID != userID {
		return nil, errors.New("input user id does not match with your id")
	}
	contacts, err := c.repos.Contact.GetByUserID(user.ID)
	if err != nil {
		logger.Logger().Errorw("error while getting user contacts", "error", err)
		return nil, errors.New("error while getting user contacts")
	}
	return &payloads.ShowContactsResponse{
		Contacts: contacts,
	}, nil

}

func (c *Contact) AddContact(ctx context.Context, userID uint64, payload payloads.AddContactRequest) (*payloads.AddContactResponse, error) {
	user := domain.GetUserDTO(ctx)
	if user == nil {
		return nil, errors.New("invalid jwt user")
	}
	if user.ID != userID {
		return nil, errors.New("input user id does not match with your id")
	}
	contact := domain.Contact{
		UserID:      userID,
		ContactID:   payload.ContactID,
		ContactName: payload.ContactName,
	}
	if err := c.repos.Contact.Upsert(&contact); err != nil {
		return nil, errors.New("error while")
	}
	return &payloads.AddContactResponse{
		Contact: contact,
	}, nil
}

func (c *Contact) DeleteContact(ctx context.Context, userID uint64, contactID uint64) (*payloads.GenericsSuccessFlagResponse, error) {
	user := domain.GetUserDTO(ctx)
	if user == nil {
		return nil, errors.New("invalid jwt user")
	}
	if user.ID != userID {
		return nil, errors.New("input user id does not match with your id")
	}

	if err := c.repos.Contact.Delete(userID, contactID); err != nil {
		return nil, errors.New("error while deleting contact")
	}
	return &payloads.GenericsSuccessFlagResponse{
		Successful: true,
		Message:    "contact deleted successfully",
	}, nil
}
