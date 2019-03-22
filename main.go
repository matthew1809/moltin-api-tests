package main

import (
	"github.com/matthew1809/cart-tests/tests"
	"github.com/matthew1809/cart-tests/config"
	"net/http"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"strings"
	"github.com/fatih/color"
)

func main() {

	test := flag.String("test", "", "a valid test name")
	flag.Parse()

	if(*test == "") {
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
		variables := config.NewConfig()
		client := &http.Client{}

		switch *test {
			case "TestPromotions": tests.TestPromotions(variables, client)
			case "TestTaxItems": tests.TestTaxItems(variables, client)
			case "TestCheckout": tests.TestCheckout(variables, client)
			case "TestCartWithTaxesAndPromotion": tests.TestCartWithTaxesAndPromotion(variables, client)
			default: fmt.Println("You have not entered a valid test")
		}
	}
}
