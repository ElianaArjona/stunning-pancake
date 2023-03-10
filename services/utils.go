package services

import (
	"bufio"
	"fmt"
	"os"
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
