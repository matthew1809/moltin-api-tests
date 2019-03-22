package remove

import (
	"fmt"
	"net/http"
	"github.com/matthew1809/cart-tests/request"
)

// PromotionItems is part of the remove package, it will remove all promotion items in a cart
func PromotionItems(baseURL string, client *http.Client, promoItemID string, token string, cartID string) {
	fullURL := baseURL + "/carts/" + cartID + "/items/" + promoItemID

	request.GenericRequest(fullURL, *client, "DELETE",  nil, token, "add.remove.itemFromCart", 200)
	fmt.Println("Remove Promotion: Completed")
}