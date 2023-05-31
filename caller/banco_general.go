package caller

import (
	"encoding/json"
	"log"
	"os"

	services "github.com/elianaarjona/stunning-pancake/services"
	source "github.com/elianaarjona/stunning-pancake/source"
	utils "github.com/elianaarjona/stunning-pancake/utils"
)

type BankConfig struct {
	BankName string `json:"bank_name,omitempty"`
	FileType string `json:"file_type,omitempty"`
	FilePath string `json:"file_path,omitempty"`
}

func (c *BankConfig) ReportBG() (*source.BgEntries, error) {

	var entries = &source.BgEntries{}

	if c.FileType == "excel" {

		rows, err := utils.OpenExcelFile(c.FilePath, source.SheetName)
		if err != nil {
			log.Fatal()
		}

		rawData, err := services.ParseExcelFile(rows)
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

	jsonData, err := json.MarshalIndent(entries, "", "  ")
	if err != nil {
		log.Fatal(err)
	}

	// fmt.Println(string(jsonData))

	// Create or open the output file
	file, _ := os.Create("output.json")
	defer file.Close()
	// Write the JSON data to the file
	_, _ = file.Write(jsonData)

	return entries, nil

}
