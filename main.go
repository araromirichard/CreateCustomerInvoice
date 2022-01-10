package main

import (
	"context"

	"github.com/aws/aws-lambda-go/lambda"
)

func HandleRequest(ctx context.Context, ciie CreateInvoiceInputEvent) (*CreateInvoiceOutputEvent, error) {
	return createInvoice(ciie)
}

func main() {
	lambda.Start(HandleRequest)

}
