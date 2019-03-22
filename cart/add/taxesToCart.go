package add

import (
	"fmt"
	"net/http"
	"sync"

	"github.com/matthew1809/cart-tests/models"
)

// TaxesToCart uses a waitgroup and go routines to add a single product to the cart multiple times
func TaxesToCart(baseURL string, client *http.Client, items models.CartItemResponse, token string) {
	// fmt.Println("Adding taxes")
	var wg sync.WaitGroup

	for i := 0; i < len(items.Data); i++ {
		// fmt.Println("Main: Starting worker", i)
		wg.Add(1)
		go TaxToCartItem(&wg, baseURL, items.Data[i].ID, *client, token)
	}

	// fmt.Println("Main: Waiting for workers to finish")
	wg.Wait()
	fmt.Println("Adding taxes: Completed")
	return
}
