package cryptoix

import "testing"

func TestWebhookSignatureVector(t *testing.T) {
	payload := []byte(`{"uuid":"tx_test_123","status":"completed","amount_fiat":100,"timestamp":1780574400}`)
	secret := "whsec_test_secret_1234567890"
	sig := "9bc6d23d70a52f2960d404a6bf2fd067ee8f15e2e30c4e53cfde3559251a7e25"
	if ComputeWebhookSignature(payload, secret) != sig { t.Fatal("signature mismatch") }
	if !VerifyWebhookSignature(payload, sig, secret) { t.Fatal("signature should verify") }
	if VerifyWebhookSignature(payload, "bad", secret) { t.Fatal("bad signature verified") }
}
