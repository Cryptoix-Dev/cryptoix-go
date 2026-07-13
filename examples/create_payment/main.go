package main

import (
	"context"
	"fmt"
	"os"

	cryptoix "github.com/cryptoix/cryptoix-go"
)

func main() {
	key := os.Getenv("CRYPTOIX_API_KEY")
	if key == "" { key = "ak_live_xxxxxxxxxxxx" }
	client := cryptoix.NewClient(key)
	payment, err := client.CreatePayment(context.Background(), map[string]any{"amount": 49.99, "order_id": "ORDER-1001"})
	if err != nil { panic(err) }
	fmt.Printf("%v\n", payment)
}
