package repository

import (
	"github.com/ABazshoushtari/Web-App-Messenger/domain"
	"gorm.io/gorm"
)

type contactRepository struct {
	db *gorm.DB
}

func newContactRepository(db *gorm.DB) *contactRepository {
	return &contactRepository{
		db: db,
	}
}

func (c *contactRepository) Upsert(contact *domain.Contact) error {
	return c.db.Save(&contact).Error
}

func (c *contactRepository) GetByUserID(userID uint64) ([]domain.Contact, error) {
	var contacts []domain.Contact
	err := c.db.Where("user_id = ?", userID).Find(&contacts).Error
	return contacts, err
}

func (c *contactRepository) Delete(userID uint64, contactID uint64) error {
	return c.db.Where("user_id = ? AND contact_id = ?", userID, contactID).Delete(&domain.Contact{}).Error
}
