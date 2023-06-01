package caller

import (
	"log"

	services "github.com/elianaarjona/stunning-pancake/services"
	source "github.com/elianaarjona/stunning-pancake/source"
)

type BankConfig struct {
	BankName  string `json:"bank_name,omitempty"`
	FileType  string `json:"file_type,omitempty"`
	FilePath  string `json:"file_path,omitempty"`
	SheetName string `json:"sheet_name,omitempty"`
}

func (c *BankConfig) ReportBG(rows [][]string, ExcelDataStartRow int) (*source.BgEntries, error) {

	var entries = &source.BgEntries{}

	if c.FileType == "excel" {

		rawData, err := services.ParseExcelFile(rows, ExcelDataStartRow)
		if err != nil {
			log.Fatal()
		}

		for _, raw := range rawData {
			entry, err := services.ProcessRawToEntry(raw, c.FileType)
			if err != nil {
				log.Fatal()
			}

			entries.Entries = append(entries.Entries, entry)
		}
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
