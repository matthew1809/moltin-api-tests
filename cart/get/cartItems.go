package get

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/matthew1809/cart-tests/models"
	"github.com/matthew1809/cart-tests/request"
)

// CartItems fetches cart items
func CartItems(baseURL string, client http.Client, token string, cartID string) models.CartItemResponse {

	cartItemsBaseURL := baseURL + "/carts/" + cartID + "/items"

	res := request.GenericRequest(cartItemsBaseURL, client, "GET", nil, token, "get.CartItems", 200)

	var items models.CartItemResponse
	unmarshallErr := json.Unmarshal(res, &items)

	if unmarshallErr != nil {
		fmt.Printf("Error reading get cart items response, failed with %s\n", unmarshallErr)
	} else {
		fmt.Println("Number of unique cart items: ", len(items.Data))
	}

	return items
}
