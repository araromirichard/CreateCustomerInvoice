package main

import (
	"os"
	"strconv"

	"github.com/karosaxy/paystack-client/pkg/client/paystack"
)

var client *paystack.Client

func init() {

	client = paystack.NewClient(os.Getenv("PAYSTACK_BASE_URL"), os.Getenv("PAYSTACK_SECRET_KEY"))
}

type CreateInvoiceInputEvent struct {
	CustomerEmail    string `json:"customer_email"`
	CustomerId       string `json:"customer_id"`
	Amount           int32  `json:"amount"`
	Currency         string `json:"currency"`
	DueDate          string `json:"due_date"`
	SendNotification bool   `json:"send_notification"`
	Draft            bool   `json:"draft"`
}

type CreateInvoiceOutputEvent struct {
	Message string            `json:"message"`
	Invoice *paystack.Invoice `json:"data"`
}

func createInvoice(ciie CreateInvoiceInputEvent) (*CreateInvoiceOutputEvent, error) {

	// get or create customer
	customerId, err := getOrCreateCustomer(ciie.CustomerId, ciie.CustomerEmail)

	if err != nil {
		return nil, err
	}

	// create invoice
	invoice, err := createCustomerInvoice(customerId, ciie)

	if err != nil {
		return nil, err
	}

	return &CreateInvoiceOutputEvent{"Customer invoice successfully created", invoice}, nil
}

func getOrCreateCustomer(customerId, customerEmail string) (string, error) {
	cId := customerId

	_, err := client.GetCustomer(customerId)

	if err != nil {

		if err == paystack.ErrCustomerNotFound {
			ccr, err := client.CreateCustomer(paystack.CreateCustomerRequest{Email: customerEmail})

			if err != nil {
				return "", err
			}
			cId = strconv.Itoa(ccr.Data.ID)
		} else {
			return "", err
		}

	}
	return cId, nil
}

func createCustomerInvoice(customerId string, ciie CreateInvoiceInputEvent) (*paystack.Invoice, error) {
	// payload
	cri := paystack.CreateInvoiceRequest{
		DueDate:          ciie.DueDate,
		Amount:           ciie.Amount,
		Currency:         ciie.Currency,
		SendNotification: ciie.SendNotification,
		Draft:            ciie.Draft,
		CustomerId:       customerId,
	}
	crr, err := client.CreateInvoice(cri)

	if err != nil {
		return nil, err
	}

	return &crr.Invoice, nil
}
