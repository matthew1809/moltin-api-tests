package get

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sync"
	"github.com/matthew1809/cart-tests/models"
	"github.com/matthew1809/cart-tests/request"
)

// Cart fetches cart
func Cart(baseURL string, client http.Client, token string) int {

	var wg sync.WaitGroup
	wg.Add(1)
	res := request.GenericRequest(&wg, 0, baseURL, client, "GET", nil, token, "get.Cart", 200)
	
	var cart models.CartResponse
	unmarshallErr := json.Unmarshal(res, &cart)

	if unmarshallErr != nil {
		fmt.Printf("Error reading get cart response, failed with %s\n", unmarshallErr)
	}

	return cart.Data.Meta.DisplayPrice.Tax.Amount
}
