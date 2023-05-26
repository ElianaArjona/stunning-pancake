package main

import (
	"log"

	"github.com/elianaarjona/stunning-pancake/caller"
	"github.com/elianaarjona/stunning-pancake/services"
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
	}

	reports, err := bg.ReportBG()
	if err != nil {
		log.Fatal()
	}

	services.GetIncomeTops(reports)

	// services.CalculateIncome(bg)

	// caller.ConnectionDB()

}
