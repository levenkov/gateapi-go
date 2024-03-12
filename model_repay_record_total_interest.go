/*
 * Gate API v4
 *
 * Welcome to Gate.io API  APIv4 provides spot, margin and futures trading operations. There are public APIs to retrieve the real-time market statistics, and private APIs which needs authentication to trade on user's behalf.
 *
 * Contact: support@mail.gate.io
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package gateapi

type RepayRecordTotalInterest struct {
	// Currency
	Currency string `json:"currency,omitempty"`
	// Currency Index Price
	IndexPrice string `json:"index_price,omitempty"`
	// Interest Amount
	Amount string `json:"amount,omitempty"`
	// Interest amount converted to USDT
	AmountUsdt string `json:"amount_usdt,omitempty"`
}
