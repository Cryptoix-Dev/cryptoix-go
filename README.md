# Cryptoix Go SDK

Official Go SDK for the Cryptoix.io Merchant API.

## Install

```bash
go get github.com/cryptoix/cryptoix-go
```

## Quick start

```go
client := cryptoix.NewClient("ak_live_xxxxxxxxxxxx")
payment, err := client.CreatePayment(ctx, map[string]any{"amount": 49.99})
```

Never send API keys in query strings. Use header authentication only.
