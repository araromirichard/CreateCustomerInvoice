package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/karosaxy/paystack-client/pkg/client/paystack"
)

const SECRET_KEY = "sk_test_3cec88e7ece6d9f5b69f33c013b94cc9142bf161"
const BASE_URL = "https://api.paystack.co"

func HandleRequest(ctx context.Context, event paystack.CreateInvoiceRequest) (string, error) {
	return fmt.Sprintf("%+v", event), nil
}

func main() {

	lambda.Start(HandleRequest)

	secretKey := os.Getenv("PAYSTACK_SECRET_KEY")
	client := paystack.NewClient(BASE_URL, secretKey)

	ccq := paystack.CreateCustomerRequest{"man@go.com"}

	if err := client.CreateCustomer(ccq); err != nil {
		log.Fatalf("could not create customer. %v", err)
	}

	log.Println("Customer successfully created")

}
