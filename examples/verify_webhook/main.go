package main

import (
	"fmt"

	cryptoix "github.com/cryptoix/cryptoix-go"
)

func main() {
	payload := []byte(`{"uuid":"tx_test_123","status":"completed","amount_fiat":100,"timestamp":1780574400}`)
	signature := "9bc6d23d70a52f2960d404a6bf2fd067ee8f15e2e30c4e53cfde3559251a7e25"
	fmt.Println(cryptoix.VerifyWebhookSignature(payload, signature, "whsec_test_secret_1234567890"))
}
