package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	filePath := "/home/samlas/workspace/advent-of-code/2015/02/input.csv"
	inputs := readCsv(filePath)

	requiredWrappingPaperTotal, requiredRibbonTotal := 0, 0
	for index := range len(inputs) {
		input := inputs[index][0]
		//fmt.Println(input)
		l, w, h := parseDimensions(input)

		wrappingPaper, ribbon := calculateMaterial(l, w, h)

		requiredWrappingPaperTotal += wrappingPaper
		requiredRibbonTotal += ribbon
	}

	fmt.Printf("Required wrapping paper: %d\n", requiredWrappingPaperTotal)
	fmt.Printf("Required ribbon: %d", requiredRibbonTotal)
}

func readCsv(filePath string) [][]string {
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

func parseDimensions(record string) (int, int, int) {
	dimensions := strings.Split(record, "x")
	length, _ := strconv.Atoi(dimensions[0])
	width, _ := strconv.Atoi(dimensions[1])
	height, _ := strconv.Atoi(dimensions[2])

	return length, width, height
}

func calculateMaterial(length int, width int, height int) (int, int) {
	dimensions := []int{length, width, height}
	slices.Sort(dimensions)
	//fmt.Println(sides)

	surfaceArea := (2 * length * height) + (2 * length * width) + (2 * width * height)

	requiredWrappingPaper := surfaceArea + (dimensions[0] * dimensions[1])
	requiredRibbon := 2*(dimensions[0]+dimensions[1]) + (length * width * height)
	return requiredWrappingPaper, requiredRibbon
}
