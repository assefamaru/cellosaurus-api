package main

import (
	"bufio"
	"log"
	"os"
	"strings"
)

// scanRefsTXT scans cellosaurus_refs.txt, writing parsed output to csv file(s).
func scanRefsTXT(firstLineNum int, sourceFile, destFileCells, destFileAttrs string) error {
	sourceTXT, err := os.Open(sourceFile)
	if err != nil {
		return err
	}
	defer sourceTXT.Close()

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
	if _, err := cellsWriter.WriteString("\"\",\"identifier\",\"citation\"\n"); err != nil {
		return err
	}

	attrsWriter := bufio.NewWriter(attrsCSV)
	if _, err := attrsWriter.WriteString("\"\",\"identifier\",\"attribute\",\"content\"\n"); err != nil {
		return err
	}

	var identifier, citation string
	lineNumRef := 1
	lineNumAttr := 1
	currLineNum := 1
	scanner := bufio.NewScanner(sourceTXT)
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
		case "RX":
			identifier = value
		case "RL":
			citation = value
		case "//":
			out := formattedCSVLine(true, lineNumRef, identifier, citation)
			if _, err := cellsWriter.WriteString(out); err != nil {
				return err
			}
			identifier = ""
			citation = ""
			lineNumRef += 1
		default:
			out := formattedCSVLine(true, lineNumAttr, identifier, code, value)
			if _, err := attrsWriter.WriteString(out); err != nil {
				return err
			}
			lineNumAttr += 1
		}
	}

	cellsWriter.Flush()
	attrsWriter.Flush()

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return nil
}
