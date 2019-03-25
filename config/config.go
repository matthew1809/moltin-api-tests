package config

import (
	"github.com/matthew1809/cart-tests/authentication"
	uuid "github.com/satori/go.uuid"
)

// Config contains the env variable structure for the project
type Config struct {
	BaseURL    string
	Token      string
	CartID     string
	ProductIDs []string
	Runs       int
}

// NewConfig returns an object containing the required env variables for the application
func NewConfig(clientID string, clientSecret string) Config {
	token := authentication.Authenticate(clientID, clientSecret)
	cartUUID := uuid.Must(uuid.NewV4()).String()

	config := Config{
		BaseURL:    "https://api.moltin.com/v2",
		Token:      token,
		CartID:     cartUUID,
		ProductIDs: []string{"ac5fd15e-e822-4d19-9457-bf6c5a83ac11", "f6b139b3-7363-4b72-9f0b-9a72e6627607"},
		Runs:       25,
	}

	return config
}
