package get

import (
	"github.com/matthew1809/cart-tests/models"
)

// Tax prints the tax amount on each cart item
func Tax(items models.CartItemResponse) (int, int, int) {

	withTax := 0
	withoutTax := 0
	tax := 0

	for i := 0; i < len(items.Data); i++ {
		withTax = withTax + items.Data[i].Meta.DisplayPrice.WithTax.Value.Amount
		withoutTax = withoutTax + items.Data[i].Meta.DisplayPrice.WithoutTax.Value.Amount
		tax = tax + items.Data[i].Meta.DisplayPrice.Tax.Value.Amount
	}

	return withTax, withoutTax, tax
}
