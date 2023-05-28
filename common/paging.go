package common

import "strings"

// Paging is used for store paging options
//
// # The default Order of paging is ASC
//
// The default Limit is 10 and must be between 5 and 20
type Paging struct {
	LastId *string `form:"last_id"`
	Order  *Order  `form:"order"`
	Limit  *int    `form:"limit"`
}

// Order is a enum type, it is used for determine the order of the query
//
// # Order includes 2 values ASC and DESC, respectively mean Ascending and Descending order
//
// Default values of Order if not provided is ASC
type Order string

var (
	ASC  Order = "asc"
	DESC       = "desc"
)

var defaultLimit = 10

func (p *Paging) Process() {
	if p.Order == nil ||
		strings.TrimSpace(strings.ToLower(string(*p.Order))) == strings.TrimSpace(strings.ToLower(DESC)) {
		p.Order = &ASC
	}
	if p.Limit == nil || *p.Limit < 5 || *p.Limit > 20 {
		p.Limit = &defaultLimit
	}
}
