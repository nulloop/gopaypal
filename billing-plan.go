package gopaypal

import (
	"fmt"
	"time"
)

type BillingPlanTerm struct {
	ID               string              `json:"id,omitempty"`
	Type             BillingPlanTermType `json:"type,omitempty"`
	MaxBillingAmount Amount              `json:"max_billing_amount,omitempty"`
	Occurences       string              `json:"occurences,omitempty"`
	AmountRange      Amount              `json:"amount_range,omitempty"`
	BuyerEditable    string              `json:"buyer_editable,omitempty"`
}

type Link struct {
	Href   string `json:"href,omitempty"`
	Rel    string `json:"rel,omitempty"`
	Method string `json:"method,omitempty"`
}

type BillingPlan struct {
	ID                  string               `json:"id,omitempty"`
	Name                string               `json:"name,omitempty"`
	Description         string               `json:"description,omitempty"`
	Type                BillingPlanType      `json:"type,omitempty"`
	State               BillingPlanState     `json:"state,omitempty"`
	CreateTime          *time.Time           `json:"create_time,omitempty"`
	UpdateTime          *time.Time           `json:"update_time,omitempty"`
	PaymentDefinitions  []*PaymentDefinition `json:"payment_definitions,omitempty"`
	Terms               []*BillingPlanTerm   `json:"terms,omitempty"`
	MerchantPreferences *MerchantPreferences `json:"merchant_preferences,omitempty"`
	Links               []*Link              `json:"links,omitempty"`
}

func (b *BillingPlan) Create(paypal *Paypal) error {
	resp, err := paypal.clientWithAuth("POST", "v1/payments/billing-plans", anyToReader(b))
	if err != nil {
		return err
	}

	// catches any api errors
	if err = catchError(resp); err != nil {
		return err
	}

	return readerToAny(resp.Body, b)
}

func (b *BillingPlan) UpdateState(state BillingPlanState, paypal *Paypal) error {
	patch := &Patch{
		Operation: ReplaceOperation,
		Path:      "/",
		Value: &BillingPlan{
			State: state,
		},
	}

	resp, err := paypal.clientWithAuth("PATCH", fmt.Sprintf("v1/payments/billing-plans/%s", b.ID), anyToReader(patch))
	if err != nil {
		return err
	}

	return catchError(resp)
}

type ChargeModel struct {
	ID     string          `json:"id,omitempty"`
	Type   ChargeModelType `json:"type,omitempty"`
	Amount *Amount         `json:"amount,omitempty"`
}

type PaymentDefinition struct {
	ID                 string                `json:"id,omitempty"` // OPTIONAL
	Name               string                `json:"name,omitempty"`
	Type               PaymentDefinitionType `json:"type,omitempty"`
	Frequency          FrequencyType         `json:"frequency,omitempty"`
	FrequenceyInterval string                `json:"frequency_interval,omitempty"`
	Cycles             string                `json:"cycles,omitempty"`
	Amount             *Amount               `json:"amount,omitempty"`
	ChargeModels       []*ChargeModel        `json:"charge_models,omitempty"` // OPTIONAL
}

type Amount struct {
	Currency Currency `json:"currency,omitempty"`
	Value    string   `json:"value,omitempty"`
}

type MerchantPreferences struct {
	SetupFee                *Amount       `json:"setup_fee,omitempty"` // OPTIONAL
	CancelURL               string        `json:"cancel_url,omitempty"`
	ReturnURL               string        `json:"return_url,omitempty"`
	NotifyURL               string        `json:"notify_url,omitempty"`                 // OPTIONAL
	MaxFailAttempts         string        `json:"max_fail_attempts,omitempty"`          // OPTIONAL
	AutoBillAmount          Bool          `json:"auto_bill_amount,omitempty"`           // OPTIONAL
	InitialFailAmountAction InitialAction `json:"initial_fail_amount_action,omitempty"` // OPTIONAL
	AcceptedPaymentType     string        `json:"accepted_payment_type,omitempty"`      // READONLY
	CharSet                 string        `json:"char_set,omitempty"`                   // READONLY
}
