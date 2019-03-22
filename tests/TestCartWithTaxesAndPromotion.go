package tests

import (
	"net/http"
	"github.com/matthew1809/cart-tests/config"
	"github.com/matthew1809/cart-tests/checkout"
	cartsGet "github.com/matthew1809/cart-tests/cart/get"
	cartsAdd "github.com/matthew1809/cart-tests/cart/add"
	cartsRemove "github.com/matthew1809/cart-tests/cart/remove"
	promotionsAdd "github.com/matthew1809/cart-tests/promotions/add"
	"fmt"
)

func TestCartWithTaxesAndPromotion(variables config.Config, client *http.Client) {
	
	cartItems := cartsGet.CartItems(variables.BaseURL, *client, variables.Token, variables.CartID)

	if(len(cartItems.Data) > 0) { 
		cartsRemove.AllProductsFromCart(variables.BaseURL, client, cartItems, variables.Token, variables.CartID)
		cartsGet.CartItems(variables.BaseURL, *client, variables.Token, variables.CartID)
	}

	cartsAdd.ProductsToCart(variables.BaseURL, client, variables.ProductIDs, variables.Token, variables.CartID)
	
	cartItems2 := cartsGet.CartItems(variables.BaseURL, *client, variables.Token, variables.CartID)

	cartsAdd.TaxesToCart(variables.BaseURL, client, cartItems2, variables.Token, variables.CartID)

	promotionsAdd.PromotionItem(variables.BaseURL, client, variables.Token, variables.CartID)
	
	cartsGet.CartItems(variables.BaseURL, *client, variables.Token, variables.CartID)

	cartTaxAmount := cartsGet.Cart(variables.BaseURL, *client, variables.Token, variables.CartID)

	fmt.Println("cart tax amount before checkout:", cartTaxAmount)

	order := checkout.Checkout(variables.BaseURL, client, variables.Token, variables.CartID)

	fmt.Println("order tax amount after checkout:", order.Data.Meta.DisplayPrice.Tax.Amount)
}