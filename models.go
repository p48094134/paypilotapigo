package paypilot

// PaymentRequest - структура для создания платежа
type PaymentRequest struct {
	OrderID     string  `json:"order_id"`
	Amount      float64 `json:"amount"`
	Currency    string  `json:"currency"` // RUB, USD, etc.
	Description string  `json:"description"`
	CallbackURL string  `json:"callback_url"`
	CustomerIP  string  `json:"customer_ip,omitempty"`
}

// PaymentResponse - ответ от API
type PaymentResponse struct {
	Success    bool   `json:"success"`
	PaymentURL string `json:"payment_url"` // Ссылка на оплату
	PaymentID  string `json:"payment_id"`
	Error      string `json:"error,omitempty"`
}

// CheckStatusResponse - ответ проверки статуса
type CheckStatusResponse struct {
	Success bool   `json:"success"`
	Status  string `json:"status"` // paid, pending, failed
}
