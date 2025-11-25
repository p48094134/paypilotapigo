package main

import (
	"fmt"
	"log"
	"github.com/username/paypilot-go"
)

func main() {
	// Initialize client
	client := paypilot.NewClient("YOUR_API_KEY", "YOUR_SECRET_KEY")

	// Create payment
	req := paypilot.PaymentRequest{
		OrderID:     "ORDER-12345",
		Amount:      1500.00,
		Currency:    "RUB",
		Description: "Payment for services",
		CallbackURL: "https://mysite.com/callback",
	}

	resp, err := client.CreatePayment(req)
	if err != nil {
		log.Fatalf("Error creating payment: %v", err)
	}

	fmt.Printf("Payment created! URL: %s\n", resp.PaymentURL)
}
