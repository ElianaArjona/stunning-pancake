package services

import (
	"fmt"
	"strconv"

	source "github.com/elianaarjona/stunning-pancake/source"
	utils "github.com/elianaarjona/stunning-pancake/utils"
)

type Column int

const (
	A Column = iota
	B
	C
	D
	E
	F
	G
	H
	I
	J
	K
)

func ParseExceMovimientos(rows [][]string, ExcelDataStartRow int) ([]*source.RawEntry, error) {

	var entries = []*source.RawEntry{}

	for i, row := range rows {

		//Skip unsed excel rows
		if i >= ExcelDataStartRow {

			// Each Row is a new Entry
			entry := &source.RawEntry{}
			entry.BankName = "Banco General"

			entry.TransactionDate = row[B]
			entry.Description = row[D]
			money, err := strconv.ParseFloat(row[E], 64)
			if err != nil {
				return nil, err
			}
			if money < 0 {
				entry.Expense = money
				entry.Income = 0.0
			} else {
				entry.Income = money
				entry.Expense = 0.0
			}
			entry.Balance, err = strconv.ParseFloat(row[F], 64)
			if err != nil {
				return nil, err
			}

			entries = append(entries, entry)
		}

	}

	// jsonData, err := json.MarshalIndent(bg, "", "  ")
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// fmt.Println(string(jsonData))

	return entries, nil
}

func ParseExcelEstadosCuenta(rows [][]string, ExcelDataStartRow int) ([]*source.RawEntry, error) {

	var entries = []*source.RawEntry{}

	for i, row := range rows {

		//Skip unsed excel rows
		if i >= ExcelDataStartRow {
			expense := 0.0
			income := 0.0

			// Each Row is a new Entry
			entry := &source.RawEntry{}
			entry.BankName = "Banco General"

			if row[A] != "" {

				entry.TransactionDate = row[A]
			} else {
				entry.TransactionDate = row[B]
			}

			entry.Description = row[E]

			expense, err := strconv.ParseFloat(row[F], 64)
			if err != nil {
				income, err = strconv.ParseFloat(row[G], 64)
				if err != nil {
					income, err = strconv.ParseFloat(row[H], 64)
					if err != nil {
						return nil, err
					}
				}
				entry.Income = income
				entry.Expense = 0.0
			}

			income, err = strconv.ParseFloat(row[G], 64)
			if err != nil {
				income, err = strconv.ParseFloat(row[H], 64)
				if err != nil {
					expense, err = strconv.ParseFloat(row[F], 64)
					if err != nil {
						return nil, err
					}
				}
				entry.Expense = expense
				entry.Income = 0.0
			}

			entry.Balance, err = strconv.ParseFloat(row[I], 64)
			if err != nil {
				return nil, err
			}

			entries = append(entries, entry)
		}

	}

	// jsonData, err := json.MarshalIndent(bg, "", "  ")
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// fmt.Println(string(jsonData))

	return entries, nil
}

func ProcessRawToEntry(r *source.RawEntry, fileType string) (*source.Entry, error) {

	var e = &source.Entry{}

	e.Bank = r.BankName

	e.Income = r.Income
	e.Expense = r.Expense
	e.Balance = r.Balance

	r.GetServicesType()
	e.Type = r.Type

	r.IncomeDescription()
	e.Description = r.Description

	tm, err := utils.BuildTimestamp(r.TransactionDate, source.ExcelEpoch)
	if err != nil {
		fmt.Println("Error Time Parse Excel", err)
		return nil, err
	}

	e.Day = int64(tm.Day())
	e.Month = int64(tm.Month())
	e.Year = int64(tm.Year())

	return e, nil

}
