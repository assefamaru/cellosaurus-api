package data

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

type Reference struct {
	Identifier []string `json:"identifier"`
	Authors    []string `json:"authors,omitempty"`
	Consortium []string `json:"group/consortium,omitempty"`
	Title      string   `json:"title"`
	Citation   string   `json:"citation"`
}

type Statistics struct {
	Version string `json:"version"`
	Cells   int    `json:"cellLines"`
	Refs    int    `json:"references"`
	XRefs   int    `json:"crossReferences"`
}

type XRef struct {
	Abbreviation string `json:"abbreviation"`
	Name         string `json:"name"`
	Server       string `json:"server"`
	URL          string `json:"dbURL"`
	Terminology  string `json:"terminology"`
	Category     string `json:"category"`
}
