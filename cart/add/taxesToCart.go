package add

import (
	"fmt"
	"net/http"
	"sync"
	"encoding/json"
	"bytes"
	"github.com/matthew1809/cart-tests/request"
	"github.com/matthew1809/cart-tests/models"
)

// TaxesToCart uses a waitgroup and go routines to add a single product to the cart multiple times
func TaxesToCart(baseURL string, client *http.Client, items models.CartItemResponse, token string, cartID string) {
		// fmt.Println("Adding taxes")
		var wg sync.WaitGroup

		for i := 0; i < len(items.Data); i++ {

		fullURL := baseURL + "/carts/" + cartID + "/items/" + items.Data[i].ID + "/taxes"

		addTaxToCartItemData := &models.AddTaxToCartItemRequest{
			Type:         "tax_item",
			Name:         "taxTest",
			Jurisdiction: "Dublin",
			Code:         "NA",
			Rate:         2,
		}

		FullPayload := &models.TopLevelRequest{
			Data: *addTaxToCartItemData,
		}

		bytesRepresentation, err := json.Marshal(FullPayload)

		if err != nil {
			fmt.Printf("Error marshalling JSON, failed with %s\n", err)
		}


		wg.Add(1)
		go request.AsyncGenericRequest(&wg, i, fullURL, *client, "POST",  bytes.NewBuffer(bytesRepresentation), token, "add.taxToCartItem", 201)
	}

	// fmt.Println("Main: Waiting for workers to finish")
	wg.Wait()
	fmt.Println("Adding taxes: Completed")
	return
}
