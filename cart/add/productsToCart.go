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

// ProductsToCart uses a waitgroup and go routines to add a single product to the cart multiple times
func ProductsToCart(baseURL string, client *http.Client, productIDs []string, token string, cartID string) {

	var wg sync.WaitGroup

	fullURL := baseURL + "/carts/" + cartID + "/items"

	for a := 0; a < 4; a++ {
		for i := 0; i < len(productIDs); i++ {
			wg.Add(1)

			addToCartData := &models.AddToCartRequest{
				ID:       productIDs[i],
				Quantity: 2,
				ItemType: "cart_item",
			}
		
			FullPayload := &models.TopLevelRequest{
				Data: *addToCartData,
			}
		
			bytesRepresentation, err := json.Marshal(FullPayload)
			
			if err != nil {
				fmt.Printf("Error marshalling JSON, failed with %s\n", err)
			}
	
			go request.AsyncGenericRequest(&wg, a + i, fullURL, *client, "POST",  bytes.NewBuffer(bytesRepresentation), token, "add.ProductToCart", 201)
		}
		wg.Wait()
	}

	wg.Wait()
	fmt.Println("Add Products: Completed")
}
