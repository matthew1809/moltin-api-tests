package tests

import (
	"github.com/matthew1809/cart-tests/config"
	cartsGet "github.com/matthew1809/cart-tests/cart/get"
	cartsAdd "github.com/matthew1809/cart-tests/cart/add"
	cartsRemove "github.com/matthew1809/cart-tests/cart/remove"
	promotionsGet "github.com/matthew1809/cart-tests/promotions/get"
	promotionsAdd "github.com/matthew1809/cart-tests/promotions/add"
	//promotionsRemove "github.com/matthew1809/cart-tests/promotions/remove"
	"net/http"
)

// TestPromotions does a thing
func TestPromotions(variables config.Config, client *http.Client) {
	
	cartItems := cartsGet.CartItems(variables.BaseURL, *client, variables.Token, variables.CartID)

	if(len(cartItems.Data) > 0) { 
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