package services

import (
	"fmt"
	"log"
	"sort"

	"github.com/360EntSecGroup-Skylar/excelize"
	source "github.com/elianaarjona/stunning-pancake/source"
)

const (
	SheetCal string = "Calulations"
	SheetTop string = "Top Income"
)

type CountTotalResult struct {
	Year    int64   `json:"year"`
	Month   int64   `json:"month"`
	Type    string  `json:"type"`
	Income  float64 `json:"income"`
	Expense float64 `json:"expense"`
	Total   float64 `json:"total"`
	// Description string  `json:"description"`
}
type CountByIncomeType struct {
	Year        int64   `json:"year"`
	Month       int64   `json:"month"`
	Type        string  `json:"type"`
	Total       float64 `json:"total"`
	Description string  `json:"description"`
}
type ReportIncome struct {
	AccountName string              `json:"account_name,omitempty"`
	Report      []CountByIncomeType `json:"report,omitempty"`
	Total       []CountTotalResult  `json:"totsld,omitempty"`
}

func GetTotals(entries *source.BgEntries) []CountTotalResult {

	countResults := make(map[string]CountTotalResult)

	for _, e := range entries.Entries {
		key := fmt.Sprintf("%d-%d-%s", e.Year, e.Month, e.Type)

		result := countResults[key]

		result.Year = e.Year
		result.Month = e.Month

		result.Type = e.Type

		result.Income += e.Income
		result.Expense += e.Expense

		result.Total += e.Income + e.Expense

		countResults[key] = result

	}

	// Convert the count results to a slice of CountResult structs
	var results []CountTotalResult
	for _, result := range countResults {
		results = append(results, result)
	}

	return results
}

func GetIncomeTops(entries *source.BgEntries) []CountByIncomeType {

	countResults := make(map[string]CountByIncomeType)

	for _, e := range entries.Entries {
		key := fmt.Sprintf("%d-%d-%s", e.Year, e.Month, e.Type)

		if e.Expense >= 0 {
			result := countResults[key]

			result.Year = e.Year
			result.Month = e.Month

			result.Type = e.Type

			result.Description = e.Description

			result.Total += e.Income

			countResults[key] = result
		}

	}

	// Convert the count results to a slice of CountResult structs
	var results []CountByIncomeType
	for _, result := range countResults {
		if result.Type != "NA" && result.Total > 0 {
			results = append(results, result)
		}
	}

	// Sorting the slice by multiple fields
	sort.SliceStable(results, func(i, j int) bool {
		// Sort by Recent Year in descending order
		if results[i].Year != results[j].Year {
			return results[i].Year > results[j].Year
		}

		// Sort by Month in descending order
		if results[i].Month != results[j].Month {
			return results[i].Month > results[j].Month
		}

		// Sort by Total in descending order
		if results[i].Total != results[j].Total {
			return results[i].Total > results[j].Total
		}

		// Sort by Type in descending order
		if results[i].Type != results[j].Type {
			return results[i].Type > results[j].Type
		}

		// Sort by Description in descending order
		return results[i].Description > results[j].Description
	})

	return results
}
func createFileCalculationIncome(f *excelize.File, inc *ReportIncome) (*excelize.File, error) {

	index := f.NewSheet("Sheet1")
	f.SetActiveSheet(index)
	f.SetSheetName("Sheet1", SheetCal)

	header := []string{"Year", "Month", "Type", "Income", "Expense", "Total"}
	for i, value := range header {
		col := string('A' + i)
		cell := col + "1"
		f.SetCellValue(SheetCal, cell, value)
	}

	row := 2
	for _, result := range inc.Total {
		f.SetCellValue(SheetCal, fmt.Sprintf("A%d", row), result.Year)
		f.SetCellValue(SheetCal, fmt.Sprintf("B%d", row), result.Month)
		f.SetCellValue(SheetCal, fmt.Sprintf("C%d", row), result.Type)
		f.SetCellValue(SheetCal, fmt.Sprintf("D%d", row), result.Income)
		f.SetCellValue(SheetCal, fmt.Sprintf("E%d", row), result.Expense)
		f.SetCellValue(SheetCal, fmt.Sprintf("F%d", row), result.Total)
		row++
	}

	return f, nil
}

func (inc *ReportIncome) ExportToExcel(filename string) error {

	f := excelize.NewFile()

	createFileCalculationIncome(f, inc)

	err := f.SaveAs(filename)
	if err != nil {
		log.Fatal(err)
		return err
	}

	return nil
}
