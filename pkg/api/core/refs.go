package core

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/assefamaru/cellosaurus-api/pkg/db"
)

type Reference struct {
	Identifier []string `json:"identifier"`
	Authors    []string `json:"authors,omitempty"`
	Consortium []string `json:"group/consortium,omitempty"`
	Title      string   `json:"title"`
	Citation   string   `json:"citation"`
}

// ListReferences returns a list of paginated References.
func ListReferences(from, to int) ([]*Reference, error) {
	mysql, err := db.NewMySQLFromEnv()
	if err != nil {
		return nil, err
	}
	conn, err := mysql.Connect()
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	query := fmt.Sprintf("SELECT identifier, citation FROM refs LIMIT %d,%d;", from, to)
	rows, err := conn.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var references []*Reference

	for rows.Next() {
		var ref Reference
		var identifier string

		if err := rows.Scan(&identifier, &ref.Citation); err != nil {
			return nil, err
		}
		ref.Identifier = strings.Split(strings.TrimSuffix(identifier, ";"), "; ")

		query := fmt.Sprintf("SELECT attribute, content FROM ref_attributes WHERE identifier = '%s';", identifier)
		detailRows, err := conn.Query(query)
		if err != nil {
			return nil, err
		}
		defer detailRows.Close()

		for detailRows.Next() {
			var attr, content string
			if err := detailRows.Scan(&attr, &content); err != nil {
				return nil, err
			}
			switch attr {
			case "RA":
				ref.Authors = append(
					ref.Authors,
					strings.Split(strings.TrimSuffix(content, ";"), ", ")...,
				)
			case "RG":
				ref.Consortium = append(
					ref.Consortium,
					strings.Split(content, ";")...,
				)
			case "RT":
				ref.Title = ref.Title + strings.TrimSuffix(
					strings.TrimPrefix(content, "\""), "\";",
				)
			}
		}

		references = append(references, &ref)
	}

	return references, nil
}

// CountReferences returns the total number of References.
func CountReferences() (int, error) {
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
	query := "SELECT count FROM statistics WHERE attribute = 'references';"
	if err := conn.QueryRow(query).Scan(&count); err != nil {
		return -1, err
	}
	return strconv.Atoi(count)
}
