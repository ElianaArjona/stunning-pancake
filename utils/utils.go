package services

import (
	"bufio"
	"fmt"
	"os"
	"time"

	excel "github.com/xuri/excelize/v2"
)

func OpenExcelFile(filepath string, sheetName string) ([][]string, error) {

	f, err := excel.OpenFile(filepath)
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

// func BuildTimestamp(source string, timeLayout string) (time.Time, error) {
// 	tt, err := time.Parse(source, timeLayout)
// 	if err != nil {
// 		return time.Time{}, err
// 	}
// 	return tt, nil
// }

func BuildTimestamp(source string, excelEpoch time.Time) (time.Time, error) {
	fmt.Println(source)

	// Attempt to parse the source as "1/28/2023" format
	tm, err := time.Parse("1/2/2006", source)
	if err != nil {
		// If parsing as "1/28/2023" fails, try parsing as "02-Jan-2006" format
		tm, err = time.Parse("02-Jan-2006", source)
		if err != nil {
			return time.Time{}, err
		}
	}

	// Adjust the timestamp relative to the excelEpoch
	duration := tm.Sub(excelEpoch)
	tm = excelEpoch.Add(duration)

	return tm, nil
}

// func BuildTimestamp(source string, excelEpoch time.Time) (time.Time, error) {
// 	fmt.Println(source)
// 	in, err := strconv.Atoi(source)
// 	if err != nil {
// 		return time.Time{}, err
// 	}

// 	tm := excelEpoch.Add(time.Duration(in * int(24*time.Hour)))

// 	return tm, nil
// }

func intInSlice(element int64, list []int64) bool {
	for _, x := range list {
		if x == element {
			return true
		}
	}
	return false
}
