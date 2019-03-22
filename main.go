package main

import (
	"github.com/matthew1809/cart-tests/tests"
	"github.com/matthew1809/cart-tests/config"
	"net/http"
)

func main() {
	variables := config.NewConfig()
	client := &http.Client{}

	tests.TestPromotions(variables, client)
	// tests.RunTaxItemsTest(variables)
	// tests.TestCheckout(variables, client)
}
