package gopaypal

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"path"
	"time"
)

type Token struct {
	Nonce       string    `json:"nonce,omitempty"`
	AccessToken string    `json:"access_token,omitempty"`
	TokenType   string    `json:"token_type,omitempty"`
	AppID       string    `json:"app_id,omitempty"`
	ExpiresIn   int64     `json:"expires_in,omitempty"`
	ExpiresAt   time.Time `json:"_"`
}

func (t *Token) Valid() bool {
	threshold := time.Duration(60) * time.Second
	return t.ExpiresAt.Sub(time.Now()) > threshold
}

type Paypal struct {
	Mode         Mode
	ClientID     string
	ClientSecret string
	Token        Token

	client http.Client

	updateToken func(Token)
}

func (p *Paypal) url(urlPath string) string {
	return "https://" + path.Join(p.Mode.String(), urlPath)
}

func (p *Paypal) updateAuthToken() error {
	if p.Token.Valid() {
		return nil
	}

	body := bytes.NewBuffer([]byte("grant_type=client_credentials"))

	req, err := http.NewRequest("POST", p.url("/v1/oauth2/token"), body)
	if err != nil {
		return err
	}

	req.Header.Set("Content-type", "application/x-www-form-urlencoded")
	req.SetBasicAuth(p.ClientID, p.ClientSecret)

	resp, err := p.client.Do(req)
	if err != nil {
		return err
	}

	err = json.NewDecoder(resp.Body).Decode(&p.Token)
	if err != nil {
		return err
	}

	p.Token.ExpiresAt = time.Now().Add(time.Duration(p.Token.ExpiresIn) * time.Second)

	if p.updateToken != nil {
		p.updateToken(p.Token)
	}

	return nil
}

func (p *Paypal) clientWithAuth(method, path string, body io.Reader) (*http.Response, error) {
	err := p.updateAuthToken()
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(method, p.url(path), body)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", p.Token.AccessToken))

	return p.client.Do(req)
}

func (p *Paypal) UpdateToken(updateToken func(Token)) {
	p.updateToken = updateToken
}

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
	Currency Currency `json:"currency"`
	Value    string   `json:"value"`
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
