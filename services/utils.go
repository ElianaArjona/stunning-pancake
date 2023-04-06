package services

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"time"
)

func ReadTxtFileLines(filePath string) (*bufio.Scanner, *os.File, error) {

	readFile, err := os.Open(filePath)

	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)

	return fileScanner, readFile, nil
}

func BuildTimestamp(source string) (time.Time, error) {
	tt, err := time.Parse(source, timeLayout)
	if err != nil {
		return time.Time{}, err
	}
	return tt, nil
}

func BuildTimestampExcel(source string) (time.Time, error) {
	in, err := strconv.Atoi(source)
	if err != nil {
		return time.Time{}, err
	}

	tm := excelEpoch.Add(time.Duration(in * int(24*time.Hour)))

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
