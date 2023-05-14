package services

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	excel "github.com/xuri/excelize/v2"
)

//identificar las 3 de valor mas alto y con mayor coincidencia

type BankConfig struct {
	BankName string `json:"bank_name,omitempty"`
	FileType string `json:"file_type,omitempty"`
	FilePath string `json:"file_path,omitempty"`
}

func (c *BankConfig) ParseTxtFile() ([]*RawEntry, error) {

	var entries = []*RawEntry{}

	fileLines, reader, err := ReadTxtFileLines(c.FilePath)
	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		// Close the spreadsheet.
		if err := reader.Close(); err != nil {
			fmt.Println(err)
		}
	}()

	//Drop header
	header := false

	//Parse lines to JSON
	for fileLines.Scan() {

		if header {
			data := strings.Split(fileLines.Text(), ";")

			// Each Row is a new Entry
			entry := &RawEntry{}

			if data[0] == "" {
				break
			}

			entry.Name = "Banco General"

			entry.TransactionDate = data[0]

			entry.Reference, err = strconv.ParseInt(data[1], 10, 64)
			if err != nil {
				return nil, err
			}

			entry.Description = data[2]

			entry.Expense, err = strconv.ParseFloat(data[3], 64)
			if err != nil {
				return nil, err
			}
			entry.Expense = 0

			entry.Income, err = strconv.ParseFloat(data[4], 64)
			if err != nil {
				return nil, err
			}
			entry.Income = 0

			entry.Balance, err = strconv.ParseFloat(data[5], 64)
			if err != nil {
				return nil, err
			}

			entry.Balance = 0

			entries = append(entries, entry)

		}

		header = true

	}

	// jsonData, err := json.MarshalIndent(bg, "", "  ")
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// fmt.Println(string(jsonData))

	return entries, nil
}
func (c *BankConfig) ParseExcelFile() ([]*RawEntry, error) {

	var entries = []*RawEntry{}

	f, err := excel.OpenFile(c.FilePath)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	defer func() {
		// Close the spreadsheet.
		if err := f.Close(); err != nil {
			fmt.Println(err)
		}
	}()

	// Get all the rows in the selected Sheet.
	rows, err := f.GetRows(sheetName)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	for i, row := range rows {
		//Skip unsed excel rows
		if i >= excelDataStartRow {

			// Each Row is a new Entry
			entry := &RawEntry{}

			entry.Name = "Banco General"

			// tm, err := BuildTimestampExcel(row[1])
			// if err != nil {
			// 	log.Fatal("Error Date Excel Date Parsing")
			// }
			entry.TransactionDate = row[1]
			// fmt.Println(entry.TransactionDate, "ts date")

			entry.Description = row[3]

			// fmt.Println(row[4])

			money, err := strconv.ParseFloat(row[4], 64)
			if err != nil {
				return nil, err
			}

			if money < 0 {

				entry.Expense = money
				entry.Income = 0.0

			} else {
				entry.Income = money
				entry.Expense = 0.0

			}

			entry.Balance, err = strconv.ParseFloat(row[5], 64)
			if err != nil {
				return nil, err
			}

			entries = append(entries, entry)
		}

	}

	// jsonData, err := json.MarshalIndent(bg, "", "  ")
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// fmt.Println(string(jsonData))

	return entries, nil
}

func (mc *BgServicesTypesTxt) GetServicesType(r *RawEntry) (*RawEntry, error) {

	// mapCodes := &BgServicesTypes{
	// 	Interest: 0,
	// 	ACH:      93,
	// 	Yappy:    []int64{452, 456},
	// 	Service:  []int64{105, 256},
	// }

	// mc.Interest = 0
	mc.ACH = 93
	mc.Yappy = []int64{452, 456}
	mc.Service = []int64{105, 256}

	// if r.Reference == mc.Interest {
	// 	r.Type = "Interest"
	// } else
	if r.Reference == mc.ACH {
		r.Type = "ACH"
	} else if intInSlice(r.Reference, mc.Yappy) {
		r.Type = "Yappy"
	} else if intInSlice(r.Reference, mc.Service) {
		r.Type = "Service"
	}

	return r, nil
}

func (mc *BgServicesTypesExcel) GetServicesType(r *RawEntry) string {

	// mc.Interest = "interest"
	mc.ACH = "ach"
	mc.Yappy = "yappy"
	mc.Deposito = "deposito"
	mc.Service = []string{"naturgy", "recargas", "cable onda"}

	var typeService string

	r.Description = strings.ToLower(r.Description)

	if strings.Contains(r.Description, mc.Deposito) {
		r.Type = "Deposito"
	} else if strings.Contains(r.Description, mc.ACH) {
		typeService = "ACH"
	} else if strings.Contains(r.Description, mc.Yappy) {
		typeService = "Yappy"
	}
	for _, s := range mc.Service {
		if strings.Contains(r.Description, s) {
			typeService = "Service"
			return typeService
		}
	}

	return typeService
}

func (r *RawEntry) ProcessRawToEntry(c *BankConfig) (*Entry, error) {

	var e = &Entry{}

	e.Bank = r.Name

	e.Income = r.Income
	e.Expense = r.Expense
	e.Balance = r.Balance
	e.Description = r.Description

	if c.FileType == "txt" {
		var mc = &BgServicesTypesTxt{}
		mc.GetServicesType(r)

		time, err := BuildTimestamp(r.TransactionDate)
		if err != nil {
			fmt.Println("Error Time")
			return nil, err
		}

		e.Month = int64(time.Month())
		e.Year = int64(time.Year())

	}

	if c.FileType == "excel" {
		var mc = &BgServicesTypesExcel{}
		typeService := mc.GetServicesType(r)
		e.Type = typeService

		tm, err := BuildTimestampExcel(r.TransactionDate)
		if err != nil {
			fmt.Println("Error Time Parse Excel")
			return nil, err
		}

		e.Month = int64(tm.Month())
		e.Year = int64(tm.Year())

	}

	return e, nil

}

func GetTopTransference() {

}
