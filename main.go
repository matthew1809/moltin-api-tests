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

	cartBaseURL := "https://api.moltin.com/v2/carts/test38"
	cartItemsBaseURL := "https://api.moltin.com/v2/carts/test38/items"
	token := "2701731ffae55867226a6318c4f941378623389f"
	client := &http.Client{}

	taxAmounts := []int{}

	get.Cart(cartBaseURL, *client, token)

	cartItems := get.CartItems(cartItemsBaseURL, *client, token)

	remove.ProductsFromCartAsync(cartItemsBaseURL, client, cartItems, token)

	get.CartItems(cartItemsBaseURL, *client, token)

	add.ProductsToCart(cartItemsBaseURL, client, token)

	cartItems2 := get.CartItems(cartItemsBaseURL, *client, token)

	get.Cart(cartBaseURL, *client, token)

	add.TaxesToCart(cartItemsBaseURL, client, cartItems2, token)

	taxAmountAfterFirstAdd := get.Cart(cartBaseURL, *client, token)

	taxAmounts = append(taxAmounts, taxAmountAfterFirstAdd)

	cartItems3 := get.CartItems(cartItemsBaseURL, *client, token)

	remove.TaxFromCartAsync(cartItemsBaseURL, client, cartItems3, token)

	get.Cart(cartBaseURL, *client, token)

	cartItems4 := get.CartItems(cartItemsBaseURL, *client, token)

	add.TaxesToCart(cartItemsBaseURL, client, cartItems4, token)

	taxAmountAfterSecondAdd := get.Cart(cartBaseURL, *client, token)

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
