package tests

import (
	"github.com/matthew1809/cart-tests/orders/get"
	"github.com/matthew1809/cart-tests/config"
	"net/http"
	"fmt"
)

func TestOrders(variables config.Config, client *http.Client) {
	orders := get.Orders(variables.BaseURL, *client, variables.Token)
	fmt.Println(orders)
	// get.Order(variables.BaseURL, *client, variables.Token, "f6431294-be6c-4580-a5bc-2a6b2102d94a")
}