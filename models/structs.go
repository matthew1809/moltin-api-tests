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
	Type      	  string        `json: "type"`
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

// TopLevelRequest builds the payload for add to cart
type TopLevelRequest struct {
	Data interface{} `json:"data"`
}

type AddToCartRequest struct {
	ID	string `json:"id"`
	Quantity int `json:"quantity"`
	ItemType string `json:"type"`
}

type CheckoutRequest struct {
		Customer `json:"customer"`
		BillingAddress `json:"billing_address"`
		ShippingAddress `json:"shipping_address"`
}

type BillingAddress struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Line1     string `json:"line_1"`
	Line2     string `json:"line_2"`
	Postcode  string `json:"postcode"`
	County    string `json:"county"`
	Country   string `json:"country"`
}

type Customer struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

type ShippingAddress struct {
	FirstName    string `json:"first_name"`
	LastName     string `json:"last_name"`
	PhoneNumber  string `json:"phone_number"`
	Line1        string `json:"line_1"`
	City         string `json:"city"`
	Line2        string `json:"line_2"`
	CompanyName  string `json:"company_name"`
	Postcode     string `json:"postcode"`
	County       string `json:"county"`
	Country      string `json:"country"`
	Instructions string `json:"instructions"`
}

type CheckoutResponse struct {
	Data struct {
		Type     string `json:"type"`
		ID       string `json:"id"`
		Status   string `json:"status"`
		Payment  string `json:"payment"`
		Shipping string `json:"shipping"`
		Customer struct {
			Name  string `json:"name"`
			Email string `json:"email"`
		} `json:"customer"`
		ShippingAddress struct {
			FirstName    string `json:"first_name"`
			LastName     string `json:"last_name"`
			PhoneNumber  string `json:"phone_number"`
			CompanyName  string `json:"company_name"`
			Line1        string `json:"line_1"`
			Line2        string `json:"line_2"`
			City         string `json:"city"`
			Postcode     string `json:"postcode"`
			County       string `json:"county"`
			Country      string `json:"country"`
			Instructions string `json:"instructions"`
		} `json:"shipping_address"`
		BillingAddress struct {
			FirstName   string `json:"first_name"`
			LastName    string `json:"last_name"`
			CompanyName string `json:"company_name"`
			Line1       string `json:"line_1"`
			Line2       string `json:"line_2"`
			City        string `json:"city"`
			Postcode    string `json:"postcode"`
			County      string `json:"county"`
			Country     string `json:"country"`
		} `json:"billing_address"`
		Links struct {
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
		Relationships struct {
			Items struct {
				Data []struct {
					Type string `json:"type"`
					ID   string `json:"id"`
				} `json:"data"`
			} `json:"items"`
		} `json:"relationships"`
	} `json:"data"`
}

type AddPromotionItem struct {
	Type string `json:"type"`
	Code string `json:"code"`
}