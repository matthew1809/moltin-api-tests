package remove

import (
	"fmt"
	"net/http"

	"github.com/matthew1809/cart-tests/request"
)

// Promotion is part of the remove package, it will remove all promotion items in a cart
func Promotion(baseURL string, client *http.Client, token string, promotionID string) {
	fullURL := baseURL + "/promotions/" + promotionID

	request.GenericRequest(fullURL, *client, "DELETE", nil, token, "remove.promotion", 204)
	fmt.Println("Remove Promotion: Completed")
}
