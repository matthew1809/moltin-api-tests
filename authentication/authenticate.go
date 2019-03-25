package authentication

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/matthew1809/cart-tests/models"
)

// Authenticate returns a bearer token from the API
func Authenticate(clientID string, clientSecret string) string {

	values := url.Values{"client_id": {clientID}, "client_secret": {clientSecret}, "grant_type": {"client_credentials"}}

	res, err := http.PostForm("https://api.moltin.com/oauth/token", values)
	if err != nil {
		fmt.Printf("Error making request, failed with %s\n", err)
	}

	response, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Printf("Error reading response, failed with %s\n", err)
	}

	var authResponse models.AuthenticationResponse
	unmarshallErr := json.Unmarshal(response, &authResponse)

	if unmarshallErr != nil {
		fmt.Printf("Error reading get cart items response, failed with %s\n", unmarshallErr)
	}

	res.Body.Close()

	if res.StatusCode != 200 {
		fmt.Println("Authentication failed:", string(response))
		return "0"
	}

	return authResponse.AccessToken
}
