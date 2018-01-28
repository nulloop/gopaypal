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

	payer := gopaypal.NewPayer(gopaypal.PaypalPaymentMethod)

	billingAgreement := gopaypal.NewBillingAgreement("Rondo Subscription", "Rondo Subscription Agreement", payer)

	data, _ := json.Marshal(billingAgreement)
	fmt.Println("Before sending: ", string(data))

	err = billingAgreement.Create(context, &gopaypal.BillingPlan{
		ID: "P-7VL15599HW676163HHKMVGEQ",
	})

	if err != nil {
		log.Fatal(err)
	}

	// data, err = json.Marshal(billingAgreement)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	fmt.Println("Approval URL: ", billingAgreement.ApprovalURL())
}
