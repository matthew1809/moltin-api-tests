package main

import (
	"fmt"
	"net/http"
	"sync"
	"io"
)

// GenericRequest builds and makes an API call to Moltin
func GenericRequest(wg *sync.WaitGroup, ID int, path string, client http.Client, method string, payload io.Reader, token string, name string, expectedCode int) {
	defer wg.Done()

	req, err := http.NewRequest(method, path, payload)
	req.Header.Set("Authorization", token)
	req.Header.Set("content-type", "application/json")

	res, err := client.Do(req)

	if err != nil {
		fmt.Printf("Error making request, failed with %s\n", err)
	}

	res.Body.Close()

	if res.StatusCode != expectedCode {
        fmt.Println("Bad response code from", name, ":", res.StatusCode, http.StatusText(res.StatusCode))
    }
}