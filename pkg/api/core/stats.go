package core

import "github.com/assefamaru/cellosaurus-api/pkg/db"

type Statistics struct {
	Version  string `json:"version"`
	Cells    string `json:"cellLinesTotal"`
	Human    string `json:"cellLinesHuman"`
	Mouse    string `json:"cellLinesMouse"`
	Rat      string `json:"cellLinesRat"`
	Species  string `json:"species"`
	Synonyms string `json:"synonyms"`
	XRefs    string `json:"crossReferences"`
	Refs     string `json:"references"`
	Pubs     string `json:"distinctPublications"`
	Web      string `json:"webLinks"`
	STR      string `json:"cellLinesWithStrProfiles"`
}

// FetchStatistics returns the Cellosaurus data stats.
func FetchStatistics() (*Statistics, error) {
	mysql, err := db.NewMySQLFromEnv()
	if err != nil {
		return nil, err
	}
	conn, err := mysql.Connect()
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	query := "SELECT attribute, count FROM statistics;"
	rows, err := conn.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var stats *Statistics

	for rows.Next() {
		var attr, content string
		if err := rows.Scan(&attr, &content); err != nil {
			return nil, err
		}
		switch attr {
		case "version":
			stats.Version = content
		case "cellLinesTotal":
			stats.Cells = content
		case "cellLinesHuman":
			stats.Human = content
		case "cellLinesMouse":
			stats.Mouse = content
		case "cellLinesRat":
			stats.Rat = content
		case "species":
			stats.Species = content
		case "synonyms":
			stats.Synonyms = content
		case "crossReferences":
			stats.XRefs = content
		case "references":
			stats.Refs = content
		case "distinctPublications":
			stats.Pubs = content
		case "webLinks":
			stats.Web = content
		case "cellLinesWithStrProfiles":
			stats.STR = content
		}
	}

	return stats, nil
}
