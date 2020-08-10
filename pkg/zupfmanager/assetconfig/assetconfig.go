package assetconfig

// Config is config
type Config struct {
	Produce  []int              `json:"produce"`
	Extracts map[string]Extract `json:"extract"`
}

// Extract is extract
type Extract struct {
	Notes map[string]Note `json:"notes"`
}

// Note is a note
type Note struct {
	Text string `json:"text"`
}
