package get

import (
	"encoding/json"
	"fmt"
	"net/http"
	"github.com/matthew1809/cart-tests/models"
	"github.com/matthew1809/cart-tests/request"
)

// Order fetches a single moltin order by ID
func Order(baseURL string, client http.Client, token string, orderID string) models.OrderResponse {

	fullURL := baseURL + "/orders/" + orderID
	res := request.GenericRequest(fullURL, client, "GET", nil, token, "get.Order", 200)
	
	var order models.OrderResponse
	unmarshallErr := json.Unmarshal(res, &order)

	if unmarshallErr != nil {
		fmt.Printf("Error reading get cart response, failed with %s\n", unmarshallErr)
	}

	return order
}