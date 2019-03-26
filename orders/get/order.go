package get

import (
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"
	"github.com/matthew1809/cart-tests/models"
	"github.com/matthew1809/cart-tests/request"
)

// Order fetches a single moltin order by ID
func Order(baseURL string, client http.Client, token string, orderID string) models.OrderResponse {

	fullURL := baseURL + "/orders/" + orderID
	res := request.GenericRequest(fullURL, client, "GET", nil, token, "get.Order", 200)

	var order models.OrderResponse
	unmarshallErr := json.Unmarshal(res, &order)

	if unmarshallErr != nil {
		fmt.Printf("Error reading get cart response, failed with %s\n", unmarshallErr)
	}

	if err := json.Unmarshal(res, &order.Flows); err != nil {
        panic(err)
	}
	
	fields := reflect.TypeOf(order.Data)
	num := fields.NumField()

	for i := 0; i < num; i++ {
		field := fields.Field(i)
		f1, _ := fields.FieldByName(field.Name)
		t := f1.Tag.Get("json")

		delete(order.Flows.Data, t)
	}
	
	fmt.Println("flow data", order.Flows)
	fmt.Println("order total without tax", order.Data.Meta.DisplayPrice.WithoutTax.Amount)
	fmt.Println("order total with tax", order.Data.Meta.DisplayPrice.WithTax.Amount)

	return order
}
