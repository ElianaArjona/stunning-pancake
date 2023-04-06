package services

import "time"

const (
	timeLayout        = "15/01/2006"
	sheetName         = "BGPExcelReport"
	excelDataStartRow = 8
)

var excelEpoch = time.Date(1899, 12, 30, 0, 0, 0, 0, time.UTC)

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

type BgServicesTypesTxt struct {
	Deposito int64   `json:"interest,omitempty"`
	ACH      int64   `json:"ach,omitempty"`
	Yappy    []int64 `json:"yappy,omitempty"`
	Service  []int64 `json:"service,omitempty"`
}

type BgServicesTypesExcel struct {
	// Interest string `json:"interest,omitempty"`
	ACH      string   `json:"ach,omitempty"`
	Yappy    string   `json:"yappy,omitempty"`
	Deposito string   `json:"deposito,omitempty"`
	Service  []string `json:"service,omitempty"`
}

type Entry struct {
	Bank    string  `json:"bank,omitempty"`
	Month   int64   `json:"month,omitempty"`
	Year    int64   `json:"year,omitempty"`
	Income  float64 `json:"income,omitempty"`
	Expense float64 `json:"expense,omitempty"`
	Balance float64 `json:"balance,omitempty"`
	Type    string  `json:"type,omitempty"`
}

type BgEntries struct {
	Entries []*Entry `json:"entries,omitempty"`
}
