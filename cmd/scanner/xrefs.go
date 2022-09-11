package main

import (
	"bufio"
	"os"
	"strings"
)

// XRef raw data structure.
type XRef struct {
	Abbrev string
	Name   string
	Server string
	URL    string
	Term   string
	Cat    string
}

// scanXRefsTXT scans cellosaurus_xrefs.txt, writing parsed output to csv file(s).
func scanXRefsTXT(firstLineNum int, sourceFile, destFile string) error {
	sourceTXT, err := os.Open(sourceFile)
	if err != nil {
		return err
	}
	defer sourceTXT.Close()

	xrefsCSV, err := os.Create(destFile)
	if err != nil {
		return err
	}
	defer xrefsCSV.Close()

	xrefsWriter := bufio.NewWriter(xrefsCSV)
	if _, err := xrefsWriter.WriteString("\"\",\"abbrev\",\"name\",\"server\",\"url\",\"term\",\"cat\"\n"); err != nil {
		return err
	}

	var xref XRef
	lineNum := 1
	currLineNum := 1
	scanner := bufio.NewScanner(sourceTXT)
	for scanner.Scan() {
		if currLineNum < firstLineNum {
			currLineNum += 1
			continue
		}

		var code, value string
		lineParts := strings.Split(scanner.Text(), ": ")
		code = lineParts[0]
		if len(lineParts) > 1 {
			value = lineParts[1]
		}

		switch code {
		case "Abbrev":
			xref.Abbrev = value
		case "Name  ":
			xref.Name = value
		case "Server":
			xref.Server = value
		case "Db_URL":
			xref.URL = value
		case "Term. ":
			xref.Term = value
		case "Cat   ":
			xref.Cat = value
		case "//":
			out := formattedCSVLine(true, lineNum, xref.Abbrev, xref.Name, xref.Server, xref.URL, xref.Term, xref.Cat)
			if _, err := xrefsWriter.WriteString(out); err != nil {
				return err
			}
			lineNum += 1
			xref = XRef{}
		}
	}

	xrefsWriter.Flush()

	if err := scanner.Err(); err != nil {
		return err
	}

	return nil
}
