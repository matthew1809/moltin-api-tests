package tests

import (
	"fmt"
	"net/http"

	"github.com/matthew1809/cart-tests/promotions/remove"

	"github.com/matthew1809/cart-tests/config"
	"github.com/matthew1809/cart-tests/promotions/create"
	"github.com/matthew1809/cart-tests/promotions/get"
)

// TestPromotions does a thing
func TestPromotions(variables config.Config, client *http.Client) {
	fmt.Println("TestPromotions running")
	get.Promotions(variables.BaseURL, *client, variables.Token)
	promotion := create.Promotion(variables.BaseURL, client, variables.Token)
	remove.Promotion(variables.BaseURL, client, variables.Token, promotion.Data.ID)
}
