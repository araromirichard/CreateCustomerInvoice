package paystack

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
)

type Client struct {
	client    http.Client
	baseUrl   string
	secretKey string
}

func NewClient(baseUrl, secretKey string) *Client {

	return &Client{*http.DefaultClient, baseUrl, secretKey}

}

func (cl *Client) GetCustomer(customerEmail string) (*GetCustomerResponse, error) {
	req, err := http.NewRequest(http.MethodGet, cl.baseUrl+"/customer/"+customerEmail, nil)

	if err != nil {
		return nil, err
	}

	cl.addRequiredHeaders(req)

	response, err := cl.client.Do(req)

	if err != nil {
		return nil, err
	}

	defer response.Body.Close()

	if response.StatusCode == http.StatusNotFound {
		return nil, ErrCustomerNotFound

	}
	if response.StatusCode == http.StatusOK {

		var customer GetCustomerResponse

		body, err := io.ReadAll(response.Body)
		if err != nil {
			return nil, err
		}
		log.Printf("Get customer response from Paystack: %s \n", string(body))

		if err := json.NewDecoder(bytes.NewReader(body)).Decode(&customer); err != nil {
			return nil, err
		}

		return &customer, nil

	}

	// switch response.StatusCode {
	// case http.StatusNotFound:
	// 	return nil, ErrCustomerNotFound
	// case http.StatusUnauthorized:
	// 	return nil, fmt.Errorf("invalid secret key")
	// case http.StatusOK:
	// 	return getCustomer(response)
	// default:
	// 	return nil, fmt.Errorf("unknown response code: %d", response.StatusCode)

	// }
	return nil, err
}

func (cl *Client) CreateCustomer(ccq CreateCustomerRequest) (*CreateCustomerResponse, error) {

	body, err := json.Marshal(ccq)

	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(http.MethodPost, cl.baseUrl+"/customer", bytes.NewReader(body))

	if err != nil {
		return nil, err
	}

	cl.addRequiredHeaders(req)
	req.Header.Add("Content-Type", "application/json")

	response, err := cl.client.Do(req)

	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	if response.StatusCode == http.StatusOK || response.StatusCode == http.StatusCreated {
		var ccr CreateCustomerResponse
		if err = json.NewDecoder(response.Body).Decode(&ccr); err != nil {
			return nil, err
		}
		return &ccr, nil
	}

	return nil, fmt.Errorf("could not create customer. Invalid status : %d", response.StatusCode)
}

func (cl *Client) addRequiredHeaders(req *http.Request) {
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", cl.secretKey))
	req.Header.Add("Accept", "application/json")

}

func (cl *Client) CreateInvoice(ciq CreateInvoiceRequest) (*CreateInvoiceResponse, error) {

	payload, _ := json.Marshal(ciq)

	log.Printf("Payload for createInvoice: %s \n", string(payload))

	//fmt.Printf("Payload for createInvoice: %s \n", string(payload))

	req, err := http.NewRequest(http.MethodPost, cl.baseUrl+"/paymentrequest", bytes.NewReader(payload))

	if err != nil {
		return nil, err
	}

	cl.addRequiredHeaders(req)
	req.Header.Add("Content-Type", "application/json")

	response, err := cl.client.Do(req)

	if err != nil {
		return nil, err
	}

	defer response.Body.Close()

	body, _ := ioutil.ReadAll(response.Body)

	log.Printf("response from paystack: %s \n", string(body))

	if response.StatusCode == http.StatusOK || response.StatusCode == http.StatusCreated {
		var cir CreateInvoiceResponse

		if err = json.Unmarshal(body, &cir); err != nil {
			return nil, err
		}
		return &cir, nil
	}
	
	return nil, fmt.Errorf("could not create Invoice. Invalid status : %d, %s", response.StatusCode, string(body))
}
