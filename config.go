package main

import "github.com/satori/go.uuid"

// Config contains the env variable structure for the project
type Config struct {
	baseURL string
	token string
	cartID string
	productIDs []string
}

// NewConfig returns an object containing the required env variables for the application
func NewConfig() Config {
	
	cartUUID := uuid.Must(uuid.NewV4()).String()

	config := Config{
		baseURL: "https://api.moltin.com/v2",
		token: "b0d1093887e6990cb583211bac3536c109cc3e50",
		cartID: cartUUID,
		productIDs: []string{"0542b5e4-cbb6-4960-bdb4-906c68512ed3", "cca740cf-0dab-4fb2-a71c-0391c48b37d0"},
	} 

	return config
}
