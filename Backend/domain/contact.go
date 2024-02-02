package domain

type Contact struct {
	UserID      uint64 `gorm:"primaryKey" json:"user_id"`
	ContactID   uint64 `gorm:"primaryKey" json:"contact_id"`
	ContactName string `json:"contact_name"`
}
