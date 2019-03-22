package remove

import (
	"fmt"
	"net/http"
	"sync"

	"github.com/matthew1809/cart-tests/models"
)

// TaxItemFromCartItem removes tax from a given cart item ID
func TaxItemFromCartItem(wg *sync.WaitGroup, baseURL string, item models.CartItem, client http.Client, token string) {
	// fmt.Println("Removing tax from item", item.ID)
	defer wg.Done()

	for i := 0; i < len(item.Relationships.Taxes.Data); i++ {
		taxItemID := item.Relationships.Taxes.Data[i].ID

		fullURL := baseURL + "/" + item.ID + "/taxes/" + taxItemID

		req, err := http.NewRequest("DELETE", fullURL, nil)
		req.Header.Set("Authorization", token)
		req.Header.Set("content-type", "application/json")

		if err != nil {
			fmt.Printf("The HTTP request failed with error %s\n", err)
		}

		res, err := client.Do(req)

		if err != nil {
			fmt.Printf("The HTTP request failed with error %s\n", err)
		}

		if res.StatusCode != 204 {
			fmt.Println("Bad response code from remove.TaxItemFromCartItem:", res.StatusCode, http.StatusText(res.StatusCode))
		}

		// fmt.Println("deleted tax item", item.ID)
	}
}
