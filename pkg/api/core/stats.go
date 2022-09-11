package core

type Statistics struct {
	Version string `json:"version"`
	Cells   int    `json:"cellLines"`
	Refs    int    `json:"references"`
	XRefs   int    `json:"crossReferences"`
}

// ListStatistics returns a list of Cellosaurus data stats.
func ListStatistics(version string) (*Statistics, error) {
	cells, err := CountCells()
	if err != nil {
		return nil, err
	}
	refs, err := CountReferences()
	if err != nil {
		return nil, err
	}
	xrefs, err := CountCrossReferences()
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
