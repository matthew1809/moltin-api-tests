package tests

import (
	cartsAdd "github.com/matthew1809/cart-tests/cart/add"
	cartsGet "github.com/matthew1809/cart-tests/cart/get"
	cartsRemove "github.com/matthew1809/cart-tests/cart/remove"
	"github.com/matthew1809/cart-tests/config"
	promotionsAdd "github.com/matthew1809/cart-tests/promotions/add"
	promotionsGet "github.com/matthew1809/cart-tests/promotions/get"

	//promotionsRemove "github.com/matthew1809/cart-tests/promotions/remove"
	"net/http"
)

// TestPromotionItems does a thing
func TestPromotionItems(variables config.Config, client *http.Client) {

	cartItems := cartsGet.CartItems(variables.BaseURL, *client, variables.Token, variables.CartID)

	if len(cartItems.Data) > 0 {
		cartsRemove.AllProductsFromCart(variables.BaseURL, client, cartItems, variables.Token, variables.CartID)
		cartsGet.CartItems(variables.BaseURL, *client, variables.Token, variables.CartID)
	}

	cartsAdd.ProductsToCart(variables.BaseURL, client, variables.ProductIDs, variables.Token, variables.CartID)
	cartsGet.CartItems(variables.BaseURL, *client, variables.Token, variables.CartID)

	promotionsAdd.PromotionItem(variables.BaseURL, client, variables.Token, variables.CartID)

	promotionsGet.PromotionItems(variables.BaseURL, *client, variables.Token, variables.CartID)

	// if(len(promotionItems) > 0) {
	// 	promotionItemID := promotionItems[0].ID
	// 	promotionsRemove.PromotionItems(variables.BaseURL, client, promotionItemID, variables.Token, variables.CartID)
	// }
}
