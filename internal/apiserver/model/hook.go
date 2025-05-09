package model

import (
	"github.com/MortalSC/FastGO/internal/pkg/rid"
	"gorm.io/gorm"
)

// AfterCreate
func (m *Post) AfterCreate(tx *gorm.DB) error {
	m.PostID = rid.PostID.New(uint64(m.ID))
	return tx.Save(m).Error
}

func (m *User) AfterCreate(tx *gorm.DB) error {
	m.UserID = rid.UserID.New(uint64(m.ID))
	return tx.Save(m).Error
}
