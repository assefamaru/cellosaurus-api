package main

import (
	"bufio"
	"os"
	"strings"
)

// Cell line attributes that appear exactly once
// (one line per cell) in text file raw data.
// Multiple line entries are parsed in the
// form [{accession, attribute, content}...].
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

// scanCellsTXT scans cellosaurus.txt, writing parsed output to csv file(s).
func scanCellsTXT(firstLineNum int, sourceFile, destFileCells, destFileAttrs string) error {
	sourceTxt, err := os.Open(sourceFile)
	if err != nil {
		return err
	}
	defer sourceTxt.Close()

	cellsCSV, err := os.Create(destFileCells)
	if err != nil {
		return err
	}
	defer cellsCSV.Close()

	attrsCSV, err := os.Create(destFileAttrs)
	if err != nil {
		return err
	}
	defer attrsCSV.Close()

	cellsWriter := bufio.NewWriter(cellsCSV)
	cellsHeader := "\"identifier\",\"accession\",\"secondary\",\"synonyms\",\"sex\",\"age\",\"category\",\"date\"\n"
	if _, err := cellsWriter.WriteString(cellsHeader); err != nil {
		return err
	}

	attrsWriter := bufio.NewWriter(attrsCSV)
	attrsHeader := "\"\",\"accession\",\"attribute\",\"content\"\n"
	if _, err := attrsWriter.WriteString(attrsHeader); err != nil {
		return err
	}

	var cell Cell
	lineNum := 1
	currLineNum := 1
	scanner := bufio.NewScanner(sourceTxt)
	for scanner.Scan() {
		if currLineNum < firstLineNum {
			currLineNum += 1
			continue
		}

		var code, value string
		lineParts := strings.Split(scanner.Text(), "   ")
		code = lineParts[0]
		if len(lineParts) > 1 {
			value = lineParts[1]
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
			out := formattedCSVLine(false, -1, cell.Identifier, cell.Accession, cell.Secondary,
				cell.Synonyms, cell.Sex, cell.Age, cell.Category, cell.Date)
			if _, err := cellsWriter.WriteString(out); err != nil {
				return err
			}
			cell = Cell{}
		default:
			out := formattedCSVLine(true, lineNum, cell.Accession, code, value)
			if _, err := attrsWriter.WriteString(out); err != nil {
				return err
			}
			lineNum += 1
		}
	}

	cellsWriter.Flush()
	attrsWriter.Flush()

	if err := scanner.Err(); err != nil {
		return err
	}

	return nil
}
