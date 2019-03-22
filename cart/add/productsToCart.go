package add

import (
	"fmt"
	"net/http"
	"sync"
)

// ProductsToCart uses a waitgroup and go routines to add a single product to the cart multiple times
func ProductsToCart(baseURL string, client *http.Client, token string) {

	productIDs := []string{"0542b5e4-cbb6-4960-bdb4-906c68512ed3", "cca740cf-0dab-4fb2-a71c-0391c48b37d0"}
	var wg sync.WaitGroup

	for a := 0; a < 4; a++ {
		for i := 0; i < len(productIDs); i++ {
			// fmt.Println("Main: Starting worker", a, i)
			wg.Add(1)
			go ProductToCart(&wg, a, i, baseURL, *client, productIDs[i], token)
		}
		wg.Wait()
	}

	// fmt.Println("Main: Waiting for workers to finish")
	wg.Wait()
	fmt.Println("Add Products: Completed")
}
