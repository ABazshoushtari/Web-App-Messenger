package payloads

import "github.com/ABazshoushtari/Web-App-Messenger/domain"

type AddContactRequest struct {
	ContactID   uint64 `json:"contact_id"`
	ContactName string `json:"contact_name"`
}

type AddContactResponse struct {
	Contact domain.Contact `json:"contact"`
}
type ShowContactsResponse struct {
	Contacts []domain.Contact `json:"contacts"`
}
