package add

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"sync"

	"github.com/matthew1809/cart-tests/models"
)

// TaxToCartItem adds tax to a single cart item
func TaxToCartItem(wg *sync.WaitGroup, baseURL string, itemID string, client http.Client, token string) {
	// fmt.Println("Adding tax to item", itemID)
	defer wg.Done()

	fullURL := baseURL + "/" + itemID + "/taxes"

	addTaxToCartItemData := &models.AddTaxToCartItemRequest{
		Type:         "tax_item",
		Name:         "taxTest",
		Jurisdiction: "Dublin",
		Code:         "NA",
		Rate:         2,
	}

	FullPayload := &models.TopLevelAddToCartRequest{
		Data: *addTaxToCartItemData,
	}

	bytesRepresentation, err := json.Marshal(FullPayload)

	if err != nil {
		fmt.Printf("Error marshalling JSON, failed with %s\n", err)
	}

	req, err := http.NewRequest("POST", fullURL, bytes.NewBuffer(bytesRepresentation))
	req.Header.Set("Authorization", token)
	req.Header.Set("content-type", "application/json")

	// response, err :=
	client.Do(req)

	// if err != nil {
	// 	fmt.Printf("The HTTP request failed with error %s\n", err)
	// } else {
	// 	data, _ := ioutil.ReadAll(response.Body)c
	// 	fmt.Println("added tax to cart item with ID of", itemID, "and the tax item is:", string(data))
	// }
}
