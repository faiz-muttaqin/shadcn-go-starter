package modelMerchant

import (
	"time"

	"gorm.io/gorm"
)

type Merchant struct {
	No                       int64     `json:"no" form:"no" gorm:"column:no;primarykey"`     // char 1-15
	MainNo                   int64     `json:"main_no" form:"main_no" gorm:"column:main_no"` // char 16-30
	LoginBlocked             bool      `json:"login_blocked" gorm:"column:login_blocked"`
	RegistrasiName           string    `json:"RegistrasiName" form:"RegistrasiName" gorm:"column:registration_name"`                        // char 31-60
	MainName                 string    `json:"-" form:"-" gorm:"column:main_registration_name"`                                             // char 61-90
	NameContactBusiness      string    `json:"NameContactBusiness" form:"NameContactBusiness" gorm:"column:business_contact_name"`          // char 91-190
	MobileContactBusiness    string    `json:"MobileContactBusiness" form:"MobileContactBusiness" gorm:"column:business_contact_mobile_no"` // char 191-210
	EmailContactBusiness     string    `json:"EmailContactBusiness" form:"EmailContactBusiness" gorm:"column:business_contact_email"`       // char 211-330
	BusinessContactTel       string    `json:"-" form:"-" gorm:"column:business_contact_tel"`                                               // char 331-350
	TechnicalContactName     string    `json:"-" form:"-" gorm:"column:technical_contact_name"`                                             // char 351-450
	TechnicalContactMobileNo string    `json:"-" form:"-" gorm:"column:technical_contact_mobile_no"`                                        // char 451-470
	TechnicalContactEmail    string    `json:"-" form:"-" gorm:"column:technical_contact_email"`                                            // char 471-590
	TechnicalContactTel      string    `json:"-" form:"-" gorm:"column:technical_contact_tel"`                                              // char 591-610
	RegistrationAddress      string    `json:"-" form:"-" gorm:"column:registration_address"`                                               // char 611-770
	AreaCode                 string    `json:"-" form:"-" gorm:"column:area_code"`                                                          // char 771-785
	MerchantCategoryCode     string    `json:"MerchantCategoryCode" form:"MerchantCategoryCode" gorm:"column:category_code"`                // char 786-789
	CreditLevel              string    `json:"-" form:"-" gorm:"column:credit_level"`                                                       // char 790-792
	StartValidDate           string    `json:"StartValidDate" form:"StartValidDate" gorm:"column:valid_start_date"`                         // char 793-800
	StopTrxDate              string    `json:"-" form:"-" gorm:"column:stop_trx_date"`                                                      // char 801-808
	IsMonitored              string    `json:"-" form:"-" gorm:"column:is_monitored"`                                                       // char 809-809
	AmountType               string    `json:"-" form:"-" gorm:"column:amount_type"`                                                        // char 810-819
	IsCreditChecked          string    `json:"-" form:"-" gorm:"column:is_credit_checked"`                                                  // char 820-820
	BusinessStartTime        string    `json:"business_start_time" form:"business_start_time" gorm:"column:business_start_time"`            // char 821-824
	BusinessEndTime          string    `json:"business_end_time" form:"business_end_time" gorm:"column:business_end_time"`                  // char 825-828
	PayCycle                 string    `json:"-" form:"-" gorm:"column:pay_cycle"`                                                          // char 829-848
	Status                   string    `json:"-" form:"-" gorm:"column:status"`                                                             // char 849-849
	NameMerchant             string    `json:"NameMerchant" form:"NameMerchant" gorm:"column:name_en"`                                      // char 850-909
	AddrMerchant             string    `json:"AddrMerchant" form:"AddrMerchant" gorm:"column:addr_en"`                                      // char 910-1109
	CityMerchant             string    `json:"CityMerchant" form:"CityMerchant" gorm:"column:city_en"`                                      // char 1110-1129
	ZipCode                  string    `json:"ZipCode" form:"ZipCode" gorm:"column:zip_code_en"`                                            // char 1130-1135
	Cardtype                 string    `json:"Cardtype" form:"Cardtype" gorm:"column:card_type"`                                            // char 1136-1165
	Secret                   string    `json:"-" gorm:"column:secret"`
	Password                 string    `json:"-" gorm:"column:password"`
	PasswordTemp             string    `json:"-" gorm:"column:password_temp"`
	IsEmailSended            bool      `json:"-" gorm:"column:is_email_sended"`
	IsPasswordRenewed        bool      `json:"-" gorm:"column:is_password_renewed"`
	CreatedAt                time.Time `json:"created_at" gorm:"column:created_at"`
	UpdatedAt                time.Time
	LoginBlockedTime         time.Time `json:"login_blocked_time" gorm:"column:login_blocked_time"`
	LoginUnblockedTime       time.Time `json:"login_unblocked_time" gorm:"column:login_unblocked_time"`
	ProfileImage             string    `json:"-"`
	RoleID                   int       `json:"-" gorm:"default:1"`
	LastLogin                time.Time `json:"last_login"`
	Session                  string    `json:"-"`
	SessionExpired           int64     `json:"-"`
	IPAdress                 string    `json:"ip_address" gorm:"column:ip_address"`
	IP                       string    `json:"-"`
}

func (Merchant) TableName() string {
	return "merchants"
}
func (m *Merchant) BeforeCreate(tx *gorm.DB) (err error) {
	now := time.Now()

	if m.LoginBlockedTime.IsZero() {
		m.LoginBlockedTime = now
	}
	if m.LoginUnblockedTime.IsZero() {
		m.LoginUnblockedTime = now
	}
	m.LastLogin = now
	return
}
func (m *Merchant) BeforeUpdate(tx *gorm.DB) (err error) {
	now := time.Now()

	if m.LoginBlockedTime.IsZero() {
		m.LoginBlockedTime = now
	}
	if m.LoginUnblockedTime.IsZero() {
		m.LoginUnblockedTime = now
	}
	return
}
