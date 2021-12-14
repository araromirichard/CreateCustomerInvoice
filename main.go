package main

import (
	"context"

	"github.com/aws/aws-lambda-go/lambda"
)

// const SECRET_KEY = "sk_test_3cec88e7ece6d9f5b69f33c013b94cc9142bf161"
// const BASE_URL = "https://api.paystack.co"

func HandleRequest(ctx context.Context, ciie CreateInvoiceInputEvent) (*CreateInvoiceOutputEvent, error) {
	return createInvoice(ciie)
}

func main() {
	lambda.Start(HandleRequest)
}
