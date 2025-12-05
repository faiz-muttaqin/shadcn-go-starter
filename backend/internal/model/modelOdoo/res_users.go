package modelOdoo

import (
	"database/sql"
)

type ResUser struct {
	ID                     int           `gorm:"column:id;primaryKey" json:"id"`
	CompanyID              int           `gorm:"column:company_id" json:"company_id"`
	PartnerID              int           `gorm:"column:partner_id" json:"partner_id"`
	Active                 sql.NullBool  `gorm:"column:active" json:"active"`
	CreateDate             sql.NullTime  `gorm:"column:create_date" json:"create_date"`
	Login                  string        `gorm:"column:login" json:"login"`
	Password               string        `gorm:"column:password" json:"password"`
	ActionID               sql.NullInt64 `gorm:"column:action_id" json:"action_id"`
	CreateUID              sql.NullInt64 `gorm:"column:create_uid" json:"create_uid"`
	WriteUID               sql.NullInt64 `gorm:"column:write_uid" json:"write_uid"`
	Signature              string        `gorm:"column:signature" json:"signature"`
	Share                  sql.NullBool  `gorm:"column:share" json:"share"`
	WriteDate              sql.NullTime  `gorm:"column:write_date" json:"write_date"`
	TotpSecret             string        `gorm:"column:totp_secret" json:"totp_secret"`
	NotificationType       string        `gorm:"column:notification_type" json:"notification_type"`
	OdoobotState           string        `gorm:"column:odoobot_state" json:"odoobot_state"`
	OdoobotFailed          sql.NullBool  `gorm:"column:odoobot_failed" json:"odoobot_failed"`
	SaleTeamID             sql.NullInt64 `gorm:"column:sale_team_id" json:"sale_team_id"`
	WebsiteID              sql.NullInt64 `gorm:"column:website_id" json:"website_id"`
	Karma                  sql.NullInt64 `gorm:"column:karma" json:"karma"`
	RankID                 sql.NullInt64 `gorm:"column:rank_id" json:"rank_id"`
	NextRankID             sql.NullInt64 `gorm:"column:next_rank_id" json:"next_rank_id"`
	OAuthProviderID        sql.NullInt64 `gorm:"column:oauth_provider_id" json:"oauth_provider_id"`
	OAuthUID               string        `gorm:"column:oauth_uid" json:"oauth_uid"`
	OAuthAccessToken       string        `gorm:"column:oauth_access_token" json:"oauth_access_token"`
	IIDStoreID             sql.NullInt64 `gorm:"column:iid_store_id" json:"iid_store_id"`
	EmailVerificationToken string        `gorm:"column:email_verification_token" json:"email_verification_token"`
	IsEmailVerified        sql.NullBool  `gorm:"column:is_email_verified" json:"is_email_verified"`
	LoginFailCount         sql.NullInt64 `gorm:"column:login_fail_count" json:"login_fail_count"`
	LoginSuspendedUntil    sql.NullTime  `gorm:"column:login_suspended_until" json:"login_suspended_until"`
	NIKSignup              string        `gorm:"column:nik_signup" json:"nik_signup"`
	NIK                    string        `gorm:"column:nik" json:"nik"`
	Phone                  string        `gorm:"column:phone" json:"phone"`
	IsChangePassword       sql.NullBool  `gorm:"column:is_change_password" json:"is_change_password"`
	LastLogin              sql.NullTime  `gorm:"column:last_login" json:"last_login"`
	LastLoginDaysAgo       sql.NullInt64 `gorm:"column:last_login_days_ago" json:"last_login_days_ago"`
}

func (ResUser) TableName() string {
	return "res_users"
}
