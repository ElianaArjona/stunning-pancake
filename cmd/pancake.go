package main

import (
	"pancake/caller"
	"pancake/services"
)

func main() {

	// bg := &services.BgEntries{}
	// bg.ParseTxtFile("./source/bg-sample.txt")
	// services.ParseExcelFile("source/bg-excel-sample.xlsx")

	var bg = &services.BankConfig{
		BankName: "Banco General",
		FilePath: "source/bg-excel-sample.xlsx",
		FileType: "excel",
	}

	caller.ReportBG(bg)

	// services.CalculateIncome(bg)

	// caller.ConnectionDB()

}
