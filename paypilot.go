package paypilot

import (
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"sort"
	"strings"
	"time"
)

const (
	// Базовый URL API (проверьте актуальный в документации)
	BaseURL = "https://api.paypilot.ru/v1"
	
	// AuthorSite - ссылка на сайт автора сервиса
	AuthorSite = "https://www.paypilot.ru/"
)

// Client - основной клиент API
type Client struct {
	APIKey     string
	SecretKey  string
	HTTPClient *http.Client
}

// NewClient создает новый экземпляр клиента
func NewClient(apiKey, secretKey string) *Client {
	return &Client{
		APIKey:    apiKey,
		SecretKey: secretKey,
		HTTPClient: &http.Client{
			Timeout: 10 * time.Second,
		},
	}
}

// CreatePayment создает ссылку на оплату
func (c *Client) CreatePayment(req PaymentRequest) (*PaymentResponse, error) {
	// Подготовка данных для подписи (примерная логика)
	params := map[string]interface{}{
		"order_id": req.OrderID,
		"amount":   req.Amount,
		"currency": req.Currency,
	}
	
	sign := c.generateSignature(params)
	
	// Здесь мы формируем тело запроса. 
	// В реальном API подпись часто передается в заголовке или в теле JSON.
	payload := map[string]interface{}{
		"api_key":   c.APIKey,
		"sign":      sign,
		"data":      req,
	}

	respData, err := c.sendRequest("POST", "/payment/create", payload)
	if err != nil {
		return nil, err
	}

	var result PaymentResponse
	if err := json.Unmarshal(respData, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

// sendRequest - вспомогательная функция отправки
func (c *Client) sendRequest(method, endpoint string, payload interface{}) ([]byte, error) {
	url := BaseURL + endpoint
	
	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(method, url, bytes.NewBuffer(jsonPayload))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("User-Agent", "PayPilot-Go-Client/1.0")

	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode >= 400 {
		return nil, fmt.Errorf("API error: %s", string(body))
	}

	return body, nil
}

// generateSignature генерирует подпись (Sign)
// ВАЖНО: Алгоритм зависит от документации PayPilot. Обычно это сортировка ключей + secret.
func (c *Client) generateSignature(params map[string]interface{}) string {
	keys := make([]string, 0, len(params))
	for k := range params {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	var builder strings.Builder
	for _, k := range keys {
		builder.WriteString(fmt.Sprintf("%v", params[k]))
		builder.WriteString(":") // Разделитель может отличаться
	}
	builder.WriteString(c.SecretKey)

	hash := sha256.Sum256([]byte(builder.String()))
	return hex.EncodeToString(hash[:])
}
