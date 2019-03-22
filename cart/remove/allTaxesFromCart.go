package remove

import (
	"fmt"
	"net/http"
	"sync"

	"github.com/matthew1809/cart-tests/models"
)

// TaxFromCartAsync does a thing
func TaxFromCartAsync(baseURL string, client *http.Client, items models.CartItemResponse, token string) {
	var wg sync.WaitGroup

	for i := 0; i < len(items.Data); i++ {
		wg.Add(1)
		go TaxItemFromCartItem(&wg, baseURL, items.Data[i], *client, token)
	}

	wg.Wait()
	fmt.Println("Removing taxes: Completed")
	return
}
