// utils/hmac.go
package utils

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"net/http"
	"time"
)

const HMACTimeout = 5 // seconds

func SignRequest(payload, secret string) string {
	mac := hmac.New(sha256.New, []byte(secret))
	mac.Write([]byte(payload))
	return base64.StdEncoding.EncodeToString(mac.Sum(nil))
}

func VerifySignature(expectedMAC, payload, secret string) bool {
	expected, _ := base64.StdEncoding.DecodeString(expectedMAC)
	mac := hmac.New(sha256.New, []byte(secret))
	mac.Write([]byte(payload))
	return hmac.Equal(mac.Sum(nil), expected)
}

func IsRequestFresh(timestampHeader string) bool {
	timestamp := ParseTimestamp(timestampHeader)
	if timestamp.IsZero() {
		return false
	}

	now := time.Now().UTC()
	diff := now.Sub(timestamp).Seconds()

	return diff >= -HMACTimeout && diff <= HMACTimeout
}

func ParseTimestamp(ts string) time.Time {
	t, _ := time.Parse(time.RFC3339, ts)
	return t
}

func RejectUnauthorized(w http.ResponseWriter, msg string) {
	fmt.Fprintf(w, "❌ Unauthorized: %s\n", msg)
	w.WriteHeader(http.StatusUnauthorized)
}
