package services

import (
	"strings"
	"time"
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
	"deposito",
	"ach",
	"transferencia",
	"yappy",
	"NA",
}

type RawEntry struct {
	BankName        string  `json:"bank_name,omitempty"`
	TransactionDate string  `json:"transaction_date,omitempty"`
	Reference       int64   `json:"reference,omitempty"`
	Description     string  `json:"description,omitempty"`
	Type            string  `json:"type,omitempty"`
	Expense         float64 `json:"expenses,omitempty"`
	Income          float64 `json:"income,omitempty"`
	Balance         float64 `json:"balance,omitempty"`
}

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

	if strings.Contains(r.Description, serviceTypes[Deposit]) && (!strings.Contains(r.Description, serviceTypes[Yappy]) ||
		!strings.Contains(r.Description, serviceTypes[ACH]) ||
		!strings.Contains(r.Description, serviceTypes[Tranference])) {
		r.Type = serviceTypes[Deposit]

	} else if strings.Contains(r.Description, serviceTypes[Yappy]) && (!strings.Contains(r.Description, serviceTypes[Deposit]) ||
		!strings.Contains(r.Description, serviceTypes[ACH]) ||
		!strings.Contains(r.Description, serviceTypes[Tranference])) {
		r.Type = serviceTypes[Yappy]

	} else if strings.Contains(r.Description, serviceTypes[ACH]) && (!strings.Contains(r.Description, serviceTypes[Deposit]) ||
		!strings.Contains(r.Description, serviceTypes[Yappy]) ||
		!strings.Contains(r.Description, serviceTypes[Tranference])) {
		r.Type = serviceTypes[ACH]

	} else if strings.Contains(r.Description, serviceTypes[Tranference]) && (!strings.Contains(r.Description, serviceTypes[Deposit]) ||
		!strings.Contains(r.Description, serviceTypes[Yappy]) ||
		!strings.Contains(r.Description, serviceTypes[ACH])) {
		r.Type = serviceTypes[ACH]
	} else {
		r.Type = serviceTypes[NA]
	}

}

func (r *RawEntry) IncomeDescription() {
	input := r.Description

	// Check if the input contains "DE" followed by the desired text
	if strings.Contains(r.Description, "transferencia de ") {
		index := strings.Index(r.Description, " de ")
		text := input[index+4:]
		r.Description = text

	} else if strings.HasPrefix(r.Description, "ach - ") {
		// Check if the r.Description starts with "ACH - "
		text := input[6:]
		r.Description = text

	} else if strings.Contains(r.Description, "yappy de ") && strings.Contains(r.Description, " por ") {
		// Check if the r.Description starts with "YAPPY DE " and contains " POR "
		index := strings.Index(r.Description, " por ")
		text := input[9:index]
		r.Description = text

	} else if strings.Contains(r.Description, "yappy de ") {
		// Check if the r.Description starts with "YAPPY DE "
		text := input[9:]
		r.Description = text

	} else {
		r.Description = input
	}

}
