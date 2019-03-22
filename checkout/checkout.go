package checkout

import (
	"github.com/matthew1809/cart-tests/models"
	"github.com/matthew1809/cart-tests/request"
	"net/http"
	"encoding/json"
	"fmt"
	"bytes"
)

// Checkout converts a moltin cart into an unpaid, unshipped order
func Checkout(baseURL string, client *http.Client, token string, cartID string) models.CheckoutResponse {
	fullURL := baseURL + "/carts/" + cartID + "/checkout/"

	checkoutPayload := &models.CheckoutRequest{
			Customer: models.Customer{
				Name: "John",
				Email: "jj@mail.com",
			},
			BillingAddress: models.BillingAddress{
				FirstName: "John",
				LastName: "Bloggs",
				Line1: "house",
				Line2: "street",
				Postcode: "DH12HU",
				County: "Durham",
				Country: "UK",
			},
			ShippingAddress: models.ShippingAddress{
				FirstName: "John",
				LastName: "Bloggs",
				PhoneNumber: "123456789",
				Line1: "house",
				City: "Durham",
				Line2: "street",
				CompanyName: "name",
				Postcode: "DH12HU",
				County: "Durham",
				Country: "UK",
				Instructions: "none",
			},
	}

	FullPayload := &models.TopLevelRequest{
		Data: *checkoutPayload,
	}

	bytesRepresentation, err := json.Marshal(FullPayload)

	if err != nil {
		fmt.Printf("Error marshalling JSON, failed with %s\n", err)
	}

	res := request.GenericRequest(fullURL, *client, "POST",  bytes.NewBuffer(bytesRepresentation), token, "add.checkout", 201)

	var order models.CheckoutResponse
	unmarshallErr := json.Unmarshal(res, &order)

	if unmarshallErr != nil {
		fmt.Printf("Error reading get cart response, failed with %s\n", unmarshallErr)
	}

	return order
}