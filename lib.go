package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/karosaxy/paystack-client/pkg/client/paystack"
)

var client *paystack.Client

func init() {

	client = paystack.NewClient("https://api.paystack.co", "sk_test_1374752eb291108be0fff0424ca81369b0f266ac")
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

	_, err := client.GetCustomer(customerEmail)

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

	//send Response body to a webhook
	postBody, err := json.Marshal(crr)

	if err != nil {
		return nil, err
	}

	resp, err := http.Post("https://webhook.site/a1b30d87-faeb-4dbc-8025-5f370470a0d5", "application/json", bytes.NewBuffer(postBody))

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	if resp.StatusCode == http.StatusCreated {
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			//Failed to read response.
			return nil, err
		}

		//Convert bytes to String and print
		jsonStr := string(body)
		fmt.Println("Response: ", jsonStr)

	} else {
		//The status is not Created. print the error.
		fmt.Println("Get failed with error: ", resp.Status)
	}

	return &crr.Invoice, nil
}

