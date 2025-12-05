package modelWA

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type Message struct {
	ID         uint      `gorm:"primaryKey;column:id" json:"id"`
	UserJID    string    `gorm:"column:user_jid;size:50;index" json:"user_jid"`
	MessageID  string    `gorm:"column:message_id;size:100;index" json:"message_id"`
	ChatJID    string    `gorm:"column:chat_jid;size:50;index" json:"chat_jid"`
	SenderJID  string    `gorm:"column:sender_jid;size:50;index" json:"sender_jid"`
	SenderName string    `gorm:"column:sender_name;size:200" json:"sender_name"`
	IsGroup    bool      `gorm:"column:is_group" json:"is_group"`
	GroupName  string    `gorm:"column:group_name;size:500" json:"group_name"`
	IsFromMe   bool      `gorm:"column:is_from_me" json:"is_from_me"`
	Timestamp  time.Time `gorm:"index;column:timestamp" json:"timestamp"`
	Edit       string    `gorm:"column:edit" json:"edit"`

	MessageType string `gorm:"column:message_type;size:100" json:"message_type"` // text, image, video, audio, document, sticker, contact, location, poll, etc
	Text        string `gorm:"type:text;column:text" json:"text"`                // isi teks
	Caption     string `gorm:"type:text;column:caption;size:500" json:"caption"` // caption pada media
	MimeType    string `gorm:"column:mime_type;size:100" json:"mime_type"`       // media MIME type (image/jpeg, video/mp4, etc)
	FileName    string `gorm:"column:file_name;size:100" json:"file_name"`       // nama file kalau ada
	FileSize    int64  `gorm:"column:file_size" json:"file_size"`                // path file lokal atau URL upload
	MediaType   string `gorm:"column:media_type;size:100" json:"media_type"`
	MediaURL    string `gorm:"column:media_url" json:"media_url"`
	MediaPath   string `gorm:"column:media_path" json:"media_path"`

	Latitude     sql.NullFloat64 `gorm:"column:latitude" json:"latitude"`
	Longitude    sql.NullFloat64 `gorm:"column:longitude" json:"longitude"`
	LocationName string          `gorm:"column:location_name;size:100" json:"location_name"`

	ContactName   string `gorm:"column:contact_name;size:100" json:"contact_name"`
	ContactNumber string `gorm:"column:contact_number;size:100" json:"contact_number"`

	PollTitle   string `gorm:"column:poll_title;size:100" json:"poll_title"`
	PollOptions string `gorm:"type:text;column:poll_options;size:500" json:"poll_options"`
	PollVotes   string `gorm:"type:text;column:poll_votes;size:100" json:"poll_votes"`

	EventTitle string       `gorm:"column:event_title;size:100" json:"event_title"`
	EventTime  sql.NullTime `gorm:"column:event_time" json:"event_time"`

	// Reply Info
	ReplyToID    string `gorm:"column:reply_to_id;size:100" json:"reply_to_id"` // MessageID yang di-reply
	ReplyText    string `gorm:"column:reply_text;type:text" json:"reply_text"`  // Isi pesan dari message yang di-reply
	ReplyChatJID string `gorm:"column:reply_chat_jid;size:50" json:"reply_chat_jid"`
	// Reactions (emoji)
	ReactionEmojis string          `gorm:"column:reaction_emojis;type:text" json:"reaction_emojis"` // List emoji (misalnya: 👍,❤️,😂,...)
	Interactive    json.RawMessage `gorm:"type:json" json:"interactive,omitempty"`

	CreatedAt time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at" json:"updated_at"`
}

// TableName returns the name of the table for Message
func (Message) TableName() string {
	return "wa_messages"
}
func (m *Message) BeforeCreate(tx *gorm.DB) (err error) {
	if m.ReactionEmojis == "" {
		m.ReactionEmojis = "{}"
	}
	return nil
}

// UpsertMessage tries to insert a new message or update it if edited
func (m *Message) UpsertMessage(db *gorm.DB) error {
	m.UpdatedAt = time.Now()

	var existingMsg Message
	db.Where("user_jid = ? AND message_id = ? AND chat_jid = ?", m.UserJID, m.MessageID, m.ChatJID).Limit(1).Find(&existingMsg)
	if existingMsg.ID == 0 {
		return db.Create(&m).Error
	}
	// Jika ditemukan dan merupakan pesan yang diedit
	if m.Edit != "" {
		// Update hanya beberapa field yang boleh diubah pada edit
		existingMsg.Text = m.Text
		existingMsg.Caption = m.Caption
		existingMsg.Edit = m.Edit
		existingMsg.ReactionEmojis = m.ReactionEmojis
		existingMsg.UpdatedAt = time.Now()
		return db.Save(&existingMsg).Error
	}
	return nil
}
func (Message) UpdateReactionEmoji(db *gorm.DB, MessageID, jid string, reaction string) error {
	var m Message
	if err := db.Where("message_id = ?", MessageID).First(&m).Error; err != nil {
		return fmt.Errorf("failed to find message: %w", err)
	}

	reactions := map[string]string{}

	if m.ReactionEmojis == "" || m.ReactionEmojis == "{}" {
		// Initialize with an empty map if no reactions exist
		reactions = make(map[string]string)
	} else {
		// Unmarshal existing reactions
		if err := json.Unmarshal([]byte(m.ReactionEmojis), &reactions); err != nil {
			fmt.Println(err)
			return err // Return original if error
		}
	}

	// Update atau hapus reaksi
	if reaction == "" {
		delete(reactions, jid)
	} else {
		reactions[jid] = reaction
	}

	// Marshal ke JSON
	data, err := json.Marshal(reactions)
	if err != nil {
		return fmt.Errorf("failed to marshal reactions: %w", err)
	}

	// Update field & simpan ke DB
	m.ReactionEmojis = string(data)
	m.UpdatedAt = time.Now()

	if err := db.Model(&Message{}).
		Where("id = ?", m.ID).
		UpdateColumns(map[string]interface{}{
			"reaction_emojis": m.ReactionEmojis,
			"updated_at":      m.UpdatedAt,
		}).Error; err != nil {
		return fmt.Errorf("failed to update reaction in database: %w", err)
	}

	return nil
}

// Custom type for sorting messages by CreateAt
type ByCreateAt []Message

// Implement the sort.Interface for ByCreateAt
func (a ByCreateAt) Len() int           { return len(a) }
func (a ByCreateAt) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByCreateAt) Less(i, j int) bool { return a[i].CreatedAt.Before(a[j].CreatedAt) }
