package gopaypal

import "fmt"

type ErrorDetails struct {
	Field string `json:"field,omitempty"`
	Issue string `json:"issue,omitempty"`
}

func (ed *ErrorDetails) String() string {
	return fmt.Sprintf("\n%s: %s\n", ed.Field, ed.Issue)
}

type PayPalError struct {
	Name            string          `json:"name,omitempty"`
	Message         string          `json:"message,omitempty"`
	InformationLink string          `json:"information_link,omitempty"`
	DebugID         string          `json:"debug_id,omitempty"`
	Details         []*ErrorDetails `json:"details,omitempty"`
}

func (e *PayPalError) Error() string {
	return fmt.Sprintf("%s %s - %s: %s", e.Name, e.Message, e.InformationLink, e.Details)
}
