package caller

import (
	"fmt"
	"log"

	"github.com/elianaarjona/stunning-pancake/services"
	source "github.com/elianaarjona/stunning-pancake/source"
)

type BankConfig struct {
	BankName  string `json:"bank_name,omitempty"`
	FileType  string `json:"file_type,omitempty"`
	FilesPath string `json:"file_path,omitempty"`
	SheetName string `json:"sheet_name,omitempty"`
}

func (c *BankConfig) ReportBG(rows [][]string, ExcelDataStartRow int, fileType string) (*source.BgEntries, error) {

	var entries = &source.BgEntries{}
	var rawData []*source.RawEntry
	var err error

	if c.FileType == "excel" {
		switch fileType {
		case "Movimiento":
			rawData, err = services.ParseExceMovimientos(rows, ExcelDataStartRow)
			if err != nil {
				fmt.Println(err)
				log.Fatal()
			}
		case "Template-1":
			rawData, err = services.ParseExcelTemplate1(rows, ExcelDataStartRow)
			if err != nil {
				fmt.Println(err)
				log.Fatal()
			}
		case "Template-2":
			rawData, err = services.ParseExcelTemplate2(rows, ExcelDataStartRow)
			if err != nil {
				fmt.Println(err)
				log.Fatal()
			}
		case "EstadoCuenta":
			rawData, err = services.ParseExcelEstadosCuenta(rows, ExcelDataStartRow)
			if err != nil {
				fmt.Println(err)
				log.Fatal()
			}
		default:
			rawData = nil // or any other appropriate default value
			fmt.Println("Unknown file type")
			return nil, err
		}
	}

	for _, raw := range rawData {
		entry, err := services.ProcessRawToEntry(raw, c.FileType)
		if err != nil {
			log.Fatal()
		}

		entries.Entries = append(entries.Entries, entry)
	}

	// jsonData, err := json.MarshalIndent(entries, "", "  ")
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// fmt.Println(string(jsonData))

	// Create or open the output file
	// file, _ := os.Create("output.json")
	// defer file.Close()
	// // Write the JSON data to the file
	// _, _ = file.Write(jsonData)

	return entries, nil

}
