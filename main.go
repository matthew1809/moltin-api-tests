package main

import (
	// "github.com/matthew1809/cart-tests/tests"
	// "github.com/matthew1809/cart-tests/config"
	// "net/http"
	// "os"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"strings"
	"github.com/fatih/color"
)

func main() {
	// variables := config.NewConfig()
	// client := &http.Client{}

	// tests.TestPromotions(variables, client)
	// tests.RunTaxItemsTest(variables)
	// tests.TestCheckout(variables, client)
	// tests.TestCartWithTaxesAndPromotion(variables, client)

	flag.String("test", "checkoutACart", "a valid test name")
	flag.Parse()

	// fmt.Println("test:", *wordPtr)

	color.Cyan("Welcome to the API testing facility. You can run any of the following tests by running the program with a `-test=` flag followed by the test name:")
	// switch argsWithoutProg := os.Args[1:]; argsWithoutProg {
	// 	case
	// }

	files, err := ioutil.ReadDir("./tests")
    if err != nil {
        log.Fatal(err)
    }

    for _, f := range files {
		name := strings.Split(f.Name(), ".")
        fmt.Println(name[0])
    }
}
