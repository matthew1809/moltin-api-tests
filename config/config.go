package config

import "github.com/satori/go.uuid"

// Config contains the env variable structure for the project
type Config struct {
	BaseURL string
	Token string
	CartID string
	ProductIDs []string
	Runs int
}

// NewConfig returns an object containing the required env variables for the application
func NewConfig() Config {
	
	cartUUID := uuid.Must(uuid.NewV4()).String()

	config := Config{
		BaseURL: "https://api.moltin.com/v2",
		Token: "9fa54a17e5f396a23326dd3560c6ad6976e6ebee",
		CartID: cartUUID,
		ProductIDs: []string{"0542b5e4-cbb6-4960-bdb4-906c68512ed3", "cca740cf-0dab-4fb2-a71c-0391c48b37d0"},
		Runs: 25,
	} 

	return config
}
