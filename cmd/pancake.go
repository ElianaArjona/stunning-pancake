package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"

	"github.com/elianaarjona/stunning-pancake/caller"
	"github.com/elianaarjona/stunning-pancake/services"
	utils "github.com/elianaarjona/stunning-pancake/utils"
)

const (
	TimeLayout        string = "15/01/2006"
	TimeLayoutMonth   string = "02-Jan-06"
	SheetNameMov      string = "BGPExcelReport"
	SheetNameEst      string = "BGRExcelContReport"
	ExcelDataStartRow int    = 8
	AccountNameLine   int    = 3
)

type FileType int

const (
	Movimiento FileType = iota
	EstadoCuenta
)

func (ft FileType) String() string {
	switch ft {
	case Movimiento:
		return "Movimiento"
	case EstadoCuenta:
		return "EstadoCuenta"
	default:
		return "Unknown"
	}
}

func getFileType(fileName string) FileType {
	lowerFileName := strings.ToLower(fileName)

	if strings.Contains(lowerFileName, "estado") {
		return EstadoCuenta
	} else if strings.Contains(lowerFileName, "movimientos") {
		return Movimiento
	} else {
		fmt.Println("No available parsing")
		return -1 // or any other appropriate default value
	}
}

func main() {

	// bg := &services.BgEntries{}
	// bg.ParseTxtFile("./source/bg-sample.txt")
	// services.ParseExcelFile("source/bg-excel-sample.xlsx")

	var bg = &caller.BankConfig{
		BankName:  "Banco General",
		FilesPath: "samples",
		FileType:  "excel",
	}

	files, err := ioutil.ReadDir(bg.FilesPath)
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		if !file.IsDir() {

			report := services.ReportIncome{}

			fileDir := fmt.Sprintf("%s/%s", bg.FilesPath, file.Name())
			typeFile := getFileType(fileDir)

			if typeFile.String() == "Movimiento" {
				bg.SheetName = SheetNameMov
			} else {
				bg.SheetName = SheetNameEst
			}

			fmt.Println("Procesing File ", file.Name())

			rows, err := utils.OpenExcelFile(fileDir, bg.SheetName)
			if err != nil {
				fmt.Errorf("error opening file ", fileDir)
				log.Fatal()
			}

			if typeFile.String() == "Movimiento" {
				report.AccountName = utils.GetAccountNameMov(rows[AccountNameLine])
			} else {
				report.AccountName = utils.GetAccountNameEst(rows[AccountNameLine])
			}

			reports, err := bg.ReportBG(rows, ExcelDataStartRow, typeFile.String())
			if err != nil {
				fmt.Errorf("error procesing file ", file.Name())
				log.Fatal()
			}

			fmt.Println("Saving Files")
			report.Total = services.GetTotals(reports)
			report.Report = services.GetIncomeTops(reports)

			// report.ExportToExcel("./outputs/sample.xlsx")
			report.ExportToExcel("./outputs/" + file.Name())

			// services.CalculateIncome(bg)

			// caller.ConnectionDB()
		}
	}

}
