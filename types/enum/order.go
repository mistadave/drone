// Copyright 2021 Harness Inc. All rights reserved.
// Use of this source code is governed by the Polyform Free Trial License
// that can be found in the LICENSE.md file for this repository.

package enum

import "strings"

// Order defines the sort order.
type Order int

// Order enumeration.
const (
	OrderDefault Order = iota
	OrderAsc
	OrderDesc
)

// String returns the Order as a string.
func (e Order) String() string {
	if e == OrderDesc {
		return "desc"
	}
	return "asc"
}

// ParseOrder parses the order string and returns
// an order enumeration.
func ParseOrder(s string) Order {
	switch strings.ToLower(s) {
	case "asc", "ascending":
		return OrderAsc
	case "desc", "descending":
		return OrderDesc
	default:
		return OrderDefault
	}
}