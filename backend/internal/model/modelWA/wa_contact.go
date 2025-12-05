package modelWA

import (
	"time"

	"gorm.io/gorm"
)

type Contact struct {
	ID        uint           `gorm:"primaryKey" json:"id"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`

	JID          string `gorm:"column:jid;size:128;uniqueIndex" json:"jid"` // e.g., "628xxx@s.whatsapp.net"
	FullName     string `gorm:"column:full_name;size:128" json:"full_name"`
	FirstName    string `gorm:"column:first_name;size:64" json:"first_name"`
	PushName     string `gorm:"column:push_name;size:64" json:"push_name"`
	BusinessName string `gorm:"column:business_name;size:128" json:"business_name"`

	ProfilePictureURL  string `gorm:"column:profile_picture_url;size:512" json:"profile_picture_url"`   // URL to photo
	ProfilePicturePath string `gorm:"column:profile_picture_path;size:512" json:"profile_picture_path"` // local path to photo
}

func (Contact) TableName() string {
	return "wa_contacts"
}
