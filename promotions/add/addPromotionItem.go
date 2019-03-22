package add

import (
	"fmt"
	"net/http"
	"github.com/matthew1809/cart-tests/models"
	"github.com/matthew1809/cart-tests/request"
	"encoding/json"
	"bytes"
)

// PromotionItem is part of the add package, it will add a single promo to the cart
func PromotionItem(baseURL string, client *http.Client, token string, cartID string) models.CartItemResponse {
	fullURL := baseURL + "/carts/" + cartID + "/items"

	addPromotionItemToCartData := &models.AddPromotionItem{
		Type:         "promotion_item",
		Code:         "bob",
	}

	FullPayload := &models.TopLevelRequest{
		Data: *addPromotionItemToCartData,
	}

	bytesRepresentation, err := json.Marshal(FullPayload)

	if err != nil {
		fmt.Printf("Error marshalling JSON, failed with %s\n", err)
	}

	res := request.GenericRequest(fullURL, *client, "POST",  bytes.NewBuffer(bytesRepresentation), token, "add.PromotionItemToCart", 201)

	var items models.CartItemResponse
	unmarshallErr := json.Unmarshal(res, &items)

	if unmarshallErr != nil {
		fmt.Printf("Error reading get cart items response, failed with %s\n", unmarshallErr)
	}
	
	fmt.Println("Add Promotion: Completed")
	return items
}