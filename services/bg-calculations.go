package services

import (
	"encoding/csv"
	"fmt"
	"os"

	source "github.com/elianaarjona/stunning-pancake/source"
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
	Report []CountByIncomeType
}

func GetTotals(entries *source.BgEntries) {

	countResults := make(map[string]CountTotalResult)

	for _, e := range entries.Entries {
		key := fmt.Sprintf("%d-%d-%s", e.Year, e.Month, e.Type)

		result := countResults[key]

		result.Year = e.Year
		result.Month = e.Month

		result.Type = e.Type
		// result.Description = e.Description

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

	// Print the count results
	for _, result := range results {
		fmt.Printf("Year: %d, Month: %d, Type: %s, Income: %.2f, Expense: %.2f, Total: %.2f\n",
			result.Year, result.Month, result.Type, result.Income, result.Expense, result.Total)
	}
}

func GetIncomeTops(entries *source.BgEntries) {

	countResults := make(map[string]CountByIncomeType)

	for _, e := range entries.Entries {
		key := fmt.Sprintf("%d-%d-%s", e.Year, e.Month, e.Type)

		result := countResults[key]

		result.Year = e.Year
		result.Month = e.Month

		result.Type = e.Type
		result.Description = e.Description

		result.Total += e.Income

		countResults[key] = result
	}

	// Convert the count results to a slice of CountResult structs
	var results []CountByIncomeType
	for _, result := range countResults {
		if result.Type != "NA" && result.Total > 0 {
			results = append(results, result)
		}

	}

	// Print the count results
	for _, result := range results {
		fmt.Printf("Year: %d, Month: %d, Type: %s, Description: %s, Total: %.2f\n",
			result.Year, result.Month, result.Type, result.Description, result.Total)
	}

	report := &ReportIncome{
		Report: results,
	}
	report.exportToCSV("./outputs/sample.csv")
}

func (inc *ReportIncome) exportToCSV(filename string) error {
	// Create the CSV file
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	// Create a CSV writer
	writer := csv.NewWriter(file)
	defer writer.Flush()

	// Write the CSV header
	header := []string{"Year", "Month", "Type", "Description", "Total"}
	err = writer.Write(header)
	if err != nil {
		return err
	}

	// Write the count results to the CSV file
	for _, result := range inc.Report {
		row := []string{
			fmt.Sprintf("%d", result.Year),
			fmt.Sprintf("%d", result.Month),
			result.Type,
			result.Description,
			fmt.Sprintf("%.2f", result.Total),
		}
		err = writer.Write(row)
		if err != nil {
			return err
		}
	}

	return nil
}
