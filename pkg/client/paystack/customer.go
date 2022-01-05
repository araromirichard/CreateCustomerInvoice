package paystack

import (
	"errors"
	"time"
)

var ErrCustomerNotFound = errors.New("Customer with ID/Email not found")

type Customer struct {
	Transactions   []interface{} `json:"transactions"`
	Subscriptions  []interface{} `json:"subscriptions"`
	Authorizations []interface{} `json:"authorizations"`
	FirstName      string        `json:"first_name"`
	LastName       string        `json:"last_name"`
	Email          string        `json:"email"`
	Phone          string        `json:"phone"`
	Metadata       struct {
	} `json:"metadata"`
	Domain                string        `json:"domain"`
	CustomerCode          string        `json:"customer_code"`
	RiskAction            string        `json:"risk_action"`
	ID                    int           `json:"id"`
	Integration           int           `json:"integration"`
	CreatedAt             time.Time     `json:"createdAt"`
	UpdatedAt             time.Time     `json:"updatedAt"`
	TotalTransactions     int           `json:"total_transactions"`
	TotalTransactionValue []interface{} `json:"total_transaction_value"`
	DedicatedAccount      interface{}   `json:"dedicated_account"`
	Identified            bool          `json:"identified"`
	Identifications       interface{}   `json:"identifications"`
}

type Amount int

type DateTime string

type LineItem struct {
	Name   string `json:"name"`
	Amount Amount `json:"amount"`
}

type Tax struct {
	Name   string `json:"name"`
	Amount Amount `json:"amount"`
}

type Invoice struct {
	ID               int        `json:"id"`
	Domain           string     `json:"domain"`
	Amount           Amount     `json:"amount"`
	Currency         string     `json:"currency"`
	DueDate          DateTime   `json:"due_date"`
	HasInvoice       bool       `json:"has_invoice"`
	InvoiceNumber    int        `json:"invoice_number"`
	Description      string     `json:"description"`
	PdfUrl           string     `json:"pdf_url"`
	LineItems        []LineItem `json:"line_items"`
	Tax              []Tax      `json:"tax"`
	Customer         Customer   `json:"customer"`
	RequestCode      string     `json:"request_code"`
	Status           string     `json:"status"`
	Paid             bool       `json:"paid"`
	PaidAt           DateTime   `json:"paid_at"`
	CreatedAt        DateTime   `json:"created_at"`
	SendNotification bool       `json:"send_notification"`
}
type Meta struct {
	Total     int `json:"total"`
	Skipped   int `json:"skipped"`
	PerPage   int `json:"perPage"`
	Page      int `json:"page"`
	PageCount int `json:"pageCount"`
}

type ListInvoicesResponse struct {
	Status  bool      `json:"status"`
	Message string    `json:"message"`
	Meta    Meta      `json:"meta"`
	Data    []Invoice `json:"data"`
}

type CreateInvoiceRequest struct {
	CustomerId       string `json:"customer"`
	Amount           int32  `json:"amount"`
	Currency         string `json:"currency"`
	DueDate          string `json:"due_date"`
	SendNotification bool   `json:"send_notification"`
	Draft            bool   `json:"draft"`
}

type CreateInvoiceResponse struct {
	Invoice Invoice `json:"data"`
}

type GetCustomerResponse struct {
	Status  bool     `json:"status"`
	Message string   `json:"message"`
	Data    Customer `json:"data"`
}
type CreateCustomerResponse struct {
	Customer Customer `json:"data"`
}

type CreateCustomerRequest struct {
	Email string `json:"email"`
}
