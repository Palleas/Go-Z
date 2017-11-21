package main

import (
	"fmt"
	"os"
	"strings"
)

// Item represents an Item in the master catalog
type Item struct {
	diamondCode, stockCode, title, publisher string
}

func parseLine(line string) Item {
	row := strings.Split(line, "\t")

	return Item{row[0], row[1], row[4], row[13]}
}

func main() {
	const path = "/Users/romain/Downloads/MasterDataFile-ITEMS_201710.csv"
	file, error := os.Open(path)
	if error != nil {
		fmt.Printf("An error occured: %s\n", error)
	}
	defer file.Close()

	// get the file size
	stat, err := file.Stat()
	if err != nil {
		return
	}
	fmt.Printf("File size is %d\n", stat.Size())

	// read the file
	bs := make([]byte, stat.Size())
	_, err = file.Read(bs)
	if err != nil {
		return
	}

	str := string(bs)

	lines := strings.Split(str, "\n")
	fmt.Printf("File has %d lines\n", len(lines))

	for index, line := range lines {
		if line == "" {
			continue
		}

		item := parseLine(line)
		fmt.Printf("(%d) Item %s (%s) - %s at %s\n",
			index,
			item.diamondCode,
			item.stockCode,
			item.title,
			item.publisher,
		)
	}
}
