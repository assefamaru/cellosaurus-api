package data

import (
	"fmt"
	"strconv"
	"strings"
)

// ListReferences returns a list of paginated references.
func (s *ObjectStore) ListReferences(from, to int) ([]*Reference, error) {
	query := fmt.Sprintf("SELECT identifier, citation FROM refs LIMIT %d,%d;", from, to)
	rows, err := s.client.Conn().Query(query)
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
		detailRows, err := s.client.Conn().Query(query)
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
				ref.Authors = append(ref.Authors, strings.Split(strings.TrimSuffix(content, ";"), ", ")...)
			case "RG":
				ref.Consortium = append(ref.Consortium, strings.Split(content, ";")...)
			case "RT":
				ref.Title = ref.Title + strings.TrimSuffix(strings.TrimPrefix(content, "\""), "\";")
			}
		}
		references = append(references, &ref)
	}
	return references, nil
}

// CountReferences returns the total number of references.
func (s *ObjectStore) CountReferences() (int, error) {
	var count string
	query := "SELECT COUNT(*) FROM refs;"
	if err := s.client.Conn().QueryRow(query).Scan(&count); err != nil {
		return -1, err
	}
	return strconv.Atoi(count)
}
