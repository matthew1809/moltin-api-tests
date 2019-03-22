package remove

import (
	"fmt"
	"net/http"
	"sync"

	"github.com/matthew1809/cart-tests/models"
)

// ItemFromCart removes an item from the cart
func ItemFromCart(wg *sync.WaitGroup, ID int, baseURL string, client http.Client, item models.CartItem, token string) {
	defer wg.Done()

	itemID := item.ID
	fullURL := baseURL + "/" + itemID

	req, err := http.NewRequest("DELETE", fullURL, nil)
	req.Header.Set("Authorization", token)
	req.Header.Set("content-type", "application/json")

	res, err := client.Do(req)

	if err != nil {
		fmt.Printf("Error removing from cart, failed with %s\n", err)
	}

	res.Body.Close()
}
