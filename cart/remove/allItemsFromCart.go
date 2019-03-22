package remove

import (
	"fmt"
	"net/http"
	"sync"
	"github.com/matthew1809/cart-tests/request"
	"github.com/matthew1809/cart-tests/models"
)

// AllProductsFromCart uses a waitgroup and go routines to remove all products from the cart
func AllProductsFromCart(baseURL string, client *http.Client, items models.CartItemResponse, token string, cartID string) {

	var wg sync.WaitGroup

	for i := 0; i < len(items.Data); i++ {
		wg.Add(1)

		fullURL := baseURL + "/carts/" + cartID + "/items/" + items.Data[i].ID

		go request.AsyncGenericRequest(&wg, i, fullURL, *client, "DELETE",  nil, token, "add.remove.itemFromCart", 200)
	}

	wg.Wait()
	fmt.Println("Removing products: Completed")
}
