package get

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/matthew1809/cart-tests/models"
)

// Cart fetches cart
func Cart(baseURL string, client http.Client, token string) int {

	req, _ := http.NewRequest("GET", baseURL, nil)
	req.Header.Set("Authorization", token)
	req.Header.Set("content-type", "application/json")
	res, err := client.Do(req)

	if err != nil {
		fmt.Printf("Error making get cart request, failed with %s\n", err)
	}

	response, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Printf("Error reading get cart response, failed with %s\n", err)
	}

	res.Body.Close()

    if res.StatusCode != 200 {
        fmt.Println("Bad response code from get.Cart:", res.StatusCode, http.StatusText(res.StatusCode))
    }

	var cart models.CartResponse
	unmarshallErr := json.Unmarshal(response, &cart)

	if unmarshallErr != nil {
		fmt.Printf("Error reading get cart response, failed with %s\n", unmarshallErr)
	} else {
		fmt.Println("Tax amount is:", cart.Data.Meta.DisplayPrice.Tax.Amount)
	}

	return cart.Data.Meta.DisplayPrice.Tax.Amount
}
