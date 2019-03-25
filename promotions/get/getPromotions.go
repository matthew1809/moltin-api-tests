package get

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/matthew1809/cart-tests/models"
	"github.com/matthew1809/cart-tests/request"
)

// Promotions is part of the get package, it will fetch all promotions for a given Moltin store
func Promotions(baseURL string, client http.Client, token string) models.PromotionsResponse {

	fmt.Println("get.Promotions running")
	fullURL := baseURL + "/promotions"

	res := request.GenericRequest(fullURL, client, "GET", nil, token, "get.promotions", 200)

	var promotions models.PromotionsResponse
	unmarshallErr := json.Unmarshal(res, &promotions)

	if unmarshallErr != nil {
		fmt.Printf("Error reading get promotion response, failed with %s\n", unmarshallErr)
	}

	fmt.Println("Get Promotions: Completed")
	return promotions
}
