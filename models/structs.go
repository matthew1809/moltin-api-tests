package models

import "time"

type AddTaxToCartItemRequest struct {
	Type         string  `json:"type"`
	Name         string  `json:"name"`
	Jurisdiction string  `json:"jurisdiction"`
	Code         string  `json:"code"`
	Rate         float32 `json:"rate"`
}

type CartResponse struct {
	Data struct {
		ID    string `json:"id"`
		Type  string `json:"type"`
		Links struct {
			Self string `json:"self"`
		} `json:"links"`
		Meta struct {
			DisplayPrice struct {
				WithTax struct {
					Amount    int    `json:"amount"`
					Currency  string `json:"currency"`
					Formatted string `json:"formatted"`
				} `json:"with_tax"`
				WithoutTax struct {
					Amount    int    `json:"amount"`
					Currency  string `json:"currency"`
					Formatted string `json:"formatted"`
				} `json:"without_tax"`
				Tax struct {
					Amount    int    `json:"amount"`
					Currency  string `json:"currency"`
					Formatted string `json:"formatted"`
				} `json:"tax"`
			} `json:"display_price"`
			Timestamps struct {
				CreatedAt time.Time `json:"created_at"`
				UpdatedAt time.Time `json:"updated_at"`
			} `json:"timestamps"`
		} `json:"meta"`
	} `json:"data"`
}

type CartItemResponse struct {
	Data []CartItem `json: "data"`
}

type CartItem struct {
	ID            string        `json: "id"`
	ItemType      string        `json: "type"`
	Description   string        `json: "description"`
	Sku           string        `json: "sku"`
	Quantity      int           `json: "quantity"`
	ManageStock   bool          `json: "manage_stock"`
	UnitPrice     Price         `json: "unit_price"`
	Value         Value         `json: "value"`
	Image         Image         `json: "image"`
	Relationships Relationships `json: "relationships"`
}

type Relationships struct {
	Taxes struct {
		Data []struct {
			Type string `json:"type"`
			ID   string `json:"id"`
		} `json:"data"`
	} `json:"taxes"`
}

type Image struct {
	MimeType string `json: "mime_type"`
	FileName string `json: "file_name"`
	Href     string `json: "href"`
}

type Price struct {
	Amount      int    `json: "amount"`
	Currency    string `json: "currency"`
	IncludesTax bool   `json: "includes_tax"`
}

type Value struct {
	Amount      int    `json: "amount"`
	Currency    string `json: "currency"`
	IncludesTax bool   `json: "includes_tax"`
}

// TopLevelAddToCartRequest builds the payload for add to cart
type TopLevelAddToCartRequest struct {
	Data interface{} `json:"data"`
}
