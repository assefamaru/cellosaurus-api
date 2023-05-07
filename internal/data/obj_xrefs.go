package data

import "strconv"

// ListXRefs returns a list of cross-references.
func (s *ObjectStore) ListCrossReferences() ([]*XRef, error) {
	query := "SELECT abbrev, name, server, url, term, cat FROM xrefs;"
	rows, err := s.client.Conn().Query(query)
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
func (s *ObjectStore) CountCrossReferences() (int, error) {
	var count string
	query := "SELECT COUNT(*) FROM xrefs;"
	if err := s.client.Conn().QueryRow(query).Scan(&count); err != nil {
		return -1, err
	}
	return strconv.Atoi(count)
}
