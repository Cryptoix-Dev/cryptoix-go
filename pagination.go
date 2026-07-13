package cryptoix

import "net/url"

type Pagination struct {
	Page    int
	PerPage int
}

func (p Pagination) Values(extra map[string]string) url.Values {
	v := url.Values{}
	if p.Page > 0 { v.Set("page", intString(p.Page)) }
	if p.PerPage > 0 { v.Set("per_page", intString(p.PerPage)) }
	for k, val := range extra { if val != "" { v.Set(k, val) } }
	return v
}

func intString(n int) string { return fmtInt(n) }
