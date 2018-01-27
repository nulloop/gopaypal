package gopaypal

import "time"

type OutstandingBalance struct {
}

type AgreementDetails struct {
	OutstandingBalance *Amount    `json:"outstanding_balance,omitempty"`
	CyclesRemaining    string     `json:"cycles_remaining,omitempty"`
	CyclesCompleted    string     `json:"cycles_completed,omitempty"`
	NextBillingDate    *time.Time `json:"next_billing_date,omitempty"`
	LastPaymentDate    *time.Time `json:"last_payment_date,omitempty"`
	LastPaymentAmount  *Amount    `json:"last_payment_amount,omitempty"`
	FinalPaymentDate   *time.Time `json:"final_payment_date,omitempty"`
	FailedPaymentCount string     `json:"failed_payment_count,omitempty"`
}

type Address struct {
	Line1       string      `json:"line1,omitempty"`
	Line2       string      `json:"line2,omitempty"`
	City        string      `json:"city,omitempty"`
	State       string      `json:"state,omitempty"`
	CountryCode CountryCode `json:"country_code,omitempty"`
	PostalCode  string      `json:"postal_code,omitempty"`
}

type CreditCard struct {
	ID                 string           `json:"id,omitempty"`
	Number             string           `json:"number,omitempty"`
	Type               string           `json:"type,omitempty"`
	ExpireMonth        int              `json:"expire_month,omitempty"`
	ExpireYear         int              `json:"expire_year,omitempty"`
	CVV2               int              `json:"ccv2,omitempty"`
	FirstName          string           `json:"first_name,omitempty"`
	LastName           string           `json:"last_name,omitempty"`
	BillingAddress     *Address         `json:"billing_address,omitempty"`
	ExternalCustomerID string           `json:"external_customer_id,omitempty"`
	State              *CreditCardState `json:"state,omitempty"`
	ValidUntil         string           `json:"valid_until,omitempty"`
	Links              []*Link          `json:"links,omitempty"`
}

type FundingInstrument struct {
	CreditCard CreditCard `json:"credit_card,omitempty"`
}

type PayerInfo struct {
	Email          string   `json:"email,omitempty"`
	FirstName      string   `json:"first_name,omitempty"`
	LastName       string   `json:"last_name,omitempty"`
	PayerID        string   `json:"payer_id,omitempty"`
	BillingAddress *Address `json:"billing_address,omitempty"`
}

type Payer struct {
	PaymentMethod      PaymentMethod        `json:"payment_method,omitempty"`
	FundingInstruments []*FundingInstrument `json:"funding_instruments,omitempty"`
	FundingOptionID    string               `json:"funding_option_id,omitempty"`
	PayerInfo          *PayerInfo           `json:"payer_info,omitempty"`
}

type BillingAgreement struct {
	Name                        string               `json:"name,omitempty"`
	Description                 string               `json:"description,omitempty"`
	StartDate                   *time.Time           `json:"start_date,omitempty"`
	AgreementDetails            *AgreementDetails    `json:"agreement_details,omitempty"`
	Payer                       *Payer               `json:"payer,omitempty"`
	ShippingAddress             *Address             `json:"shipping_address,omitempty"`
	OverrideMerchantPreferences *MerchantPreferences `json:"override_merchant_preferences,omitempty"`
	OverrideChargeModels        []*ChargeModelType   `json:"override_charge_models"`
	Plan                        *BillingPlan         `json:"plan,omitempty"`
	Links                       []*Link              `json:"links,omitempty"`
}
