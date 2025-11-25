# PayPilot Go Client

Неофициальный API клиент для платежной системы [PayPilot](https://www.paypilot.ru/).
Simple wrapper written in Go.

**Author Website / Target Service:** [https://www.paypilot.ru/](https://www.paypilot.ru/)

## Installation

```bash
go get github.com/username/paypilot-go

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


Disclaimer
This library is not affiliated with PayPilot directly. Please refer to the official documentation at https://www.paypilot.ru/ for specific endpoint details and signing algorithms.

 Инструкция по заливке на GitHub

1.  Создайте новый репозиторий на GitHub (например, `paypilot-go`).
2.  Откройте терминал в папке с этими файлами.
3.  Выполните команды:

```bash
git init
git add .
git commit -m "Initial commit: PayPilot API client implementation"
git branch -M main
# Замените username на ваш логин GitHub
git remote add origin https://github.com/username/paypilot-go.git
git push -u origin main
