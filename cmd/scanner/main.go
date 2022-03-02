package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	// parse cellosaurus.txt - version 40
	// raw data starts on line 55
	scanRawCellData(55,
		getFilePath("cellosaurus", "cellosaurus.txt"),
		getFilePath("data", "cells.csv"),
		getFilePath("data", "cell_attributes.csv"))

	// parse cellosaurus_refs.txt - version 40
	// raw data starts on line 38
	scanRawRefData(38,
		getFilePath("cellosaurus", "cellosaurus_refs.txt"),
		getFilePath("data", "refs.csv"),
		getFilePath("data", "ref_attributes.csv"))

	// stats from cellosaurus_relnotes.txt - version 40
	// manually entered for simplicity below
	scanRelNoteStats(getFilePath("data", "statistics.csv"))
}

type Cell struct {
	Identifier string
	Accession  string
	Secondary  string
	Synonyms   string
	Sex        string
	Age        string
	Category   string
	Date       string
}

// Reads and parses cellosaurus.txt.
// Writes parsed data to csv file(s).
func scanRawCellData(lineStart int, sourceFile string, destFiles ...string) {
	if len(destFiles) < 2 {
		log.Fatal("Error: need at least two destination file paths")
	}

	txt, err := os.Open(sourceFile)
	if err != nil {
		log.Fatal(err)
	}
	defer txt.Close()

	csvCells, err := os.Create(destFiles[0])
	if err != nil {
		log.Fatal(err)
	}
	defer csvCells.Close()

	csvAttrs, err := os.Create(destFiles[1])
	if err != nil {
		log.Fatal(err)
	}
	defer csvAttrs.Close()

	scanner := bufio.NewScanner(txt)
	writerCells := bufio.NewWriter(csvCells)
	writerAttrs := bufio.NewWriter(csvAttrs)

	if _, err := writerCells.WriteString(
		"\"identifier\",\"accession\",\"secondary\",\"synonyms\",\"sex\",\"age\",\"category\",\"date\"\n",
	); err != nil {
		log.Fatal(err)
	}
	if _, err := writerAttrs.WriteString("\"accession\",\"attribute\",\"content\"\n"); err != nil {
		log.Fatal(err)
	}

	start := 1
	cell := Cell{}
	code := ""
	value := ""
	for scanner.Scan() {
		if start < lineStart {
			start = start + 1
			continue
		}

		line := scanner.Text()
		dict := strings.Split(line, "   ")
		code = dict[0]
		if len(dict) > 1 {
			value = dict[1]
		}

		switch code {
		case "ID":
			cell.Identifier = value
		case "AC":
			cell.Accession = value
		case "AS":
			cell.Secondary = value
		case "SY":
			cell.Synonyms = value
		case "SX":
			cell.Sex = value
		case "AG":
			cell.Age = value
		case "CA":
			cell.Category = value
		case "DT":
			cell.Date = value
		case "//":
			if _, err := writerCells.WriteString(
				csvSprintf(8,
					cell.Identifier,
					cell.Accession,
					cell.Secondary,
					cell.Synonyms,
					cell.Sex,
					cell.Age,
					cell.Category,
					cell.Date),
			); err != nil {
				log.Fatal(err)
			}
			cell = Cell{}
		default:
			if _, err := writerAttrs.WriteString(
				csvSprintf(3, cell.Accession, code, value),
			); err != nil {
				log.Fatal(err)
			}
		}
	}

	writerCells.Flush()
	writerAttrs.Flush()

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

// Reads and parses cellosaurus_refs.txt.
// Writes parsed data to csv file(s).
func scanRawRefData(lineStart int, sourceFile string, destFiles ...string) {
	if len(destFiles) < 2 {
		log.Fatal("Error: need at least two destination file paths")
	}

	txt, err := os.Open(sourceFile)
	if err != nil {
		log.Fatal(err)
	}
	defer txt.Close()

	csvCells, err := os.Create(destFiles[0])
	if err != nil {
		log.Fatal(err)
	}
	defer csvCells.Close()

	csvAttrs, err := os.Create(destFiles[1])
	if err != nil {
		log.Fatal(err)
	}
	defer csvAttrs.Close()

	scanner := bufio.NewScanner(txt)
	writerCells := bufio.NewWriter(csvCells)
	writerAttrs := bufio.NewWriter(csvAttrs)

	if _, err := writerCells.WriteString("\"identifier\",\"citation\"\n"); err != nil {
		log.Fatal(err)
	}
	if _, err := writerAttrs.WriteString("\"identifier\",\"attribute\",\"content\"\n"); err != nil {
		log.Fatal(err)
	}

	start := 1
	identifier := ""
	citation := ""
	code := ""
	value := ""
	for scanner.Scan() {
		if start < lineStart {
			start = start + 1
			continue
		}

		line := scanner.Text()
		dict := strings.Split(line, "   ")
		code = dict[0]
		if len(dict) > 1 {
			value = dict[1]
		}

		switch code {
		case "RX":
			identifier = value
		case "RL":
			citation = value
		case "//":
			if _, err := writerCells.WriteString(
				csvSprintf(2, identifier, citation),
			); err != nil {
				log.Fatal(err)
			}
			identifier = ""
			citation = ""
		default:
			if _, err := writerAttrs.WriteString(
				csvSprintf(3, identifier, code, value),
			); err != nil {
				log.Fatal(err)
			}
		}
	}

	writerCells.Flush()
	writerAttrs.Flush()

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

// Manually writes stats data from cellosaurus_relnotes.txt to csv.
func scanRelNoteStats(destFile string) {
	csv, err := os.Create(destFile)
	if err != nil {
		log.Fatal(err)
	}
	defer csv.Close()

	writer := bufio.NewWriter(csv)
	if _, err := writer.WriteString("\"\",\"attribute\",\"count\"\n"); err != nil {
		log.Fatal(err)
	}
	if _, err := writer.WriteString("1,\"totalCells\",\"134839\"\n"); err != nil {
		log.Fatal(err)
	}
	if _, err := writer.WriteString("2,\"humanCellLines\",\"101276\"\n"); err != nil {
		log.Fatal(err)
	}
	if _, err := writer.WriteString("3,\"mouseCellLines\",\"22999\"\n"); err != nil {
		log.Fatal(err)
	}
	if _, err := writer.WriteString("\"ratCellLines\",\"2498\"\n"); err != nil {
		log.Fatal(err)
	}
	if _, err := writer.WriteString("\"species\",\"747\"\n"); err != nil {
		log.Fatal(err)
	}
	if _, err := writer.WriteString("\"synonyms\",\"96745\"\n"); err != nil {
		log.Fatal(err)
	}
	if _, err := writer.WriteString("\"crossReferences\",\"396097\"\n"); err != nil {
		log.Fatal(err)
	}
	if _, err := writer.WriteString("\"references\",\"138234\"\n"); err != nil {
		log.Fatal(err)
	}
	if _, err := writer.WriteString("\"distinctPublications\",\"23257\"\n"); err != nil {
		log.Fatal(err)
	}
	if _, err := writer.WriteString("\"webLinks\",\"23257\"\n"); err != nil {
		log.Fatal(err)
	}
	if _, err := writer.WriteString("\"cellLinesWithStrProfiles\",\"8032\"\n"); err != nil {
		log.Fatal(err)
	}

	writer.Flush()
}

// Returns formatted string for csv file lines.
func csvSprintf(placeholders int, words ...string) string {
	if placeholders == 2 {
		return fmt.Sprintf("\"%s\",\"%s\"\n", words[0], words[1])
	}
	if placeholders == 3 {
		return fmt.Sprintf("\"%s\",\"%s\",\"%s\"\n", words[0], words[1], words[2])
	}
	if placeholders == 8 {
		return fmt.Sprintf(
			"\"%s\",\"%s\",\"%s\",\"%s\",\"%s\",\"%s\",\"%s\",\"%s\"\n",
			words[0],
			words[1],
			words[2],
			words[3],
			words[4],
			words[5],
			words[6],
			words[7],
		)
	}
	return ""
}

// Returns the absolute path to read/write file.
func getFilePath(dir string, file string) string {
	root, err := filepath.Abs("../")
	if err != nil {
		log.Fatal(err)
	}
	return fmt.Sprintf("%s/%s/%s", root, dir, file)
}
