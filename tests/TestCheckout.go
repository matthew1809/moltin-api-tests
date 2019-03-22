package tests

import (
	"github.com/matthew1809/cart-tests/checkout"
	"github.com/matthew1809/cart-tests/config"
	"github.com/matthew1809/cart-tests/cart/add"
	"github.com/matthew1809/cart-tests/cart/get"
	"net/http"
	"fmt"
)

// TestCheckout does a thing
func TestCheckout(variables config.Config, client *http.Client) {
	
	cartItems := get.CartItems(variables.BaseURL, *client, variables.Token, variables.CartID)

	if(len(cartItems.Data) < 1) {
		add.ProductsToCart(variables.BaseURL, client, variables.ProductIDs, variables.Token, variables.CartID)
	}

	order := checkout.Checkout(variables.BaseURL, client, variables.Token, variables.CartID)
	fmt.Println(order)
}