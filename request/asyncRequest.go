package request

import (
	"fmt"
	"net/http"
	"sync"
	"io"
	"io/ioutil"
)

// AsyncGenericRequest builds and makes an API call to Moltin, but also takes a waitgroup
func AsyncGenericRequest(wg *sync.WaitGroup, ID int, path string, client http.Client, method string, payload io.Reader, token string, name string, expectedCode int) []byte {
	defer wg.Done()

	req, err := http.NewRequest(method, path, payload)
	req.Header.Set("Authorization", token)
	req.Header.Set("content-type", "application/json")

	res, err := client.Do(req)

	if err != nil {
		fmt.Printf("Error making request, failed with %s\n", err)
	}

	response, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Printf("Error reading response, failed with %s\n", err)
	}

	// res.Body.Close()
	res.Body.Close()

	if res.StatusCode != expectedCode {
		fmt.Println("Bad response code from", name, ":", res.StatusCode, http.StatusText(res.StatusCode))
		fmt.Println(string(response))
	}
	
	return response
}