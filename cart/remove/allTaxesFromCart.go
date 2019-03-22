package remove

import (
	"fmt"
	"net/http"
	"sync"
	"github.com/matthew1809/cart-tests/request"
	"github.com/matthew1809/cart-tests/models"
)

// TaxFromCartAsync does a thing
func TaxFromCartAsync(baseURL string, client *http.Client, items models.CartItemResponse, token string) {
	var wg sync.WaitGroup

	for i := 0; i < len(items.Data); i++ {
		wg.Add(1)
		for z := 0; z < len(items.Data[i].Relationships.Taxes.Data); z++ {
			
			taxItemID := items.Data[i].Relationships.Taxes.Data[z].ID
			fullURL := baseURL + "/" + items.Data[i].ID + "/taxes/" + taxItemID

			go request.GenericRequest(&wg, i, fullURL, *client, "DELETE",  nil, token, "add.remove.taxFromCartItem", 204)
			wg.Wait()
		}
	}

	wg.Wait()
	fmt.Println("Removing taxes: Completed")
	return
}
