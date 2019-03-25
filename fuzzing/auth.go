package fuzzing

import (
	fuzz "github.com/google/gofuzz"
	"github.com/matthew1809/cart-tests/authentication"
)

// Auth generates random inputs for a Moltin authentication request
func Auth() {
	randStrings()
}

func randStrings() {
	f := fuzz.New()
	var myString string
	f.Fuzz(&myString)

	authentication.Authenticate(myString, myString)
}
