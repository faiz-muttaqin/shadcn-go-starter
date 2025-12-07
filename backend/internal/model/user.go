package model

import (
	"html/template"
	"reflect"
	"slices"
	"strings"
	"time"

	"github.com/faiz-muttaqin/shadcn-admin-go-starter/backend/pkg/types"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
	"gorm.io/gorm"
)

type User struct {
	ID                 uint           `gorm:"primaryKey;column:id" json:"id" ui:"sortable"`
	ExternalID         string         `gorm:"column:external_id;size:200;unique" json:"external_id"`
	VerificationStatus string         `gorm:"column:verification_status;size:50" json:"verification_status"`
	Avatar             types.Avatar   `gorm:"column:avatar;size:255" json:"avatar" ui:"visible;visibility;editable"`
	Email              types.Email    `gorm:"column:email;size:100;unique" json:"email" ui:"creatable;visible;visibility;editable;filterable;sortable"`
	Username           string         `gorm:"column:username;size:50" json:"username" ui:"creatable;visible;visibility;editable;filterable;sortable"`
	FirstName          string         `gorm:"column:first_name;size:50" json:"first_name" ui:"creatable;visible;visibility;editable;filterable;sortable"`
	LastName           string         `gorm:"column:last_name;size:50" json:"last_name" ui:"creatable;visible;visibility;editable;filterable;sortable"`
	PhoneNumber        types.Phone    `gorm:"column:phone_number;size:20" json:"phone_number" ui:"creatable;visible;visibility;editable;filterable;sortable"`
	Password           types.Password `gorm:"column:password;size:100" json:"-"`
	Status             types.Badge    `gorm:"column:status" json:"status" ui:"visible;visibility;editable;filterable;sortable;selection:/options?data=status"`
	Session            string         `gorm:"column:session;size:120" json:"session"`
	LastLogin          time.Time      `gorm:"column:last_login" json:"last_login" ui:"visible;visibility;filterable;sortable"`
	RoleID             uint           `gorm:"column:role_id;index" json:"role_id"`
	UserRole           UserRole       `gorm:"foreignKey:RoleID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL" json:"user_role" `
	Role               types.HTML     `gorm:"-" json:"role" ui:"visible;visibility;editable;filterable;sortable;selection:/options?data=role"`

	CreatedAt time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at" json:"updated_at"`
}

// Override nama tabel jika ingin lebih eksplisit
func (User) TableName() string {
	return "users"
}

const (
	StatusInactive  = "inactive"
	StatusActive    = "active"
	StatusBanned    = "banned"
	StatusSuspended = "suspended"
)

func (u *User) StatusInactive() {
	u.Status = StatusInactive
}
func (u *User) StatusActive() {
	u.Status = StatusActive
}
func (u *User) StatusBanned() {
	u.Status = StatusBanned
}
func (u *User) StatusSuspended() {
	u.Status = StatusSuspended
}
func (m *User) BeforeCreate(tx *gorm.DB) (err error) {
	if m.RoleID == 0 {
		m.RoleID = 2 // Set ke role default
	}
	if m.Status == "" {
		m.Status = StatusInactive
	}
	if m.LastLogin.IsZero() {
		m.LastLogin = time.Now()
	}
	return nil
}
func (u *User) AfterFind(tx *gorm.DB) error {
	u.Role = types.HTML(`<i class="` + u.UserRole.Icon + `"></i> ` + u.UserRole.Name)
	return nil
}
func (m User) IgnoredColumn() []string {
	return []string{
		"username",
		"external_id",
		"verification_status",
		"updated_at",
		"password",
		"session",
		"deleted_at",
	}
}
func (m User) TableSettings(url string) map[string]any {
	if url == "" {
		url = "/users"
	}
	type col struct {
		Name       string       `json:"name"`
		Data       string       `json:"data"`
		Type       string       `json:"type"`
		Visible    bool         `json:"visible"`
		Visibility bool         `json:"visibility"` // toggle visibility / visibility control
		Sortable   bool         `json:"sortable"`
		Filterable bool         `json:"filterable"`
		Creatable  bool         `json:"creatable"`
		Editable   bool         `json:"editable"`
		Selection  template.URL `json:"selection"`
	}
	var tableHeaders []col
	// Use reflection to get the type of the struct
	t := reflect.TypeOf(m)
	// Loop through the fields of the struct
	for i := 0; i < t.NumField(); i++ {

		field := t.Field(i)
		// Get the variable name
		// varName := field.Name
		// varName = ""
		// Get the data type
		// Get the JSON key
		jsonKey := field.Tag.Get("json")
		if jsonKey == "" || jsonKey == "-" {
			continue
		}
		jsonKey = strings.Split(jsonKey, ",")[0]
		if slices.Contains(m.IgnoredColumn(), jsonKey) {
			continue
		}
		// Parse UI tag
		ui := types.ParseUIOptions(field.Tag.Get("ui"))

		dataType := string(types.DetectFieldType(field.Type))
		if dataType == "" {
			dataType = field.Type.String()
			switch dataType {
			case "model.UserRole":
				dataType = "object"
			default:
				continue
			}
		}
		title := strings.ToLower(strings.ReplaceAll(jsonKey, "_", " "))
		tableHeaders = append(tableHeaders,
			col{
				Name:       cases.Title(language.English).String(title),
				Data:       jsonKey,
				Type:       dataType,
				Visible:    ui.Visible,
				Visibility: ui.Visibility,
				Sortable:   ui.Sortable,
				Filterable: ui.Filterable,
				Creatable:  ui.Creatable,
				Editable:   ui.Editable,
				Selection:  template.URL(ui.Selection),
			},
		)
	}
	table_data := map[string]any{
		"name":         m.TableName(),
		"title":        "Users Table",
		"url":          url,
		"sort":         "id desc",
		"row":          10,
		"row_opt":      []int{10, 20, 30, 40, 50, 100, 200, 500, 1000},
		"creatable":    true,
		"checkable":    true,
		"editable":     true,
		"deletable":    true,
		"passwordable": false,
		"column":       tableHeaders,
	}
	return table_data
}
