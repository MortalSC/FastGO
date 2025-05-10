package model

import (
	"github.com/MortalSC/FastGO/internal/commonpkg/authn"
	"github.com/MortalSC/FastGO/internal/pkg/rid"
	"gorm.io/gorm"
)

// == Post ==

// AfterCreate
func (m *Post) AfterCreate(tx *gorm.DB) error {
	m.PostID = rid.PostID.New(uint64(m.ID))
	return tx.Save(m).Error
}

// == User ==
// BeforeCreate
func (m *User) BeforeCreate(tx *gorm.DB) error {
	var err error
	m.Password, err = authn.Encrypt(m.Password)
	if err != nil {
		return err
	}
	return nil
}

// AfterCreate
func (m *User) AfterCreate(tx *gorm.DB) error {
	m.UserID = rid.UserID.New(uint64(m.ID))
	return tx.Save(m).Error
}
