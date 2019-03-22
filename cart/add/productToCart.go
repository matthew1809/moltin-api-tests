package add

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"sync"
	"github.com/matthew1809/cart-tests/models"
)

type addToCartRequest struct {
	ID       string `json:"id"`
	Quantity int    `json:"quantity"`
	ItemType string `json:"type"`
}

// ProductToCart adds a single product to a Moltin cart
func ProductToCart(wg *sync.WaitGroup, outerID int, innerID int, baseURL string, client http.Client, productID string, token string) {
	defer wg.Done()

	addToCartData := &addToCartRequest{
		ID:       productID,
		Quantity: 2,
		ItemType: "cart_item",
	}

	FullPayload := &models.TopLevelAddToCartRequest{
		Data: *addToCartData,
	}

	bytesRepresentation, err := json.Marshal(FullPayload)

	if err != nil {
		fmt.Printf("Error marshalling JSON, failed with %s\n", err)
	}

	req, err := http.NewRequest("POST", baseURL, bytes.NewBuffer(bytesRepresentation))
	req.Header.Set("Authorization", token)
	req.Header.Set("content-type", "application/json")

	res, err := client.Do(req)

	if err != nil {
		fmt.Printf("Error adding to cart, failed with %s\n", err)
	}

	if res.StatusCode != 201 {
        fmt.Println("Bad response code from add.ProductToCart:", res.StatusCode, http.StatusText(res.StatusCode))
	}
	
	res.Body.Close()
	// fmt.Println("added two of", productID, "to cart")
	// fmt.Printf("Worker %v %v: Finished\n", outerID, innerID)
}
