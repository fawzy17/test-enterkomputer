package utils

import (
	"crypto/rand"
	"encoding/json"
	"fmt"
	"math/big"
	"net/http"
	"time"

	"github.com/fawzy17/test-enterkomputer/types"
	"github.com/go-playground/validator/v10"
)

var Validate = validator.New()

func ParseJSON(r *http.Request, payload any) error {
	if r.Body == nil {
		return fmt.Errorf("missing request body")
	}

	return json.NewDecoder(r.Body).Decode(payload)
}

func WriteJSON(w http.ResponseWriter, status int, v any) error {
	w.Header().Add("Content-Type", "application.json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(v)
}

func WriteError(w http.ResponseWriter, status int, err error) {
	WriteJSON(w, status, map[string]string{"error": err.Error()})
}

func GenerateRandomString(length int) (string, error) {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	result := make([]byte, length)
	for i := range result {
		num, err := rand.Int(rand.Reader, big.NewInt(int64(len(charset))))
		if err != nil {
			return "", err
		}
		result[i] = charset[num.Int64()]
	}
	return string(result), nil
}

func GenerateUniqueID() (string, error) {
	currentDate := time.Now().Format("020106")

	randomString, err := GenerateRandomString(5)
	if err != nil {
		return "", err
	}

	uniqueID := currentDate + randomString
	return uniqueID, nil
}

func RemoveDuplicate(data []types.ProductResponse) []types.ProductResponse {
	seen := make(map[int]types.ProductResponse, 0)

	for _, item := range data {
		key := item.ID
		if existing, exists := seen[key]; exists {
			existing.Quantity += item.Quantity
			seen[key] = existing
		} else {
			seen[key] = item
		}
	}

	uniqueData := []types.ProductResponse{}

	for _, v := range seen {
		uniqueData = append(uniqueData, v)
	}

	return uniqueData
}
