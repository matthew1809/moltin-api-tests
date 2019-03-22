package remove

import (
	"fmt"
	"net/http"
	"sync"

	"github.com/matthew1809/cart-tests/models"
)

// ProductsFromCartAsync uses a waitgroup and go routines to remove all products from the cart
func ProductsFromCartAsync(baseURL string, client *http.Client, items models.CartItemResponse, token string) {

	var wg sync.WaitGroup

	for i := 0; i < len(items.Data); i++ {
		wg.Add(1)
		go ItemFromCart(&wg, i, baseURL, *client, items.Data[i], token)
	}

	wg.Wait()
	fmt.Println("Removing products: Completed")
}
