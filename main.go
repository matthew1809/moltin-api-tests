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

	client := &http.Client{}
	taxAmounts := map[string]int{}

	get.Cart(variables.baseURL, *client, variables.token, variables.cartID)

	cartItems := get.CartItems(variables.baseURL, *client, variables.token, variables.cartID)

	if(len(cartItems.Data) > 0) { 
		remove.AllProductsFromCart(variables.baseURL, client, cartItems, variables.token, variables.cartID)
		get.CartItems(variables.baseURL, *client, variables.token, variables.cartID)
	}

	add.ProductsToCart(variables.baseURL, client, variables.productIDs, variables.token, variables.cartID)

	cartItems2 := get.CartItems(variables.baseURL, *client, variables.token, variables.cartID)

	get.Cart(variables.baseURL, *client, variables.token, variables.cartID)

	add.TaxesToCart(variables.baseURL, client, cartItems2, variables.token, variables.cartID)

	taxAmountAfterFirstAdd := get.Cart(variables.baseURL, *client, variables.token, variables.cartID)

	taxAmounts["first"] = taxAmountAfterFirstAdd

	cartItems3 := get.CartItems(variables.baseURL, *client, variables.token, variables.cartID)

	remove.TaxFromCartAsync(variables.baseURL, client, cartItems3, variables.token, variables.cartID)

	get.Cart(variables.baseURL, *client, variables.token, variables.cartID)

	cartItems4 := get.CartItems(variables.baseURL, *client, variables.token, variables.cartID)

	add.TaxesToCart(variables.baseURL, client, cartItems4, variables.token, variables.cartID)

	taxAmountAfterSecondAdd := get.Cart(variables.baseURL, *client, variables.token, variables.cartID)

	taxAmounts["second"] = taxAmountAfterSecondAdd

	fmt.Println(taxAmounts)

	if taxAmounts["first-taxes"] != taxAmounts["second-taxes"] {
		return false
	}

	fmt.Println("Terminating the application...")
	
	return true
}

func main() {
	variables := NewConfig()

	results := map[string]int{"accurate": 0, "inaccurate": 0}

	for i := 0; i < variables.runs; i++ {
		if(testTaxItems(variables) == true) {
			results["accurate"]++
		} else {
			results["innacurate"]++
		}
	}

	fmt.Println(results)

}
