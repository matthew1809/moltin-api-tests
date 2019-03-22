package tests

import (
	"fmt"
	"github.com/matthew1809/cart-tests/cart/add"
	"github.com/matthew1809/cart-tests/cart/get"
	"github.com/matthew1809/cart-tests/cart/remove"
	"github.com/matthew1809/cart-tests/config"
	"net/http"
)

// RunTaxItemsTest runs the test for tax items
func RunTaxItemsTest(variables config.Config) {
	results := map[string]int{"accurate": 0, "inaccurate": 0}

	for i := 0; i < variables.Runs; i++ {
		if(TestTaxItems(variables) == true) {
			results["accurate"]++
		} else {
			results["innacurate"]++
		}
	}
	fmt.Println(results)
}

// TestTaxItems does a thing
func TestTaxItems(variables config.Config) bool {

	fmt.Println("Starting the application...")

	client := &http.Client{}
	taxAmounts := map[string]int{}

	get.Cart(variables.BaseURL, *client, variables.Token, variables.CartID)

	cartItems := get.CartItems(variables.BaseURL, *client, variables.Token, variables.CartID)

	if(len(cartItems.Data) > 0) { 
		remove.AllProductsFromCart(variables.BaseURL, client, cartItems, variables.Token, variables.CartID)
		get.CartItems(variables.BaseURL, *client, variables.Token, variables.CartID)
	}

	add.ProductsToCart(variables.BaseURL, client, variables.ProductIDs, variables.Token, variables.CartID)

	cartItems2 := get.CartItems(variables.BaseURL, *client, variables.Token, variables.CartID)

	get.Cart(variables.BaseURL, *client, variables.Token, variables.CartID)

	add.TaxesToCart(variables.BaseURL, client, cartItems2, variables.Token, variables.CartID)

	taxAmountAfterFirstAdd := get.Cart(variables.BaseURL, *client, variables.Token, variables.CartID)

	taxAmounts["first"] = taxAmountAfterFirstAdd

	cartItems3 := get.CartItems(variables.BaseURL, *client, variables.Token, variables.CartID)

	remove.TaxFromCartAsync(variables.BaseURL, client, cartItems3, variables.Token, variables.CartID)

	get.Cart(variables.BaseURL, *client, variables.Token, variables.CartID)

	cartItems4 := get.CartItems(variables.BaseURL, *client, variables.Token, variables.CartID)

	add.TaxesToCart(variables.BaseURL, client, cartItems4, variables.Token, variables.CartID)

	taxAmountAfterSecondAdd := get.Cart(variables.BaseURL, *client, variables.Token, variables.CartID)

	taxAmounts["second"] = taxAmountAfterSecondAdd

	fmt.Println(taxAmounts)

	if taxAmounts["first-taxes"] != taxAmounts["second-taxes"] {
		return false
	}

	fmt.Println("Terminating the application...")
	
	return true
}