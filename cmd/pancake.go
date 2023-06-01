package main

import (
	"log"

	"github.com/elianaarjona/stunning-pancake/caller"
	"github.com/elianaarjona/stunning-pancake/services"
	utils "github.com/elianaarjona/stunning-pancake/utils"
)

const (
	TimeLayout        string = "15/01/2006"
	TimeLayoutMonth   string = "02-Jan-06"
	SheetName         string = "BGPExcelReport"
	ExcelDataStartRow int    = 8
	AccountNameLine   int    = 3
)

func main() {

	// verufy why icome is not calculate correctetly and verify the other calculation

	// bg := &services.BgEntries{}
	// bg.ParseTxtFile("./source/bg-sample.txt")
	// services.ParseExcelFile("source/bg-excel-sample.xlsx")

	var bg = &caller.BankConfig{
		BankName: "Banco General",
		FilePath: "samples/sample-new.xlsx",
		FileType: "excel",
		SheetName: "BGPExcelReport"
	}

	rows, err := utils.OpenExcelFile(bg.FilePath, bg.SheetName)
	if err != nil {
		log.Fatal()
	}

	reports, err := bg.ReportBG(rows, ExcelDataStartRow)
	if err != nil {
		log.Fatal()
	}

	report := services.ReportIncome{}

	report.AccountName = utils.GetAccountName(rows[AccountNameLine])
	report.Total = services.GetTotals(reports)
	report.Report = services.GetIncomeTops(reports)

	// report.ExportToExcel("./outputs/sample.xlsx")
	report.ExportToCSV("./outputs/" + report.AccountName + ".csv")

	// services.CalculateIncome(bg)

	// caller.ConnectionDB()

}
