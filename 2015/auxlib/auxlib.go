package auxlib

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
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

func ReadLinesTextFile(filePath string) []string {
	f, err := os.Open(filePath)

	if err != nil {
		log.Fatal("Failed to read csv file", err)
	}
	defer f.Close()

	lines := []string{}

	scanner := bufio.NewReader(f)

	for {
		textLine, err := scanner.ReadString('\n')
		if err == io.EOF {
			if len(textLine) != 0 {
				lines = append(lines, textLine)
			}
			break
		}
		if err != nil {
			fmt.Println(fmt.Errorf("error reading from file: %w", err))
		}
		lines = append(lines, textLine)
	}
	return lines

}
