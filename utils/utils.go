package services

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	excel "github.com/xuri/excelize/v2"
)

func OpenExcelFile(filepath string, sheetName string) ([][]string, error) {

	f, err := excel.OpenFile(filepath)
	if err != nil {
		fmt.Println(err, "could not open file")
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
		fmt.Println(err, "get rows errors")
		return nil, err
	}

	return rows, nil
}

func ReadTxtFileLines(filePath string) (*bufio.Scanner, *os.File, error) {

	readFile, err := os.Open(filePath)

	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)

	return fileScanner, readFile, nil
}

func BuildTimestamp(source string, excelEpoch time.Time) (time.Time, error) {
	// Attempt to parse the source as a formatted date string
	tm, err := time.Parse("1/2/2006", source)
	if err == nil {
		return tm, nil
	}

	// Attempt to parse the source as a different formatted date string
	tm, err = time.Parse("02-Jan-2006", source)
	if err == nil {
		return tm, nil
	}

	// If parsing as a formatted date string fails, attempt to parse as a float representing the Excel serial number
	excelSerial, err := strconv.ParseFloat(source, 64)
	if err != nil {
		return time.Time{}, err
	}

	// Calculate the duration in days relative to the excelEpoch
	duration := time.Duration(excelSerial-1) * 24 * time.Hour

	// Adjust the timestamp relative to the excelEpoch
	tm = excelEpoch.Add(duration)

	return tm, nil
}

func intInSlice(element int64, list []int64) bool {
	for _, x := range list {
		if x == element {
			return true
		}
	}
	return false
}

func GetAccountNameMov(row []string) string {
	accountName := strings.Replace(row[1], "Cuenta:", "", -1)
	return accountName
}

func GetAccountNameEst(row []string) string {
	accountName := strings.Replace(row[0], "Cuenta:", "", -1)
	return accountName
}

func GetAccountNameTemplate(row []string) string {
	accountName := row[1]
	return accountName
}
