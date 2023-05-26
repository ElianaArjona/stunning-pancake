package services

import (
	"fmt"
	"strconv"

	source "github.com/elianaarjona/stunning-pancake/source"
	utils "github.com/elianaarjona/stunning-pancake/utils"
)

func ParseExcelFile(rows [][]string) ([]*source.RawEntry, error) {

	var entries = []*source.RawEntry{}

	for i, row := range rows {
		//Skip unsed excel rows
		if i >= source.ExcelDataStartRow {

			// Each Row is a new Entry
			entry := &source.RawEntry{}

			entry.Name = "Banco General"

			// tm, err := BuildTimestampExcel(row[1])
			// if err != nil {
			// 	log.Fatal("Error Date Excel Date Parsing")
			// }
			entry.TransactionDate = row[1]
			// fmt.Println(entry.TransactionDate, "ts date")

			entry.Description = row[3]

			// fmt.Println(row[4])

			money, err := strconv.ParseFloat(row[4], 64)
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

			entry.Balance, err = strconv.ParseFloat(row[5], 64)
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

	e.Bank = r.Name

	e.Income = r.Income
	e.Expense = r.Expense
	e.Balance = r.Balance

	r.IncomeDescription()
	e.Description = r.Description

	if fileType == "excel" {
		// var mc = &source.BgTransactionType{}
		r.GetServicesType()
		e.Type = r.Type

		tm, err := utils.BuildTimestamp(r.TransactionDate, source.ExcelEpoch)
		if err != nil {
			fmt.Println("Error Time Parse Excel", err)
			return nil, err
		}

		e.Day = int64(tm.Day())
		e.Month = int64(tm.Month())
		e.Year = int64(tm.Year())

	}

	return e, nil

}
