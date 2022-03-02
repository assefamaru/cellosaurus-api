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
	scanRawCellData(
		55,
		getFilePath("cellosaurus", "cellosaurus.txt"),
		getFilePath("data", "cells.csv"),
		getFilePath("data", "cell_attributes.csv"),
	)

	// parse cellosaurus_refs.txt - version 40
	// raw data starts on line 38
	scanRawRefData(
		38,
		getFilePath("cellosaurus", "cellosaurus_refs.txt"),
		getFilePath("data", "refs.csv"),
		getFilePath("data", "ref_attributes.csv"),
	)

	// parse cellosaurus_xrefs.txt - version 40
	// raw data starts on line 118
	scanRawCrossRefData(
		118,
		getFilePath("cellosaurus", "cellosaurus_xrefs.txt"),
		getFilePath("data", "xrefs.csv"),
	)

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
	if _, err := writerAttrs.WriteString("\"\",\"accession\",\"attribute\",\"content\"\n"); err != nil {
		log.Fatal(err)
	}

	start := 1
	lineNumber := 1
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
				csvSprintf(
					false,
					-1,
					cell.Identifier,
					cell.Accession,
					cell.Secondary,
					cell.Synonyms,
					cell.Sex,
					cell.Age,
					cell.Category,
					cell.Date,
				),
			); err != nil {
				log.Fatal(err)
			}
			cell = Cell{}
		default:
			if _, err := writerAttrs.WriteString(
				csvSprintf(true, lineNumber, cell.Accession, code, value),
			); err != nil {
				log.Fatal(err)
			}
			lineNumber = lineNumber + 1
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

	if _, err := writerCells.WriteString("\"\",\"identifier\",\"citation\"\n"); err != nil {
		log.Fatal(err)
	}
	if _, err := writerAttrs.WriteString("\"\",\"identifier\",\"attribute\",\"content\"\n"); err != nil {
		log.Fatal(err)
	}

	start := 1
	lineNumberRef := 1
	lineNumberAttr := 1
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
				csvSprintf(true, lineNumberRef, identifier, citation),
			); err != nil {
				log.Fatal(err)
			}
			identifier = ""
			citation = ""
			lineNumberRef = lineNumberRef + 1
		default:
			if _, err := writerAttrs.WriteString(
				csvSprintf(true, lineNumberAttr, identifier, code, value),
			); err != nil {
				log.Fatal(err)
			}
			lineNumberAttr = lineNumberAttr + 1
		}
	}
	writerCells.Flush()
	writerAttrs.Flush()
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

type XRef struct {
	Abbrev string
	Name   string
	Server string
	URL    string
	Term   string
	Cat    string
}

func scanRawCrossRefData(lineStart int, sourceFile string, destFile string) {
	txt, err := os.Open(sourceFile)
	if err != nil {
		log.Fatal(err)
	}
	defer txt.Close()

	csv, err := os.Create(destFile)
	if err != nil {
		log.Fatal(err)
	}
	defer csv.Close()

	scanner := bufio.NewScanner(txt)
	writer := bufio.NewWriter(csv)

	if _, err := writer.WriteString(
		"\"\",\"abbrev\",\"name\",\"server\",\"url\",\"term\",\"cat\"\n",
	); err != nil {
		log.Fatal(err)
	}

	start := 1
	lineNumber := 1
	code := ""
	value := ""
	xRef := XRef{}
	for scanner.Scan() {
		if start < lineStart {
			start = start + 1
			continue
		}

		line := scanner.Text()
		dict := strings.Split(line, ": ")
		code = dict[0]
		if len(dict) > 1 {
			value = dict[1]
		}

		switch code {
		case "Abbrev":
			xRef.Abbrev = value
		case "Name  ":
			if strings.HasPrefix(value, "Istituto Zooprofilattico") { // sanitize special char from raw data
				value = "Istituto Zooprofilattico Sperimentale della Lombardia e dell Emilia Romagna biobank"
			}
			xRef.Name = value
		case "Server":
			xRef.Server = value
		case "Db_URL":
			xRef.URL = value
		case "Term. ":
			xRef.Term = value
		case "Cat   ":
			xRef.Cat = value
		case "//":
			if _, err := writer.WriteString(
				csvSprintf(
					true,
					lineNumber,
					xRef.Abbrev,
					xRef.Name,
					xRef.Server,
					xRef.URL,
					xRef.Term,
					xRef.Cat,
				),
			); err != nil {
				log.Fatal(err)
			}
			lineNumber = lineNumber + 1
			xRef = XRef{}
		}
	}
	writer.Flush()
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

// Manually writes stats data from cellosaurus_relnotes.txt.
func scanRelNoteStats(destFile string) {
	csv, err := os.Create(destFile)
	if err != nil {
		log.Fatal(err)
	}
	defer csv.Close()
	writer := bufio.NewWriter(csv)
	if _, err := writer.WriteString(
		"\"\",\"attribute\",\"count\"\n" +
			"1,\"totalCellLines\",\"134839\"\n" +
			"2,\"humanCellLines\",\"101276\"\n" +
			"3,\"mouseCellLines\",\"22999\"\n" +
			"4,\"ratCellLines\",\"2498\"\n" +
			"5,\"species\",\"747\"\n" +
			"6,\"synonyms\",\"96745\"\n" +
			"7,\"crossReferences\",\"396097\"\n" +
			"8,\"references\",\"138234\"\n" +
			"9,\"distinctPublications\",\"23257\"\n" +
			"10,\"webLinks\",\"23257\"\n" +
			"11,\"cellLinesWithStrProfiles\",\"8032\"\n" +
			"12,\"version\",\"40\"\n",
	); err != nil {
		log.Fatal(err)
	}
	writer.Flush()
}

// Returns formatted string for csv file lines.
func csvSprintf(addLineNumber bool, lineNumber int, words ...string) string {
	prefix := ""
	if addLineNumber {
		prefix = fmt.Sprintf("%d,", lineNumber)
	}
	final := ""
	for _, word := range words {
		final = final + fmt.Sprintf("\"%s\",", word)
	}
	return prefix + strings.TrimSuffix(final, ",") + "\n"
}

// Returns the absolute path to read/write file.
func getFilePath(dir string, file string) string {
	root, err := filepath.Abs("../")
	if err != nil {
		log.Fatal(err)
	}
	return fmt.Sprintf("%s/%s/%s", root, dir, file)
}
