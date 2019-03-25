package get

import (
	"encoding/json"
	"fmt"
	"net/http"
	"github.com/matthew1809/cart-tests/models"
	"github.com/matthew1809/cart-tests/request"
)

// Orders fetches the first page of moltin orders
func Orders(baseURL string, client http.Client, token string) models.OrdersResponse {

	fullURL := baseURL + "/orders"
	res := request.GenericRequest(fullURL, client, "GET", nil, token, "get.Orders", 200)
	
	var orders models.OrdersResponse
	unmarshallErr := json.Unmarshal(res, &orders)

	if unmarshallErr != nil {
		fmt.Printf("Error reading get cart response, failed with %s\n", unmarshallErr)
	}

	return orders
}