package main

import (
	"fmt"
	"github.com/matthew1809/cart-tests/cart/add"
	"github.com/matthew1809/cart-tests/cart/get"
	"github.com/matthew1809/cart-tests/cart/remove"
	"net/http"
)

func testTaxItems(variables Config) bool {

	fmt.Println("Starting the application...")

	cartBaseURL := variables.baseURL + "/carts/" + variables.cartID
	cartItemsBaseURL := variables.baseURL + "/carts/" + variables.cartID + "/items"

	client := &http.Client{}
	taxAmounts := map[string]int{}

	get.Cart(cartBaseURL, *client, variables.token)

	cartItems := get.CartItems(cartItemsBaseURL, *client, variables.token)

	remove.AllProductsFromCart(cartItemsBaseURL, client, cartItems, variables.token)

	get.CartItems(cartItemsBaseURL, *client, variables.token)

	add.ProductsToCart(cartItemsBaseURL, client, variables.productIDs, variables.token)

	cartItems2 := get.CartItems(cartItemsBaseURL, *client, variables.token)

	get.Cart(cartBaseURL, *client, variables.token)

	add.TaxesToCart(cartItemsBaseURL, client, cartItems2, variables.token)

	taxAmountAfterFirstAdd := get.Cart(cartBaseURL, *client, variables.token)

	taxAmounts["first"] = taxAmountAfterFirstAdd

	cartItems3 := get.CartItems(cartItemsBaseURL, *client, variables.token)

	remove.TaxFromCartAsync(cartItemsBaseURL, client, cartItems3, variables.token)

	get.Cart(cartBaseURL, *client, variables.token)

	cartItems4 := get.CartItems(cartItemsBaseURL, *client, variables.token)

	add.TaxesToCart(cartItemsBaseURL, client, cartItems4, variables.token)

	taxAmountAfterSecondAdd := get.Cart(cartBaseURL, *client, variables.token)

	taxAmounts["second"] = taxAmountAfterSecondAdd

	fmt.Println(taxAmounts)

	if taxAmounts["first"] != taxAmounts["second"] {
		return false
	}

	fmt.Println("Terminating the application...")
	
	return true
}

func main() {
	variables := NewConfig()

	results := map[string]int{"accurate": 0, "innacurate": 0}

	for i := 0; i < variables.runs; i++ {
		if(testTaxItems(variables) == true) {
			results["accurate"]++
		} else {
			results["innacurate"]++
		}
	}

	fmt.Println(results)

}
