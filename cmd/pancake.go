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
	TimeLayout                 string = "15/01/2006"
	TimeLayoutMonth            string = "02-Jan-06"
	SheetNameMov               string = "BGPExcelReport"
	SheetNameEst               string = "BGRExcelContReport"
	ExcelDataStartRow          int    = 8
	AccountNameLine            int    = 3
	Template_AccountNameLine   int    = 0
	Template_ExcelDataStartRow int    = 4
)

type FileType int

const (
	Movimiento FileType = iota
	EstadoCuenta
	Template_1
	Template_2
)

func (ft FileType) String() string {
	switch ft {
	case Movimiento:
		return "Movimiento"
	case EstadoCuenta:
		return "EstadoCuenta"
	case Template_1:
		return "Template-1"
	case Template_2:
		return "Template-2"
	default:
		return "Unknown"
	}
}

func getFileType(fileName string) FileType {
	lowerFileName := strings.ToLower(fileName)

	if strings.Contains(lowerFileName, "estado") {
		return EstadoCuenta
	} else if strings.Contains(lowerFileName, "template-1") {
		return Template_1
	} else if strings.Contains(lowerFileName, "template-2") {
		return Template_2
	} else if strings.Contains(lowerFileName, "movimientos") {
		return Movimiento
	} else {
		fmt.Println("No available parsing")
		return -1 // or any other appropriate default value
	}
}

func main() {

	var bg = &caller.BankConfig{
		BankName:  "Banco General",
		FilesPath: "inputs",
		FileType:  "excel",
	}

	files, err := ioutil.ReadDir(bg.FilesPath)
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		if !file.IsDir() {

			startRow := 0

			report := services.ReportIncome{}

			fileDir := fmt.Sprintf("%s/%s", bg.FilesPath, file.Name())
			typeFile := getFileType(fileDir)

			if typeFile.String() == "Movimiento" || typeFile.String() == "Template-1" {
				bg.SheetName = SheetNameMov
			} else {
				bg.SheetName = SheetNameEst
			}

			fmt.Println("Procesing File ", file.Name())

			rows, err := utils.OpenExcelFile(fileDir, bg.SheetName)
			if err != nil {
				fmt.Println("error opening file ", fileDir)
				log.Fatal()
			}

			if typeFile.String() == "Movimiento" {
				report.AccountName = utils.GetAccountNameMov(rows[AccountNameLine])
				startRow = ExcelDataStartRow

			} else if typeFile.String() == "Template-1" || typeFile.String() == "Template-2" {
				report.AccountName = utils.GetAccountNameTemplate(rows[Template_AccountNameLine])
				startRow = Template_ExcelDataStartRow

			} else {
				report.AccountName = utils.GetAccountNameEst(rows[AccountNameLine])
				startRow = ExcelDataStartRow
			}

			reports, err := bg.ReportBG(rows, startRow, typeFile.String())
			if err != nil {
				fmt.Println("error procesing file:", file.Name())
				log.Fatal()
			}

			fmt.Println("Saving Files")
			report.Total = services.GetTotals(reports)
			report.Report = services.GetIncomeTops(reports)

			report.ExportToExcel("./outputs/" + report.AccountName + "_" + file.Name())

		}
	}

}
