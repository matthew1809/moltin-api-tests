package authentication

import (
	"fmt"
	"net/http"
	"net/url"
	"io/ioutil"
	"encoding/json"
	"github.com/matthew1809/cart-tests/models"
)

func Authenticate() string {
	res, err := http.PostForm("https://api.moltin.com/oauth/token",
	url.Values{"client_id": {"9rDK38TcVqmvDzI7xtmRt44hQ6XZ0OWpMMYxxKWCdK"}, "client_secret": {"XYZ"}, "grant_type": {"client_credentials"}})


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

	return authResponse.AccessToken
}