package add

import (
	"fmt"
	"net/http"
	"sync"
)

// ProductsToCart uses a waitgroup and go routines to add a single product to the cart multiple times
func ProductsToCart(baseURL string, client *http.Client, productIDs []string, token string) {

	var wg sync.WaitGroup

	for a := 0; a < 4; a++ {
		for i := 0; i < len(productIDs); i++ {
			wg.Add(1)
			go ProductToCart(&wg, a, i, baseURL, *client, productIDs[i], token)
		}
		wg.Wait()
	}

	wg.Wait()
	fmt.Println("Add Products: Completed")
}
