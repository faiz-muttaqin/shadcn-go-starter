package modelOdoo

import (
	"database/sql"
	"time"
)

type SaleOrder struct {
	ID                     int             `gorm:"column:id;primaryKey" json:"id"`
	CampaignID             sql.NullInt64   `gorm:"column:campaign_id" json:"campaign_id"`
	SourceID               sql.NullInt64   `gorm:"column:source_id" json:"source_id"`
	MediumID               sql.NullInt64   `gorm:"column:medium_id" json:"medium_id"`
	CompanyID              int             `gorm:"column:company_id" json:"company_id"`
	PartnerID              int             `gorm:"column:partner_id" json:"partner_id"`
	JournalID              sql.NullInt64   `gorm:"column:journal_id" json:"journal_id"`
	PartnerInvoiceID       int             `gorm:"column:partner_invoice_id" json:"partner_invoice_id"`
	PartnerShippingID      int             `gorm:"column:partner_shipping_id" json:"partner_shipping_id"`
	FiscalPositionID       sql.NullInt64   `gorm:"column:fiscal_position_id" json:"fiscal_position_id"`
	PaymentTermID          sql.NullInt64   `gorm:"column:payment_term_id" json:"payment_term_id"`
	PricelistID            sql.NullInt64   `gorm:"column:pricelist_id" json:"pricelist_id"`
	CurrencyID             sql.NullInt64   `gorm:"column:currency_id" json:"currency_id"`
	UserID                 sql.NullInt64   `gorm:"column:user_id" json:"user_id"`
	TeamID                 sql.NullInt64   `gorm:"column:team_id" json:"team_id"`
	AnalyticAccountID      sql.NullInt64   `gorm:"column:analytic_account_id" json:"analytic_account_id"`
	CreateUID              sql.NullInt64   `gorm:"column:create_uid" json:"create_uid"`
	WriteUID               sql.NullInt64   `gorm:"column:write_uid" json:"write_uid"`
	AccessToken            string          `gorm:"column:access_token" json:"access_token"`
	Name                   string          `gorm:"column:name" json:"name"`
	State                  string          `gorm:"column:state" json:"state"`
	ClientOrderRef         string          `gorm:"column:client_order_ref" json:"client_order_ref"`
	Origin                 string          `gorm:"column:origin" json:"origin"`
	Reference              string          `gorm:"column:reference" json:"reference"`
	SignedBy               string          `gorm:"column:signed_by" json:"signed_by"`
	InvoiceStatus          string          `gorm:"column:invoice_status" json:"invoice_status"`
	ValidityDate           sql.NullTime    `gorm:"column:validity_date" json:"validity_date"`
	Note                   string          `gorm:"column:note" json:"note"`
	CurrencyRate           sql.NullFloat64 `gorm:"column:currency_rate" json:"currency_rate"`
	AmountUntaxed          sql.NullFloat64 `gorm:"column:amount_untaxed" json:"amount_untaxed"`
	AmountTax              sql.NullFloat64 `gorm:"column:amount_tax" json:"amount_tax"`
	AmountTotal            sql.NullFloat64 `gorm:"column:amount_total" json:"amount_total"`
	AmountToInvoice        sql.NullFloat64 `gorm:"column:amount_to_invoice" json:"amount_to_invoice"`
	Locked                 sql.NullBool    `gorm:"column:locked" json:"locked"`
	RequireSignature       sql.NullBool    `gorm:"column:require_signature" json:"require_signature"`
	RequirePayment         sql.NullBool    `gorm:"column:require_payment" json:"require_payment"`
	CreateDate             sql.NullTime    `gorm:"column:create_date" json:"create_date"`
	CommitmentDate         sql.NullTime    `gorm:"column:commitment_date" json:"commitment_date"`
	DateOrder              time.Time       `gorm:"column:date_order" json:"date_order"`
	SignedOn               sql.NullTime    `gorm:"column:signed_on" json:"signed_on"`
	WriteDate              sql.NullTime    `gorm:"column:write_date" json:"write_date"`
	PrepaymentPercent      sql.NullFloat64 `gorm:"column:prepayment_percent" json:"prepayment_percent"`
	CarrierID              sql.NullInt64   `gorm:"column:carrier_id" json:"carrier_id"`
	DeliveryMessage        string          `gorm:"column:delivery_message" json:"delivery_message"`
	DeliveryRatingSuccess  sql.NullBool    `gorm:"column:delivery_rating_success" json:"delivery_rating_success"`
	RecomputeDeliveryPrice sql.NullBool    `gorm:"column:recompute_delivery_price" json:"recompute_delivery_price"`
	ShippingWeight         sql.NullFloat64 `gorm:"column:shipping_weight" json:"shipping_weight"`
	PendingEmailTemplateID sql.NullInt64   `gorm:"column:pending_email_template_id" json:"pending_email_template_id"`
	WebsiteID              sql.NullInt64   `gorm:"column:website_id" json:"website_id"`
	ShopWarning            string          `gorm:"column:shop_warning" json:"shop_warning"`
	AccessPointAddress     string          `gorm:"column:access_point_address" json:"access_point_address"`
	CartRecoveryEmailSent  sql.NullBool    `gorm:"column:cart_recovery_email_sent" json:"cart_recovery_email_sent"`
	IidDeliveryID          sql.NullInt64   `gorm:"column:iid_delivery_id" json:"iid_delivery_id"`
	DeliveryProviderID     sql.NullInt64   `gorm:"column:delivery_provider_id" json:"delivery_provider_id"`
	DeliveryProductID      sql.NullInt64   `gorm:"column:delivery_product_id" json:"delivery_product_id"`
	Incoterm               sql.NullInt64   `gorm:"column:incoterm" json:"incoterm"`
	WarehouseID            int             `gorm:"column:warehouse_id" json:"warehouse_id"`
	ProcurementGroupID     sql.NullInt64   `gorm:"column:procurement_group_id" json:"procurement_group_id"`
	IncotermLocation       string          `gorm:"column:incoterm_location" json:"incoterm_location"`
	PickingPolicy          string          `gorm:"column:picking_policy" json:"picking_policy"`
	DeliveryStatus         string          `gorm:"column:delivery_status" json:"delivery_status"`
	EffectiveDate          sql.NullTime    `gorm:"column:effective_date" json:"effective_date"`
	SaleOrderTemplateID    sql.NullInt64   `gorm:"column:sale_order_template_id" json:"sale_order_template_id"`
	ProjectID              sql.NullInt64   `gorm:"column:project_id" json:"project_id"`
	IidStoreID             sql.NullInt64   `gorm:"column:iid_store_id" json:"iid_store_id"`
	PaymentMethodID        sql.NullInt64   `gorm:"column:payment_method_id" json:"payment_method_id"`
	DeliveryCarrierID      sql.NullInt64   `gorm:"column:delivery_carrier_id" json:"delivery_carrier_id"`
	ResponseCodeID         sql.NullInt64   `gorm:"column:response_code_id" json:"response_code_id"`
	BankID                 sql.NullInt64   `gorm:"column:bank_id" json:"bank_id"`
	AccHolderName          string          `gorm:"column:acc_holder_name" json:"acc_holder_name"`
	AccHolderNumber        string          `gorm:"column:acc_holder_number" json:"acc_holder_number"`
	PaymentURL             string          `gorm:"column:payment_url" json:"payment_url"`
	PaymentDate            sql.NullTime    `gorm:"column:payment_date" json:"payment_date"`
	ReceiptDate            sql.NullTime    `gorm:"column:receipt_date" json:"receipt_date"`
	ExpiryDate             sql.NullTime    `gorm:"column:expiry_date" json:"expiry_date"`
	ParentOrderID          sql.NullInt64   `gorm:"column:parent_order_id" json:"parent_order_id"`
	AmountUnpaid           sql.NullFloat64 `gorm:"column:amount_unpaid" json:"amount_unpaid"`
}

func (SaleOrder) TableName() string {
	return "sale_order"
}
