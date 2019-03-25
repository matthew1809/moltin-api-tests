package get

import (
	"github.com/matthew1809/cart-tests/models"
)

// Tax prints the tax item amount on each order item
func Tax(items models.OrderItemResponse) (int, int, int) {

	withTax := 0
	withoutTax := 0
	tax := 0

	for i := 0; i < len(items.Data); i++ {
		withTax = withTax + items.Data[i].Meta.DisplayPrice.WithTax.Value.Amount
		withoutTax = withoutTax + items.Data[i].Meta.DisplayPrice.WithoutTax.Value.Amount
		tax = tax + items.Data[i].Meta.DisplayPrice.Tax.Value.Amount

		// fmt.Println("value of withTax", items.Data[i].Meta.DisplayPrice.WithTax.Value.Amount)
		// fmt.Println("value of withoutTax", items.Data[i].Meta.DisplayPrice.WithoutTax.Value.Amount)
		// fmt.Println("value of tax", items.Data[i].Meta.DisplayPrice.Tax.Value.Amount)
	}

	return withTax, withoutTax, tax
}
