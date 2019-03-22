package main

import (
	"fmt"
	"github.com/matthew1809/cart-tests/cart/add"
	"github.com/matthew1809/cart-tests/cart/get"
	"github.com/matthew1809/cart-tests/cart/remove"
	"net/http"
)

func testTaxItems() {

	fmt.Println("Starting the application...")

	variables := NewConfig()
	cartBaseURL := variables.baseURL + "/carts/" + variables.cartID
	cartItemsBaseURL := variables.baseURL + "/carts/" + variables.cartID + "/items"

	client := &http.Client{}
	taxAmounts := []int{}

	get.Cart(cartBaseURL, *client, variables.token)

	cartItems := get.CartItems(cartItemsBaseURL, *client, variables.token)

	remove.AllProductsFromCart(cartItemsBaseURL, client, cartItems, variables.token)

	get.CartItems(cartItemsBaseURL, *client, variables.token)

	add.ProductsToCart(cartItemsBaseURL, client, variables.productIDs, variables.token)

	cartItems2 := get.CartItems(cartItemsBaseURL, *client, variables.token)

	get.Cart(cartBaseURL, *client, variables.token)

	add.TaxesToCart(cartItemsBaseURL, client, cartItems2, variables.token)

	taxAmountAfterFirstAdd := get.Cart(cartBaseURL, *client, variables.token)

	taxAmounts = append(taxAmounts, taxAmountAfterFirstAdd)

	cartItems3 := get.CartItems(cartItemsBaseURL, *client, variables.token)

	remove.TaxFromCartAsync(cartItemsBaseURL, client, cartItems3, variables.token)

	get.Cart(cartBaseURL, *client, variables.token)

	cartItems4 := get.CartItems(cartItemsBaseURL, *client, variables.token)

	add.TaxesToCart(cartItemsBaseURL, client, cartItems4, variables.token)

	taxAmountAfterSecondAdd := get.Cart(cartBaseURL, *client, variables.token)

	taxAmounts = append(taxAmounts, taxAmountAfterSecondAdd)

	fmt.Println(taxAmounts)

	if taxAmounts[0] != taxAmounts[1] {
		fmt.Println("innacurate taxes found")
	}

	fmt.Println("Terminating the application...")
}

func main() {
	for i := 0; i < 25; i++ {
		testTaxItems()
	}
}
