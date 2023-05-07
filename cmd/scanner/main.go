package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/assefamaru/cellosaurus-api/pkg/logging"
)

const (
	latestVersion = "42"
)

type versionConfig struct {
	startCellsTXT int // the line where raw data starts in cellosaurus.txt
	startRefsTXT  int // the line where raw data starts in cellosaurus_refs.txt
	startXRefsTXT int // the line where raw data starts in cellosaurus_xrefs.txt
}

var settings = map[string]*versionConfig{
	"41": {startCellsTXT: 55, startRefsTXT: 38, startXRefsTXT: 118},
	"42": {startCellsTXT: 55, startRefsTXT: 38, startXRefsTXT: 118},
}

func main() {
	version := flag.String("version", fromEnvOrDefaultVal("VERSION", latestVersion), "The current Cellosaurus data version")
	flag.Parse()

	if settings[*version] == nil {
		logging.NewLocalLogger().Errorf("unknown version: %v", *version)
		os.Exit(1)
	}

	config := settings[*version]

	// Parse cellosaurus.txt
	cellsTXT := absoluteFilePath("cellosaurus", "cellosaurus.txt")
	cellsCSV := absoluteFilePath("data", "cells.csv")
	cellAttrsCSV := absoluteFilePath("data", "cell_attributes.csv")
	if err := scanCellsTXT(config.startCellsTXT, cellsTXT, cellsCSV, cellAttrsCSV); err != nil {
		logging.NewLocalLogger().Errorf("scan cellosaurus.txt: %v", err)
		os.Exit(1)
	}

	// Parse cellosaurus_refs.txt
	refsTXT := absoluteFilePath("cellosaurus", "cellosaurus_refs.txt")
	refsCSV := absoluteFilePath("data", "refs.csv")
	refAttrsCSV := absoluteFilePath("data", "ref_attributes.csv")
	if err := scanRefsTXT(config.startRefsTXT, refsTXT, refsCSV, refAttrsCSV); err != nil {
		logging.NewLocalLogger().Errorf("scan cellosaurus_refs.txt: %v", err)
		os.Exit(1)
	}

	// Parse cellosaurus_xrefs.txt
	xrefsTXT := absoluteFilePath("cellosaurus", "cellosaurus_xrefs.txt")
	xrefsCSV := absoluteFilePath("data", "xrefs.csv")
	if err := scanXRefsTXT(config.startXRefsTXT, xrefsTXT, xrefsCSV); err != nil {
		logging.NewLocalLogger().Errorf("scan cellosaurus_xrefs.txt: %v", err)
		os.Exit(1)
	}
}

// formattedCSVLine returns formatted string representing a single csv line.
func formattedCSVLine(useLineNum bool, lineNum int, words ...string) string {
	var line string
	for _, w := range words {
		line += fmt.Sprintf("\"%s\",", w)
	}
	if useLineNum {
		return strings.TrimSuffix(fmt.Sprintf("%d,%s", lineNum, line), ",") + "\n"
	}
	return strings.TrimSuffix(line, ",") + "\n"
}

// absoluteFilePath returns the absolute path of a file.
func absoluteFilePath(dir string, file string) string {
	root, err := filepath.Abs(".")
	if err != nil {
		logging.NewLocalLogger().Errorf("absoluteFilePath: %v", err)
		os.Exit(1)
	}
	return fmt.Sprintf("%s/%s/%s", root, dir, file)
}

// fromEnvOrDefaultVal returns an environment variable value if it exists,
// or the specified default value.
func fromEnvOrDefaultVal(env, defaultVal string) string {
	if val := os.Getenv(env); val != "" {
		return val
	}
	return defaultVal
}
