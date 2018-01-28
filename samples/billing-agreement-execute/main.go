package main

import (
	"encoding/json"
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

	err = gopaypal.BillingAgreementExecute(context, "EC-2NM13782DU4310537")
	if err != nil {
		log.Fatal(err)
	}
}
