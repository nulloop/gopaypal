package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/nulloop/gopaypal"
)

func main() {
	context := &gopaypal.Context{}

	err := func() error {
		file, err := os.Open("./token.json")
		if err != nil {
			return err
		}
		defer file.Close()

		return json.NewDecoder(file).Decode(context)
	}()

	if err != nil {
		log.Fatal(err)
	}

	context.Mode = gopaypal.SandboxMode

	context.UpdateToken(func(token *gopaypal.Token) {
		// use this to store the token into database

		// data, _ := json.Marshal(token)
		// fmt.Printf("Token: %s\n\n", string(data))
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
			FrequenceyInterval: "1",
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

	// data, _ := json.Marshal(billingPlan)
	// fmt.Println("Req: ", string(data))

	err = billingPlan.Create(context)
	if err != nil {
		log.Fatal("Error: ", err)
	}

	err = billingPlan.UpdateState(gopaypal.ActiveBillingPlanState, context)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Created and activated Billing Plan with id : ", billingPlan.ID)
}
