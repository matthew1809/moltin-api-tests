package get

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/matthew1809/cart-tests/models"
)

// CartItems fetches cart items
func CartItems(baseURL string, client http.Client, token string) models.CartItemResponse {

	// fmt.Println("Fetching cart items")
	req, _ := http.NewRequest("GET", baseURL, nil)
	req.Header.Set("Authorization", token)
	req.Header.Set("content-type", "application/json")
	res, err := client.Do(req)

	if err != nil {
		fmt.Printf("Error making get cart items request, failed with %s\n", err)
	}

	response, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Printf("Error reading get cart response, failed with %s\n", err)
	}

	res.Body.Close()

	var items models.CartItemResponse
	unmarshallErr := json.Unmarshal(response, &items)

	if unmarshallErr != nil {
		fmt.Printf("Error reading get cart response, failed with %s\n", unmarshallErr)
	} else {
		// fmt.Println("quantity of item: ", items.Data[0].Quantity)
		// fmt.Println("quantity of item: ", items.Data[1].Quantity)
		fmt.Println("Number of unique cart items: ", len(items.Data))
	}

	return items
}
