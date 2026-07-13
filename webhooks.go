package cryptoix

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
)

func ComputeWebhookSignature(payload []byte, secret string) string {
	mac := hmac.New(sha256.New, []byte(secret))
	mac.Write(payload)
	return hex.EncodeToString(mac.Sum(nil))
}

func VerifyWebhookSignature(payload []byte, signature string, secret string) bool {
	expected := ComputeWebhookSignature(payload, secret)
	return hmac.Equal([]byte(expected), []byte(signature))
}
