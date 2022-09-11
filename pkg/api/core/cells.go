package core

import (
	"database/sql"
	"fmt"
	"strconv"
	"strings"

	"github.com/assefamaru/cellosaurus-api/pkg/db"
)

type Cell struct {
	Identifier      string            `json:"identifier"`
	Accession       Accession         `json:"accession"`
	Synonyms        []string          `json:"synonyms,omitempty"`
	CrossReferences []CrossReference  `json:"crossReferences,omitempty"`
	References      []string          `json:"references,omitempty"`
	WebPages        []string          `json:"webPages,omitempty"`
	Comments        []string          `json:"comments,omitempty"`
	STR             STR               `json:"strProfileData,omitempty"`
	Diseases        []Disease         `json:"diseases,omitempty"`
	SpeciesOfOrigin []SpeciesOfOrigin `json:"speciesOfOrigin"`
	Hierarchy       []Hierarchy       `json:"hierarchy,omitempty"`
	SameOriginAs    []SameOriginAs    `json:"sameOriginAs,omitempty"`
	Sex             string            `json:"sex,omitempty"`
	Age             string            `json:"age,omitempty"`
	Category        string            `json:"category"`
	Date            Date              `json:"date"`
}

type Accession struct {
	Primary   string   `json:"primary"`
	Secondary []string `json:"secondary,omitempty"`
}

type Date struct {
	Created string `json:"created"`
	Updated string `json:"lastUpdated"`
	Version string `json:"version"`
}

type CrossReference struct {
	Database  string `json:"database,omitempty"`
	Accession string `json:"accession,omitempty"`
}

type Marker struct {
	ID      string `json:"id,omitempty"`
	Alleles string `json:"alleles,omitempty"`
}

type STR struct {
	Sources []string `json:"sources,omitempty"`
	Markers []Marker `json:"markers,omitempty"`
}

type Disease struct {
	Terminology string `json:"terminology,omitempty"`
	Accession   string `json:"accession,omitempty"`
	Disease     string `json:"disease,omitempty"`
}

type SpeciesOfOrigin struct {
	Terminology string `json:"terminology,omitempty"`
	Accession   string `json:"accession,omitempty"`
	Species     string `json:"species,omitempty"`
}

type Hierarchy struct {
	Terminology string `json:"terminology,omitempty"`
	Accession   string `json:"accession,omitempty"`
	DerivedFrom string `json:"derivedFrom,omitempty"`
}

type SameOriginAs struct {
	Terminology string `json:"terminology,omitempty"`
	Accession   string `json:"accession,omitempty"`
	Identifier  string `json:"identifier,omitempty"`
}

// ListCells returns a list of paginated cell lines.
func ListCells(from, to int) ([]*Cell, error) {
	mysql, err := db.NewMySQLFromEnv()
	if err != nil {
		return nil, err
	}
	conn, err := mysql.Connect()
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	query := fmt.Sprintf("SELECT identifier, accession, secondary, synonyms, sex, age, category, date FROM cells LIMIT %d,%d;", from, to)
	rows, err := conn.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var cells []*Cell

	for rows.Next() {
		var cell Cell
		var sec, syn, date string

		if err := rows.Scan(&cell.Identifier, &cell.Accession.Primary, &sec, &syn, &cell.Sex, &cell.Age, &cell.Category, &date); err != nil {
			return nil, err
		}
		if sec != "" {
			cell.Accession.Secondary = strings.Split(sec, "; ")
		}
		if syn != "" {
			cell.Synonyms = strings.Split(syn, "; ")
		}
		if date != "" {
			splitDate := strings.Split(date, "; ")
			cell.Date.Created = strings.Split(splitDate[0], ": ")[1]
			cell.Date.Updated = strings.Split(splitDate[1], ": ")[1]
			cell.Date.Version = strings.Split(splitDate[2], ": ")[1]
		}

		cells = append(cells, &cell)
	}

	return cells, nil
}

// FetchCell retrieves a cell line using its attribute.
// Attribute can be one of: accession, identifier, or synonym.
func FetchCell(attribute string) (*Cell, error) {
	mysql, err := db.NewMySQLFromEnv()
	if err != nil {
		return nil, err
	}
	conn, err := mysql.Connect()
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	var cell Cell
	var sec, syn, date string
	var found bool

	attributeKeys := []string{"accession = '%s';", "identifier = '%s';", "synonyms LIKE '%%%s%%';"}
	for _, key := range attributeKeys {
		query := fmt.Sprintf("SELECT identifier, accession, secondary, synonyms, sex, age, category, date FROM cells WHERE "+key, attribute)
		row := conn.QueryRow(query)
		if err := row.Scan(&cell.Identifier, &cell.Accession.Primary, &sec, &syn, &cell.Sex, &cell.Age, &cell.Category, &date); err != nil {
			if err != sql.ErrNoRows {
				return nil, err
			}
		} else {
			found = true
			break
		}
	}

	if !found {
		return nil, fmt.Errorf("cell line not found: %s", attribute)
	}
	if sec != "" {
		cell.Accession.Secondary = strings.Split(sec, "; ")
	}
	if syn != "" {
		cell.Synonyms = strings.Split(syn, "; ")
	}
	if date != "" {
		splitDate := strings.Split(date, "; ")
		cell.Date.Created = strings.Split(splitDate[0], ": ")[1]
		cell.Date.Updated = strings.Split(splitDate[1], ": ")[1]
		cell.Date.Version = strings.Split(splitDate[2], ": ")[1]
	}

	query := fmt.Sprintf("SELECT attribute, content FROM cell_attributes WHERE accession = '%s';", cell.Accession.Primary)
	rows, err := conn.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var attr, content string
		if err := rows.Scan(&attr, &content); err != nil {
			return nil, err
		}
		switch attr {
		case "DR":
			dr := strings.Split(content, "; ")
			cell.CrossReferences = append(cell.CrossReferences, CrossReference{dr[0], dr[1]})
		case "RX":
			cell.References = append(cell.References, strings.TrimRight(content, ";"))
		case "WW":
			cell.WebPages = append(cell.WebPages, content)
		case "CC":
			cell.Comments = append(cell.Comments, content)
		case "ST":
			str := strings.Split(content, ": ")
			if str[0] == "Source(s)" {
				cell.STR.Sources = strings.Split(str[1], "; ")
			} else {
				cell.STR.Markers = append(cell.STR.Markers, Marker{str[0], str[1]})
			}
		case "DI":
			di := strings.Split(content, "; ")
			cell.Diseases = append(cell.Diseases, Disease{di[0], di[1], di[2]})
		case "OX":
			f := strings.Split(content, " ! ")
			s := strings.Split(strings.TrimRight(f[0], ";"), "=")
			cell.SpeciesOfOrigin = append(cell.SpeciesOfOrigin, SpeciesOfOrigin{"NCBI-Taxonomy", s[1], f[1]})
		case "HI":
			hi := strings.Split(content, " ! ")
			cell.Hierarchy = append(cell.Hierarchy, Hierarchy{"Cellosaurus", hi[0], hi[1]})
		case "OI":
			oi := strings.Split(content, " ! ")
			cell.SameOriginAs = append(cell.SameOriginAs, SameOriginAs{"Cellosaurus", oi[0], oi[1]})
		}
	}

	return &cell, nil
}

// CountCells returns the total number of cell lines.
func CountCells() (int, error) {
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
	query := "SELECT COUNT(*) FROM cells;"
	if err := conn.QueryRow(query).Scan(&count); err != nil {
		return -1, err
	}
	return strconv.Atoi(count)
}
