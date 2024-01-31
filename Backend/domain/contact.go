package domain

type Contact struct {
	UserID      uint64 `json:"user_id"`
	ContactID   uint64 `json:"contact_id"`
	ContactName string `json:"contact_name"`
}
