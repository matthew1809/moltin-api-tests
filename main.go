package main

import (
	"bufio"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/matthew1809/cart-tests/models"

	"github.com/fatih/color"
	"github.com/matthew1809/cart-tests/config"
	"github.com/matthew1809/cart-tests/tests"
)

func runTest(test string, clientID string, clientSecret string) {
	variables := config.NewConfig(clientID, clientSecret)
	client := &http.Client{}

	if variables.Token == "0" {
		fmt.Println("Cannot proceed without authenticating")
		return
	}

	switch test {
	case "TestPromotions":
		tests.TestPromotions(variables, client)
	case "TestTaxItems":
		tests.TestTaxItems(variables, client)
	case "TestCheckout":
		tests.TestCheckout(variables, client)
	case "TestOrders":
		tests.TestOrders(variables, client)
	case "TestCartWithTaxesAndPromotion":
		tests.TestCartWithTaxesAndPromotion(variables, client)
	default:
		fmt.Println("You have not entered a valid test")
	}

}

func readAuthFile() models.Tokens {
	file, _ := ioutil.ReadFile("auth.json")
	authData := models.Tokens{}
	_ = json.Unmarshal([]byte(file), &authData)

	return authData
}

func main() {

	test := flag.String("test", "", "a valid test name")
	flag.Parse()

	if *test == "" {
		color.Cyan("Welcome to the API testing facility. You can run any of the following tests by running the program with a `-test=` flag followed by the test name:")

		files, err := ioutil.ReadDir("./tests")
		if err != nil {
			log.Fatal(err)
		}

		for _, f := range files {
			name := strings.Split(f.Name(), ".")
			fmt.Println(name[0])
		}
	} else {

		authData := readAuthFile()

		if (models.Tokens{}) != authData {
			runTest(*test, authData.ClientID, authData.ClientSecret)
		} else {
			reader := bufio.NewReader(os.Stdin)
			fmt.Print("Enter client ID: ")
			collectedClientID, _ := reader.ReadString('\n')
			trimmedClientID := strings.TrimRight(collectedClientID, "\r\n")

			fmt.Print("Enter client secret: ")
			collectedClientSecret, _ := reader.ReadString('\n')
			trimmedClientSecret := strings.TrimRight(collectedClientSecret, "\r\n")

			tokens := models.Tokens{ClientID: trimmedClientID, ClientSecret: trimmedClientSecret}
			tokensJSON, _ := json.Marshal(tokens)
			err := ioutil.WriteFile("auth.json", tokensJSON, 0644)

			if err != nil {
				fmt.Println(err)
			}

			runTest(*test, trimmedClientID, trimmedClientSecret)

		}
	}

}
