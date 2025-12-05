package modelWA

import (
	"time"
	"unicode/utf8"

	"gorm.io/gorm"
)

type Chat struct {
	ID                int    `gorm:"column:id;primaryKey" json:"id"`
	UserJID           string `gorm:"column:user_jid;size:100;index" json:"user_jid"` // max 100 chars
	JID               string `gorm:"column:jid;size:100;index" json:"jid"`           // Optional: size = max length
	ProfilePictureURL string `gorm:"column:profile_picture_url" json:"profile_picture_url"`
	Name              string `gorm:"column:name;size:100;charset:utf8mb4;collate:utf8mb4_unicode_ci" json:"name"`
	Summary           string `gorm:"column:summary" json:"summary"` // max 100 chars
	IsGroup           bool   `gorm:"column:is_group" json:"is_group"`
	AutoRespondBot    bool   `gorm:"column:auto_respond_bot;default:true" json:"auto_respond_bot"`
	GroupName         string `gorm:"column:group_name;size:500" json:"group_name"`
	// Preview           string         `gorm:"column:preview;size:50" json:"preview"` // max 50 chars
	Preview   string         `gorm:"column:preview;size:50;charset:utf8mb4;collate:utf8mb4_unicode_ci" json:"preview"`
	ChatUrl   string         `gorm:"-" json:"chat_url"` // max 50 chars
	UpdatedAt time.Time      `gorm:"column:updated_at" json:"updated_at"`
	CreatedAt time.Time      `gorm:"column:created_at" json:"created_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}

func (Chat) TableName() string {
	return "wa_chats"
}
func truncateUTF8(s string, max int) string {
	if utf8.RuneCountInString(s) <= max {
		return s
	}
	return string([]rune(s)[:max])
}

// UpsertChat tries to insert new or update existing
// UpsertChat tries to insert new or update existing
func (chat *Chat) UpsertChat(db *gorm.DB) error { // Notice the pointer receiver here
	// Truncate preview and name if needed
	chat.Preview = truncateUTF8(chat.Preview, 50)
	chat.Name = truncateUTF8(chat.Name, 100)
	var existing_chat Chat
	db.Where("user_jid = ? AND jid = ?", chat.UserJID, chat.JID).Limit(1).Find(&existing_chat)
	if existing_chat.ID == 0 {
		// If no existing chat, create a new record in the database
		if err := db.Create(chat).Error; err != nil {
			return err
		}
	} else {
		// Prepare map for changed fields only (not empty and changed)
		updates := make(map[string]interface{})
		if chat.Name != "" && existing_chat.Name != chat.Name {
			updates["name"] = chat.Name
		}
		if chat.Preview != "" && existing_chat.Preview != chat.Preview {
			updates["preview"] = chat.Preview
		}
		if chat.GroupName != "" && existing_chat.GroupName != chat.GroupName {
			updates["group_name"] = chat.GroupName
		}
		if chat.ProfilePictureURL != "" && existing_chat.ProfilePictureURL != chat.ProfilePictureURL {
			updates["profile_picture_url"] = chat.ProfilePictureURL
		}

		if len(updates) > 0 {
			if err := db.Model(&existing_chat).Updates(updates).Error; err != nil {
				return err
			}
			// Re-fetch the updated chat
			if err := db.Where("id = ?", existing_chat.ID).First(&existing_chat).Error; err != nil {
				return err
			}
		}

		// If you want to update the original chat pointer, assign it to the updated `existing_chat`
		*chat = existing_chat // This will update the original `chat` pointer
	}

	return nil
}
