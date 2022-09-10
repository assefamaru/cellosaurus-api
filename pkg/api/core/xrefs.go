package core

import (
	"strconv"

	"github.com/assefamaru/cellosaurus-api/pkg/db"
)

type XRef struct {
	Abbreviation string `json:"abbreviation"`
	Name         string `json:"name"`
	Server       string `json:"server"`
	URL          string `json:"dbURL"`
	Terminology  string `json:"terminology"`
	Category     string `json:"category"`
}

// ListXRefs returns a list of cross-references.
func ListCrossReferences() ([]*XRef, error) {
	mysql, err := db.NewMySQLFromEnv()
	if err != nil {
		return nil, err
	}
	conn, err := mysql.Connect()
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	query := "SELECT abbrev, name, server, url, term, cat FROM xrefs;"
	rows, err := conn.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var xrefs []*XRef

	for rows.Next() {
		var xref XRef
		if err := rows.Scan(&xref.Abbreviation, &xref.Name, &xref.Server, &xref.URL, &xref.Terminology, &xref.Category); err != nil {
			return nil, err
		}
		xrefs = append(xrefs, &xref)
	}

	return xrefs, nil
}

// CountCrossReferences returns the total number of cross-references.
func CountCrossReferences() (int, error) {
	mysql, err := db.NewMySQLFromEnv()
	if err != nil {
		return -1, err
	}
	conn, err := mysql.Connect()
	if err != nil {
		return -1, err
	}
	defer conn.Close()

	var count string
	query := "SELECT COUNT(*) FROM xrefs;"
	if err := conn.QueryRow(query).Scan(&count); err != nil {
		return -1, err
	}
	return strconv.Atoi(count)
}
