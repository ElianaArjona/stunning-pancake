package services

import (
	"strings"
	"time"
)

const (
	TimeLayout        string = "15/01/2006"
	TimeLayoutMonth   string = "02-Jan-06"
	SheetName         string = "BGPExcelReport"
	ExcelDataStartRow int    = 8
)

var ExcelEpoch = time.Date(1899, 12, 30, 0, 0, 0, 0, time.UTC)

const (
	Deposit = iota
	ACH
	Tranference
	Yappy
	NA
)

var serviceTypes = []string{
	"Deposit",
	"ACH",
	"Tranference",
	"Yappy",
	"NA",
}

type RawEntry struct {
	Name            string  `json:"name,omitempty"`
	TransactionDate string  `json:"transaction_date,omitempty"`
	Reference       int64   `json:"reference,omitempty"`
	Description     string  `json:"description,omitempty"`
	Type            string  `json:"type,omitempty"`
	Expense         float64 `json:"expenses"`
	Income          float64 `json:"income"`
	Balance         float64 `json:"balance"`
}

// type BgTransactionType struct {
// 	// Interest string `json:"interest,omitempty"`
// 	ACH      bool `json:"ach,omitempty"`
// 	Yappy    bool `json:"yappy,omitempty"`
// 	Deposito bool `json:"deposito,omitempty"`
// 	// Service  []string `json:"service,omitempty"`
// }

type Entry struct {
	Bank        string  `json:"bank,omitempty"`
	Day         int64   `json:"day,omitempty"`
	Month       int64   `json:"month,omitempty"`
	Year        int64   `json:"year,omitempty"`
	Income      float64 `json:"income,omitempty"`
	Expense     float64 `json:"expense,omitempty"`
	Balance     float64 `json:"balance,omitempty"`
	Type        string  `json:"type,omitempty"`
	Description string  `json:"description,omitempty"`
}

type BgEntries struct {
	Entries []*Entry `json:"entries,omitempty"`
}

func (r *RawEntry) GetServicesType() {
	r.Description = strings.ToLower(r.Description)

	switch {
	case strings.Contains(r.Description, strings.ToLower(serviceTypes[Deposit])):
		r.Type = serviceTypes[Deposit]
	case strings.Contains(r.Description, strings.ToLower(serviceTypes[ACH])):
		r.Type = serviceTypes[ACH]
	case strings.Contains(r.Description, strings.ToLower(serviceTypes[Tranference])):
		r.Type = serviceTypes[ACH]
	case strings.Contains(r.Description, strings.ToLower(serviceTypes[Yappy])):
		r.Type = serviceTypes[Yappy]
	default:
		r.Type = serviceTypes[NA]
	}
}

func (r *RawEntry) IncomeDescription() {
	input := r.Description
	// Check if the input contains "DE" followed by the desired text
	if strings.Contains(r.Description, "TRANSFERENCIA DE ") {
		index := strings.Index(r.Description, " DE ")
		text := input[index+4:]
		r.Description = text
	}

	// Check if the r.Description starts with "ACH - "
	if strings.HasPrefix(r.Description, "ACH - ") {
		text := input[6:]
		r.Description = text
	}

	// Check if the r.Description starts with "YAPPY DE " and contains " POR "
	if strings.HasPrefix(r.Description, "YAPPY DE ") && strings.Contains(r.Description, " POR ") {
		index := strings.Index(r.Description, " POR ")
		text := input[9:index]
		r.Description = text
	}

}
