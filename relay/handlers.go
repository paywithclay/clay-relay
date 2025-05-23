// relay/handlers.go
package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/paywithclay/clay-relay/utils"
)

func ChargeHandler(w http.ResponseWriter, r *http.Request) {
	secret := os.Getenv("CLAY_RELAY_SECRET")
	if secret == "" {
		utils.RejectUnauthorized(w, "missing shared secret")
		return
	}

	timestamp := r.Header.Get("X-Clerk-Timestamp")
	signature := r.Header.Get("X-Clerk-Signature")

	if !utils.IsRequestFresh(timestamp) {
		utils.RejectUnauthorized(w, "request expired")
		return
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "failed to read body", http.StatusInternalServerError)
		return
	}

	if !utils.VerifySignature(signature, string(body), secret) {
		utils.RejectUnauthorized(w, "invalid signature")
		return
	}

	// Proceed with charge logic
	response := map[string]string{"url": "https://paystack.com/pay/demo "}
	json.NewEncoder(w).Encode(response)
}

func VerifyHandler(w http.ResponseWriter, r *http.Request) {
	// Same pattern as above
	fmt.Fprintf(w, "🔍 Verified successfully")
}

func HealthCheck(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"status": "healthy",
		"name":   "Clay Relay",
	})
}
