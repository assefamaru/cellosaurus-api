package data

// ListStatistics returns a list of Cellosaurus data stats.
func (s *ObjectStore) ListStatistics(version string) (*Statistics, error) {
	cells, err := s.CountCells()
	if err != nil {
		return nil, err
	}
	refs, err := s.CountReferences()
	if err != nil {
		return nil, err
	}
	xrefs, err := s.CountCrossReferences()
	if err != nil {
		return nil, err
	}
	stats := &Statistics{
		Version: version,
		Cells:   cells,
		Refs:    refs,
		XRefs:   xrefs,
	}
	return stats, nil
}
