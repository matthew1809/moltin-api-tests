package main

import "github.com/satori/go.uuid"

// Config contains the env variable structure for the project
type Config struct {
	baseURL string
	token string
	cartID string
	productIDs []string
	runs int
}

// NewConfig returns an object containing the required env variables for the application
func NewConfig() Config {
	
	cartUUID := uuid.Must(uuid.NewV4()).String()

	config := Config{
		baseURL: "https://api.moltin.com/v2",
		token: "669bdba38c096b5894da27476afd741fa0d59838",
		cartID: cartUUID,
		productIDs: []string{"0542b5e4-cbb6-4960-bdb4-906c68512ed3", "cca740cf-0dab-4fb2-a71c-0391c48b37d0"},
		runs: 25,
	} 

	return config
}
