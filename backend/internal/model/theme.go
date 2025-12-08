package model

import (
	"database/sql/driver"
	"encoding/json"
	"time"

	"gorm.io/gorm"
)

// ThemeStyles represents the light and dark theme configuration
type ThemeStyles struct {
	Light ThemeStyleProps `json:"light"`
	Dark  ThemeStyleProps `json:"dark"`
}

// ThemeStyleProps represents individual theme properties
type ThemeStyleProps struct {
	Background               string `json:"background"`
	Foreground               string `json:"foreground"`
	Card                     string `json:"card"`
	CardForeground           string `json:"card-foreground"`
	Popover                  string `json:"popover"`
	PopoverForeground        string `json:"popover-foreground"`
	Primary                  string `json:"primary"`
	PrimaryForeground        string `json:"primary-foreground"`
	Secondary                string `json:"secondary"`
	SecondaryForeground      string `json:"secondary-foreground"`
	Muted                    string `json:"muted"`
	MutedForeground          string `json:"muted-foreground"`
	Accent                   string `json:"accent"`
	AccentForeground         string `json:"accent-foreground"`
	Destructive              string `json:"destructive"`
	DestructiveForeground    string `json:"destructive-foreground"`
	Border                   string `json:"border"`
	Input                    string `json:"input"`
	Ring                     string `json:"ring"`
	Chart1                   string `json:"chart-1"`
	Chart2                   string `json:"chart-2"`
	Chart3                   string `json:"chart-3"`
	Chart4                   string `json:"chart-4"`
	Chart5                   string `json:"chart-5"`
	Sidebar                  string `json:"sidebar"`
	SidebarForeground        string `json:"sidebar-foreground"`
	SidebarPrimary           string `json:"sidebar-primary"`
	SidebarPrimaryForeground string `json:"sidebar-primary-foreground"`
	SidebarAccent            string `json:"sidebar-accent"`
	SidebarAccentForeground  string `json:"sidebar-accent-foreground"`
	SidebarBorder            string `json:"sidebar-border"`
	SidebarRing              string `json:"sidebar-ring"`
	FontSans                 string `json:"font-sans"`
	FontSerif                string `json:"font-serif"`
	FontMono                 string `json:"font-mono"`
	Radius                   string `json:"radius"`
	ShadowColor              string `json:"shadow-color"`
	ShadowOpacity            string `json:"shadow-opacity"`
	ShadowBlur               string `json:"shadow-blur"`
	ShadowSpread             string `json:"shadow-spread"`
	ShadowOffsetX            string `json:"shadow-offset-x"`
	ShadowOffsetY            string `json:"shadow-offset-y"`
	LetterSpacing            string `json:"letter-spacing"`
	Spacing                  string `json:"spacing,omitempty"`
}

// Scan implements the sql.Scanner interface for ThemeStyles
func (ts *ThemeStyles) Scan(value interface{}) error {
	if value == nil {
		return nil
	}
	bytes, ok := value.([]byte)
	if !ok {
		return nil
	}
	return json.Unmarshal(bytes, ts)
}

// Value implements the driver.Valuer interface for ThemeStyles
func (ts ThemeStyles) Value() (driver.Value, error) {
	return json.Marshal(ts)
}

// Theme represents a user's theme configuration
type Theme struct {
	ID        string         `gorm:"primaryKey;column:id;size:100" json:"id"`
	UserID    string         `gorm:"column:user_id;size:200;index" json:"userId"`
	Name      string         `gorm:"column:name;size:100" json:"name"`
	Styles    ThemeStyles    `gorm:"column:styles;type:jsonb" json:"styles"`
	CreatedAt time.Time      `gorm:"column:created_at" json:"createdAt"`
	UpdatedAt time.Time      `gorm:"column:updated_at" json:"updatedAt"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}

// TableName overrides the table name
func (Theme) TableName() string {
	return "themes"
}

// BeforeCreate hook to generate ID if not provided
func (t *Theme) BeforeCreate(tx *gorm.DB) error {
	if t.ID == "" {
		// Generate a simple ID (you can use a better ID generator like cuid or uuid)
		t.ID = generateThemeID()
	}
	return nil
}

// Helper function to generate theme ID (simple implementation)
func generateThemeID() string {
	// In production, use a proper ID generator like github.com/rs/xid or cuid
	return "theme_" + time.Now().Format("20060102150405")
}
