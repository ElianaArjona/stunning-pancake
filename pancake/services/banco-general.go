package services

import (
	"encoding/json"
	"fmt"
	"log"
	"strconv"
	"strings"
)

type BancoGeneralTxt struct {
	Entries []*BgEntry `json:"entries,omitempty"`
}

type BgEntry struct {
	TransactionDate string  `json:"transaction_date,omitempty"`
	Reference       string  `json:"reference,omitempty"`
	Description     string  `json:"description,omitempty"`
	Expenses        float64 `json:"expenses,omitempty"`
	Income          float64 `json:"income,omitempty"`
	Balance         float64 `json:"balance,omitempty"`
}

func (bg *BancoGeneralTxt) ParseTxtFile(filePath string) (i interface{}, err error) {

	//Drop header
	fileLines, reader, err := ReadTxtFileLines(filePath)
	if err != nil {
		log.Fatal(err)
	}

	var entries []*BgEntry
	header := false

	//Parse lines to JSON
	for fileLines.Scan() {

		if header {
			data := strings.Split(fileLines.Text(), ";")

			entry := &BgEntry{}

			if data[0] == "" {
				break
			}

			entry.TransactionDate = data[0]
			entry.Reference = data[1]
			entry.Description = data[2]

			entry.Expenses, err = strconv.ParseFloat(data[3], 64)
			if err != nil {
				entry.Expenses = 0
			}

			entry.Income, err = strconv.ParseFloat(data[4], 64)
			if err != nil {
				entry.Income = 0
			}

			entry.Balance, err = strconv.ParseFloat(data[5], 64)
			if err != nil {
				entry.Balance = 0
			}

			entries = append(entries, entry)

		}

		header = true

	}

	bg.Entries = entries

	jsonData, err := json.MarshalIndent(bg, "", "  ")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(jsonData))

	reader.Close()

	return jsonData, nil
}
