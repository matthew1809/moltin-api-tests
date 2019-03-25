package tests

import (
	"fmt"
	"net/http"

	cartsAdd "github.com/matthew1809/cart-tests/cart/add"
	cartsGet "github.com/matthew1809/cart-tests/cart/get"
	cartsRemove "github.com/matthew1809/cart-tests/cart/remove"
	"github.com/matthew1809/cart-tests/checkout"
	"github.com/matthew1809/cart-tests/config"
	ordersGet "github.com/matthew1809/cart-tests/orders/get"
	promotionsAdd "github.com/matthew1809/cart-tests/promotions/add"
)

func TestCartWithTaxesAndPromotion(variables config.Config, client *http.Client) {

	cartItems := cartsGet.CartItems(variables.BaseURL, *client, variables.Token, variables.CartID)

	if len(cartItems.Data) > 0 {
		cartsRemove.AllProductsFromCart(variables.BaseURL, client, cartItems, variables.Token, variables.CartID)
		cartsGet.CartItems(variables.BaseURL, *client, variables.Token, variables.CartID)
	}

	cartsAdd.ProductsToCart(variables.BaseURL, client, variables.ProductIDs, variables.Token, variables.CartID)

	cartItems2 := cartsGet.CartItems(variables.BaseURL, *client, variables.Token, variables.CartID)

	withTax, withoutTax, Tax := cartsGet.Tax(cartItems2)
	fmt.Printf("cart items: With tax: %v, without tax: %v, tax: %v \n", withTax, withoutTax, Tax)

	cartBeforeTaxAndPromo := cartsGet.Cart(variables.BaseURL, *client, variables.Token, variables.CartID)
	fmt.Printf("cart total: With tax: %v, without tax: %v, tax: %v \n\n", cartBeforeTaxAndPromo.Data.Meta.DisplayPrice.WithTax.Amount, cartBeforeTaxAndPromo.Data.Meta.DisplayPrice.WithoutTax.Amount, cartBeforeTaxAndPromo.Data.Meta.DisplayPrice.Tax.Amount)

	cartsAdd.TaxesToCart(variables.BaseURL, client, cartItems2, variables.Token, variables.CartID)

	cartItems3 := cartsGet.CartItems(variables.BaseURL, *client, variables.Token, variables.CartID)

	withTax, withoutTax, Tax = cartsGet.Tax(cartItems3)
	fmt.Printf("cart items: With tax: %v, without tax: %v, tax: %v \n", withTax, withoutTax, Tax)

	cartAfterTaxBeforePromo := cartsGet.Cart(variables.BaseURL, *client, variables.Token, variables.CartID)
	fmt.Printf("cart total: With tax: %v, without tax: %v, tax: %v \n\n", cartAfterTaxBeforePromo.Data.Meta.DisplayPrice.WithTax.Amount, cartAfterTaxBeforePromo.Data.Meta.DisplayPrice.WithoutTax.Amount, cartBeforeTaxAndPromo.Data.Meta.DisplayPrice.Tax.Amount)

	promotionsAdd.PromotionItem(variables.BaseURL, client, variables.Token, variables.CartID)

	cartItems4 := cartsGet.CartItems(variables.BaseURL, *client, variables.Token, variables.CartID)

	withTax, withoutTax, Tax = cartsGet.Tax(cartItems4)
	fmt.Printf("cart items: With tax: %v, without tax: %v, tax: %v \n", withTax, withoutTax, Tax)

	cartAfterTaxAndPromo := cartsGet.Cart(variables.BaseURL, *client, variables.Token, variables.CartID)
	fmt.Printf("cart total: With tax: %v, without tax: %v, tax: %v \n\n", cartAfterTaxAndPromo.Data.Meta.DisplayPrice.WithTax.Amount, cartAfterTaxAndPromo.Data.Meta.DisplayPrice.WithoutTax.Amount, cartBeforeTaxAndPromo.Data.Meta.DisplayPrice.Tax.Amount)

	fmt.Println("cart tax amount before checkout:", cartAfterTaxAndPromo.Data.Meta.DisplayPrice.Tax.Amount)

	order := checkout.Checkout(variables.BaseURL, client, variables.Token, variables.CartID)

	fmt.Println("order tax amount after checkout:", order.Data.Meta.DisplayPrice.Tax.Amount)

	orderItems := ordersGet.OrderItems(variables.BaseURL, *client, variables.Token, order.Data.ID)

	withTax, withoutTax, Tax = ordersGet.Tax(orderItems)
	fmt.Printf("\norder items: With tax: %v, without tax: %v, tax: %v \n", withTax, withoutTax, Tax)

}
