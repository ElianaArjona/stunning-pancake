package services

import "github.com/golang/protobuf/ptypes/timestamp"

type BacExcel struct {
	TransactionDate timestamp.Timestamp `json:"transaction_date,omitempty"`
	Description     string              `json:"description,omitempty"`
	Expenses        float32             `json:"expenses,omitempty"`
	Income          float32             `json:"income,omitempty"`
	Balance         float32             `json:"balance,omitempty"`
}
