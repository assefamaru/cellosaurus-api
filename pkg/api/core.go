package api

import (
	"database/sql"
	"fmt"
	"strconv"
	"strings"
)

type Cell struct {
	Identifier      string           `json:"identifier"`
	Accession       Accession        `json:"accession"`
	Synonyms        []string         `json:"synonyms,omitempty"`
	Sex             string           `json:"sex,omitempty"`
	Age             string           `json:"age,omitempty"`
	Category        string           `json:"category"`
	Date            Date             `json:"date"`
	CrossReferences []CrossReference `json:"crossReferences,omitempty"`
	References      []string         `json:"references,omitempty"`
	WebPages        []string         `json:"webPages,omitempty"`
	Comments        []string         `json:"comments,omitempty"`
	STR             STR              `json:"strProfileData,omitempty"`
	Diseases        []Disease        `json:"diseases,omitempty"`
	OriginSpecies   []OriginSpecies  `json:"speciesOfOrigin"`
	Hierarchy       []Hierarchy      `json:"hierarchy,omitempty"`
	SameOriginAs    []SameOriginAs   `json:"sameOriginAs,omitempty"`
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

type OriginSpecies struct {
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

// Metadata for paginated results.
type Meta struct {
	Page     int `json:"page"`
	PerPage  int `json:"perPage"`
	LastPage int `json:"lastPage"`
	Total    int `json:"total"`
}

// Paginated collection of cell lines.
type Cells struct {
	Meta Meta   `json:"meta"`
	Data []Cell `json:"data"`
}

// Returns a list of paginated cell lines.
func (cells *Cells) List() error {
	db, err := Database()
	if err != nil {
		return err
	}
	defer db.Close()

	from := getPaginationFrom(cells.Meta)
	to := cells.Meta.PerPage
	query := fmt.Sprintf(
		"SELECT identifier, accession, secondary, synonyms, sex, age, category, date FROM cells LIMIT %d,%d;",
		from,
		to,
	)

	rows, err := db.Query(query)
	if err != nil {
		logSentry(err)
		return err
	}
	defer rows.Close()

	for rows.Next() {
		var (
			secondary string
			synonym   string
			date      string
			cell      Cell
		)
		if err = rows.Scan(
			&cell.Identifier,
			&cell.Accession.Primary,
			&secondary,
			&synonym,
			&cell.Sex,
			&cell.Age,
			&cell.Category,
			&date,
		); err != nil {
			logSentry(err)
			return err
		}
		if secondary != "" {
			cell.Accession.Secondary = strings.Split(secondary, "; ")
		}
		if synonym != "" {
			cell.Synonyms = strings.Split(synonym, "; ")
		}
		if date != "" {
			parsed := strings.Split(date, "; ")
			cell.Date.Created = strings.Split(parsed[0], ": ")[1]
			cell.Date.Updated = strings.Split(parsed[1], ": ")[1]
			cell.Date.Version = strings.Split(parsed[2], ": ")[1]
		}
		cells.Data = append(cells.Data, cell)
	}
	return nil
}

// Finds a cell line using its attributes.
func (cell *Cell) Find() error {
	db, err := Database()
	if err != nil {
		return err
	}
	defer db.Close()

	var (
		secondary string
		synonym   string
		date      string
	)

	phrase := cell.Identifier
	query := fmt.Sprintf(
		"SELECT identifier, accession, secondary, synonyms, sex, age, category, date FROM cells WHERE accession = '%s';",
		phrase,
	)
	row := db.QueryRow(query)
	err = row.Scan(
		&cell.Identifier,
		&cell.Accession.Primary,
		&secondary,
		&synonym,
		&cell.Sex,
		&cell.Age,
		&cell.Category,
		&date,
	)
	if err != nil {
		if err != sql.ErrNoRows {
			logSentry(err)
			return err
		}

		query = fmt.Sprintf(
			"SELECT identifier, accession, secondary, synonyms, sex, age, category, date FROM cells WHERE identifier = '%s';",
			phrase,
		)
		row := db.QueryRow(query)
		err = row.Scan(
			&cell.Identifier,
			&cell.Accession.Primary,
			&secondary,
			&synonym,
			&cell.Sex,
			&cell.Age,
			&cell.Category,
			&date,
		)
		if err != nil {
			if err != sql.ErrNoRows {
				logSentry(err)
				return err
			}

			query = fmt.Sprintf(
				"SELECT identifier, accession, secondary, synonyms, sex, age, category, date FROM cells WHERE synonyms LIKE '%%%s%%';",
				phrase,
			)
			row := db.QueryRow(query)
			err = row.Scan(
				&cell.Identifier,
				&cell.Accession.Primary,
				&secondary,
				&synonym,
				&cell.Sex,
				&cell.Age,
				&cell.Category,
				&date,
			)
			if err != nil {
				if err != sql.ErrNoRows {
					logSentry(err)
				}
				return err
			}
		}
	}

	if secondary != "" {
		cell.Accession.Secondary = strings.Split(secondary, "; ")
	}
	if synonym != "" {
		cell.Synonyms = strings.Split(synonym, "; ")
	}
	if date != "" {
		parsed := strings.Split(date, "; ")
		cell.Date.Created = strings.Split(parsed[0], ": ")[1]
		cell.Date.Updated = strings.Split(parsed[1], ": ")[1]
		cell.Date.Version = strings.Split(parsed[2], ": ")[1]
	}

	query = fmt.Sprintf("SELECT attribute, content FROM cell_attributes WHERE accession = '%s';", cell.Accession.Primary)
	rows, err := db.Query(query)
	if err != nil {
		logSentry(err)
		return err
	}
	defer rows.Close()

	for rows.Next() {
		var (
			attribute string
			content   string
		)
		if err = rows.Scan(&attribute, &content); err != nil {
			logSentry(err)
			return err
		}

		switch attribute {
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
			cell.OriginSpecies = append(cell.OriginSpecies, OriginSpecies{"NCBI-Taxonomy", s[1], f[1]})
		case "HI":
			hi := strings.Split(content, " ! ")
			cell.Hierarchy = append(cell.Hierarchy, Hierarchy{"Cellosaurus", hi[0], hi[1]})
		case "OI":
			oi := strings.Split(content, " ! ")
			cell.SameOriginAs = append(cell.SameOriginAs, SameOriginAs{"Cellosaurus", oi[0], oi[1]})
		}
	}
	return nil
}

// Returns the total number of cell lines in db.
func countCells() (int, error) {
	db, err := Database()
	if err != nil {
		return -1, err
	}
	defer db.Close()

	var count string
	query := "SELECT content FROM statistics WHERE attribute = 'totalCellLines';"
	err = db.QueryRow(query).Scan(&count)
	if err != nil {
		logSentry(err)
		return -1, err
	}

	total, _ := strconv.Atoi(count)
	return total, nil
}
