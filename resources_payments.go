package cryptoix

import (
	"context"
	"net/url"
)

func (c *Client) ListCurrencies(ctx context.Context) (any, error) { return c.do(ctx, "GET", "/currencies", nil, nil, false) }
func (c *Client) ListRates(ctx context.Context) (any, error) { return c.do(ctx, "GET", "/rates", nil, nil, true) }
func (c *Client) CreatePayment(ctx context.Context, payload map[string]any) (any, error) { return c.do(ctx, "POST", "/payment/create", payload, nil, true) }
func (c *Client) CheckPayment(ctx context.Context, uuid string) (any, error) { return c.do(ctx, "GET", "/payment/"+url.PathEscape(uuid)+"/check", nil, nil, true) }
