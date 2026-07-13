package cryptoix

import (
	"context"
	"net/url"
)

func (c *Client) ListTransactions(ctx context.Context, q url.Values) (any, error) { return c.do(ctx, "GET", "/transactions", nil, q, true) }
func (c *Client) GetTransaction(ctx context.Context, uuid string) (any, error) { return c.do(ctx, "GET", "/transactions/"+url.PathEscape(uuid), nil, nil, true) }
func (c *Client) ListBalances(ctx context.Context) (any, error) { return c.do(ctx, "GET", "/balances", nil, nil, true) }
func (c *Client) ListWithdrawals(ctx context.Context, q url.Values) (any, error) { return c.do(ctx, "GET", "/withdrawals", nil, q, true) }
func (c *Client) CreateWithdrawal(ctx context.Context, payload map[string]any) (any, error) { return c.do(ctx, "POST", "/withdrawals", payload, nil, true) }
func (c *Client) GetWithdrawal(ctx context.Context, uuid string) (any, error) { return c.do(ctx, "GET", "/withdrawals/"+url.PathEscape(uuid), nil, nil, true) }
func (c *Client) ListRefunds(ctx context.Context, q url.Values) (any, error) { return c.do(ctx, "GET", "/refunds", nil, q, true) }
func (c *Client) CreateRefund(ctx context.Context, payload map[string]any) (any, error) { return c.do(ctx, "POST", "/refunds", payload, nil, true) }
func (c *Client) GetRefund(ctx context.Context, uuid string) (any, error) { return c.do(ctx, "GET", "/refunds/"+url.PathEscape(uuid), nil, nil, true) }
func (c *Client) ListPayouts(ctx context.Context, q url.Values) (any, error) { return c.do(ctx, "GET", "/payouts", nil, q, true) }
func (c *Client) CreatePayout(ctx context.Context, payload map[string]any) (any, error) { return c.do(ctx, "POST", "/payouts", payload, nil, true) }
func (c *Client) GetPayout(ctx context.Context, uuid string) (any, error) { return c.do(ctx, "GET", "/payouts/"+url.PathEscape(uuid), nil, nil, true) }
func (c *Client) SubmitPayout(ctx context.Context, uuid string, payload map[string]any) (any, error) { return c.do(ctx, "POST", "/payouts/"+url.PathEscape(uuid)+"/submit", payload, nil, true) }
func (c *Client) CancelPayout(ctx context.Context, uuid string, payload map[string]any) (any, error) { return c.do(ctx, "POST", "/payouts/"+url.PathEscape(uuid)+"/cancel", payload, nil, true) }
func (c *Client) ListPaymentLinks(ctx context.Context, q url.Values) (any, error) { return c.do(ctx, "GET", "/payment-links", nil, q, true) }
func (c *Client) CreatePaymentLink(ctx context.Context, payload map[string]any) (any, error) { return c.do(ctx, "POST", "/payment-links", payload, nil, true) }
func (c *Client) ListInvoices(ctx context.Context, q url.Values) (any, error) { return c.do(ctx, "GET", "/invoices", nil, q, true) }
func (c *Client) CreateInvoice(ctx context.Context, payload map[string]any) (any, error) { return c.do(ctx, "POST", "/invoices", payload, nil, true) }
func (c *Client) GetInvoice(ctx context.Context, uuid string) (any, error) { return c.do(ctx, "GET", "/invoices/"+url.PathEscape(uuid), nil, nil, true) }
func (c *Client) SendInvoice(ctx context.Context, uuid string) (any, error) { return c.do(ctx, "POST", "/invoices/"+url.PathEscape(uuid)+"/send", nil, nil, true) }
func (c *Client) ListEscrows(ctx context.Context, q url.Values) (any, error) { return c.do(ctx, "GET", "/escrows", nil, q, true) }
func (c *Client) CreateEscrow(ctx context.Context, payload map[string]any) (any, error) { return c.do(ctx, "POST", "/escrows", payload, nil, true) }
func (c *Client) GetEscrow(ctx context.Context, uuid string) (any, error) { return c.do(ctx, "GET", "/escrows/"+url.PathEscape(uuid), nil, nil, true) }
func (c *Client) FundEscrow(ctx context.Context, uuid string, payload map[string]any) (any, error) { return c.do(ctx, "POST", "/escrows/"+url.PathEscape(uuid)+"/fund", payload, nil, true) }
func (c *Client) RequestEscrowRelease(ctx context.Context, uuid string, payload map[string]any) (any, error) { return c.do(ctx, "POST", "/escrows/"+url.PathEscape(uuid)+"/request-release", payload, nil, true) }
func (c *Client) DisputeEscrow(ctx context.Context, uuid string, payload map[string]any) (any, error) { return c.do(ctx, "POST", "/escrows/"+url.PathEscape(uuid)+"/dispute", payload, nil, true) }
func (c *Client) ListWebhookDeliveries(ctx context.Context, q url.Values) (any, error) { return c.do(ctx, "GET", "/webhook-deliveries", nil, q, true) }
func (c *Client) ReplayWebhookDelivery(ctx context.Context, id int) (any, error) { return c.do(ctx, "POST", "/webhook-deliveries/"+fmtInt(id)+"/replay", nil, nil, true) }
