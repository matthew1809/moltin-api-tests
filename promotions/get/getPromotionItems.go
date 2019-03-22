package get

import (
	"net/http"
	"github.com/matthew1809/cart-tests/models"
	"github.com/matthew1809/cart-tests/cart/get"
)

// PromotionItems is part of the get package, it will fetch all promotion items in a cart
func PromotionItems(baseURL string, client http.Client, token string, cartID string) []models.CartItem {

	var promotionItems []models.CartItem
	items := get.CartItems(baseURL, client, token, cartID)

	for i := 0; i < len(items.Data); i++ {
		if(items.Data[i].Type == "promotion_item") {
			promotionItems = append(promotionItems, items.Data[i])
		}
	}
	return promotionItems
}