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

type Patch struct {
	Operation Operation   `json:"op,omitempty"`
	Path      string      `json:"path,omitempty"`
	Value     interface{} `json:"value,omitempty"`
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
