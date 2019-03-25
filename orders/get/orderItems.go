package get

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/matthew1809/cart-tests/models"
	"github.com/matthew1809/cart-tests/request"
)

// OrderItems fetches order items
func OrderItems(baseURL string, client http.Client, token string, orderID string) models.OrderItemResponse {

	fullURL := baseURL + "/orders/" + orderID + "/items"

	res := request.GenericRequest(fullURL, client, "GET", nil, token, "get.OrderItems", 200)

	var items models.OrderItemResponse
	unmarshallErr := json.Unmarshal(res, &items)

	if unmarshallErr != nil {
		fmt.Printf("Error reading get order items response, failed with %s\n", unmarshallErr)
	} else {
		fmt.Println("Number of unique order items: ", len(items.Data))
	}

	return items
}
