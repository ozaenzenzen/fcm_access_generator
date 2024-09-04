package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"golang.org/x/oauth2/google"
	"google.golang.org/api/option"
	"google.golang.org/api/transport"
)

func main() {
	// Path to your service account key file
	serviceAccountPath := "/Users/qoin/Downloads/service-account-file.json"

	// Read the service account JSON file
	creds, err := os.ReadFile(serviceAccountPath)
	if err != nil {
		log.Fatalf("Failed to read service account file: %v", err)
	}

	// Initialize a new Google client with the credentials
	ctx := context.Background()
	config, err := google.JWTConfigFromJSON(creds, "https://www.googleapis.com/auth/firebase.messaging")
	if err != nil {
		log.Fatalf("Failed to parse service account JSON: %v", err)
	}

	// Create a new HTTP client using the JWT configuration
	client, _, err := transport.NewHTTPClient(ctx, option.WithTokenSource(config.TokenSource(ctx)))
	if err != nil {
		log.Fatalf("Failed to create HTTP client: %v", err)
	}
	fmt.Printf("client: %s\n", client)

	// Use the HTTP client to get an access token
	tokenSource := config.TokenSource(ctx)
	token, err := tokenSource.Token()
	if err != nil {
		log.Fatalf("Failed to get token: %v", err)
	}

	fmt.Printf("Access Token: %s\n", token.AccessToken)
}
