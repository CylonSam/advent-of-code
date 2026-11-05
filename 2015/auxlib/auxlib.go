package auxlib

import (
	"encoding/csv"
	"log"
	"os"
)

type ChallengePart int

const (
	Part1 ChallengePart = 1
	Part2 ChallengePart = 2
)

func ReadCsv(filePath string) [][]string {
	f, err := os.Open(filePath)

	if err != nil {
		log.Fatal("Failed to read csv file", err)
	}
	defer f.Close()

	csvReader := csv.NewReader(f)
	records, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal("Failed to read records from csv", err)
	}

	return records
}
