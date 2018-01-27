package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/nulloop/gopaypal"
)

func main() {
	paypal := &gopaypal.Paypal{
		Mode:         gopaypal.SandboxMode,
		ClientID:     "",
		ClientSecret: "",
	}

	paypal.UpdateToken(func(token gopaypal.Token) {
		data, _ := json.Marshal(&token)
		fmt.Println("Token: ", string(data))
	})

	paymentDifinition := []*gopaypal.PaymentDefinition{
		&gopaypal.PaymentDefinition{
			Name:               "One Month Trial",
			Type:               gopaypal.TrialPaymentDefinitionType,
			Cycles:             "1",
			Frequency:          gopaypal.MonthFrequencyType,
			FrequenceyInterval: "1",
			Amount: &gopaypal.Amount{
				Currency: gopaypal.UnitedStatesDollar,
				Value:    "0",
			},
		},
		&gopaypal.PaymentDefinition{
			Name:               "Subscrtiption",
			Type:               gopaypal.RegularPaymentDefinitionType,
			Cycles:             "0",
			Frequency:          gopaypal.MonthFrequencyType,
			FrequenceyInterval: "12",
			Amount: &gopaypal.Amount{
				Currency: gopaypal.UnitedStatesDollar,
				Value:    "2",
			},
		},
	}

	merchantPreferences := &gopaypal.MerchantPreferences{
		AutoBillAmount:          gopaypal.Yes,
		ReturnURL:               "http://localhost:8080/success",
		CancelURL:               "http://localhost:8080/cancel",
		InitialFailAmountAction: gopaypal.ContinueInitialAction,
		MaxFailAttempts:         "0",
	}

	billingPlan := &gopaypal.BillingPlan{
		Name:                "Subscription",
		Description:         "Subscription for  Service",
		Type:                gopaypal.InfiniteBillingPlanType,
		PaymentDefinitions:  paymentDifinition,
		MerchantPreferences: merchantPreferences,
	}

	data, _ := json.Marshal(billingPlan)
	fmt.Println("Req: ", string(data))

	err := billingPlan.Create(paypal)
	if err != nil {
		log.Fatal("Error: ", err)
	}

	data, _ = json.Marshal(billingPlan)
	fmt.Println("Resp: ", string(data))
}
