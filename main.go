package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
)

var index int = 0

func showLines(filePath string, isIndexNumber bool) error {
	sf, err := os.Open(filePath)

	if err != nil {
		return err
	}
	defer sf.Close()

	scanner := bufio.NewScanner(sf)

	for scanner.Scan() {
		if isIndexNumber {
			index++
			fmt.Printf("%d: %s\n", index, scanner.Text())
		} else {
			fmt.Println(scanner.Text())
		}
	}

	return nil
}

func main() {
	var isIndexNumber *bool = flag.Bool("n", false, "Show line number")
	flag.Parse()

	fileList := make([]string, 0, 8)

	for _, arg := range os.Args[1:] {
		if arg == "-n" {
			*isIndexNumber = true
			continue
		}

		fileList = append(fileList, arg)
	}

	for _, file := range fileList {
		showLines(file, *isIndexNumber)
	}
}
