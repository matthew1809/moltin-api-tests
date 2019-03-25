package create

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/matthew1809/cart-tests/models"
	"github.com/matthew1809/cart-tests/request"
)

func Promotion(baseURL string, client *http.Client, token string) models.CreatePromotionResponse {

	fullURL := baseURL + "/promotions"

	createPromotionData := &models.Promotion{
		Type:          "promotion",
		PromotionType: "fixed_discount",
		Name:          "testPromo",
		Description:   "testing the promotions service",
		Enabled:       true,
		Schema: models.Schema{
			Currencies: []models.Currency{
				{
					Currency: "USD",
					Amount:   5000,
				},
			},
		},
		Start: "2017-05-12T15:04:05+00:00",
		End:   "2017-05-12T15:04:05+00:00",
	}

	FullPayload := &models.TopLevelRequest{
		Data: *createPromotionData,
	}

	bytesRepresentation, err := json.Marshal(FullPayload)

	if err != nil {
		fmt.Printf("Error marshalling JSON, failed with %s\n", err)
	}

	res := request.GenericRequest(fullURL, *client, "POST", bytes.NewBuffer(bytesRepresentation), token, "create.Promotion", 201)

	var promotion models.CreatePromotionResponse
	unmarshallErr := json.Unmarshal(res, &promotion)

	if unmarshallErr != nil {
		fmt.Printf("Error reading create promotion response, failed with %s\n", unmarshallErr)
	}

	fmt.Println("Create Promotion: Completed")
	return promotion
}
