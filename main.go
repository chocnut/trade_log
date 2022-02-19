package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {

	records, err := readData("trade.csv")

	if err != nil {
		log.Fatal(err)
	}

	for _, record := range records {

		// Check data and option records only
		if record[0] == "DATA" && record[6] == "OPT" {
			description := strings.Split(record[8], " ")

			t := Trade{
				symbol:  description[0],
				expiry:  description[1],
				strike:  description[2],
				putCall: description[3],
				buySell: record[42],
			}

			fmt.Printf("%s %s %s %s %s\n", t.symbol, t.expiry, t.strike, t.putCall, t.buySell)
		}
	}

}

func readData(fileName string) ([][]string, error) {
	f, err := os.Open(fileName)

	if err != nil {
		return [][]string{}, err
	}

	defer f.Close()

	csvReader := csv.NewReader(f)

	// Skip length check of columns
	csvReader.FieldsPerRecord = -1
	records, err := csvReader.ReadAll()

	if err != nil {
		return [][]string{}, err
	}

	return records[3:], nil
}
