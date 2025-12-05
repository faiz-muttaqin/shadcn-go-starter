package modelWA

import (
	"database/sql"
	"time"

	"gorm.io/gorm"
)

// CallEventLog represents a log of a WhatsApp call event.

type CallEventLog struct {
	ID        uint           `gorm:"primaryKey" json:"id"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`

	UserJID        string    `gorm:"column:user_jid;size:100" json:"user_jid"`                 // max 100 chars
	EventType      string    `gorm:"column:event_type;size:32;index" json:"event_type"`        // e.g., CallOffer, CallAccept, CallTerminate, etc.
	CallID         string    `gorm:"column:call_id;size:64;index" json:"call_id"`              // Unique identifier for the call
	FromJID        string    `gorm:"column:from_jid;size:128;index" json:"from_jid"`           // JID of the sender
	CallCreatorJID string    `gorm:"column:call_creator_jid;size:128" json:"call_creator_jid"` // JID of the original call creator
	RemotePlatform string    `gorm:"column:remote_platform;size:64" json:"remote_platform"`    // The platform of the caller's WhatsApp client
	RemoteVersion  string    `gorm:"column:remote_version;size:32" json:"remote_version"`      // Version of the caller's WhatsApp client
	Reason         string    `gorm:"column:reason;size:128" json:"reason"`                     // Reason for termination/rejection
	Media          string    `gorm:"column:media;size:16" json:"media"`                        // audio or video (CallOfferNotice)
	Type           string    `gorm:"column:type;size:16" json:"type"`                          // group or 1on1 (CallOfferNotice)
	RawData        string    `gorm:"column:raw_data" json:"raw_data"`                          // Optionally, XML-encoded or JSON-encoded node
	Timestamp      time.Time `gorm:"column:timestamp" json:"timestamp"`                        // Call event timestamp

	CallOfferAt        sql.NullTime `gorm:"column:call_offer_at" json:"call_offer_at"`
	CallAcceptAt       sql.NullTime `gorm:"column:call_accept_at" json:"call_accept_at"`
	CallPreAcceptAt    sql.NullTime `gorm:"column:call_pre_accept_at" json:"call_pre_accept_at"`
	CallTransportAt    sql.NullTime `gorm:"column:call_transport_at" json:"call_transport_at"`
	CallOfferNoticeAt  sql.NullTime `gorm:"column:call_offer_notice_at" json:"call_offer_notice_at"`
	CallRelayLatencyAt sql.NullTime `gorm:"column:call_relay_latency_at" json:"call_relay_latency_at"`
	CallTerminateAt    sql.NullTime `gorm:"column:call_terminate_at" json:"call_terminate_at"`
	CallRejectAt       sql.NullTime `gorm:"column:call_reject_at" json:"call_reject_at"`
}

// TableName overrides the table name used by GORM.
func (CallEventLog) TableName() string {
	return "wa_call_event_logs"
}
