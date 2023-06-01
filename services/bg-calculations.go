package services

import (
	"encoding/csv"
	"fmt"
	"os"
	"sort"

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

	// Print the count results
	// for _, result := range results {
	// 	fmt.Printf("Year: %d, Month: %d, Type: %s, Income: %.2f, Expense: %.2f, Total: %.2f\n",
	// 		result.Year, result.Month, result.Type, result.Income, result.Expense, result.Total)
	// }
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

	// Print the count results
	// for _, result := range results {
	// 	fmt.Printf("Year: %d, Month: %d, Type: %s, Description: %s, Total: %.2f\n",
	// 		result.Year, result.Month, result.Type, result.Description, result.Total)
	// }

	return results

	// report.ExportToCSV("./outputs/sample.csv")
}

func createFileTotal(writer *csv.Writer, inc *ReportIncome) error {
	// Write the CSV header
	// header := []string{"Year", "Month", "Type", "Income", "Expense", "Total"}
	header := []string{"Type", "Income", "Expense", "Total"}
	err := writer.Write(header)
	if err != nil {
		return err
	}

	// Write the count results to the CSV file
	for _, result := range inc.Total {
		row := []string{
			// fmt.Sprintf("%d", result.Year),
			// fmt.Sprintf("%d", result.Month),
			result.Type,
			fmt.Sprintf("%.2f", result.Income),
			fmt.Sprintf("%.2f", result.Expense),
			// result.Description,
			fmt.Sprintf("%.2f", result.Total),
		}
		err = writer.Write(row)
		if err != nil {
			return err
		}
	}
	return nil
}

func createFileTopResult(writer *csv.Writer, inc *ReportIncome) error {
	// Write the CSV header
	header := []string{"Year", "Month", "Type", "Description", "Total"}
	err := writer.Write(header)
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

func (inc *ReportIncome) ExportToCSV(filename string) error {
	// Create the CSV file
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	// Create a CSV writer
	writer := csv.NewWriter(file)
	defer writer.Flush()

	// createFileTotal(writer, inc)
	// if err != nil {
	// 	return err
	// }

	createFileTopResult(writer, inc)
	if err != nil {
		return err
	}

	return nil
}
