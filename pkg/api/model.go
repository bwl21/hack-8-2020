package api

// Asset represents a core asset
type Asset struct {
	ID              string    `graphql:"id"`
	Filename        string    `graphql:"filename"`
	Title           string    `graphql:"title"`
	Genre           string    `graphql:"genre"`
	Extracts        Extracts  `graphql:"extracts"`
	Ratings         Ratings   `graphql:"ratings"`
	Copyright       Copyright `graphql:"copyright"`
	ReferenceCopies []string  `graphql:"referenceCopies"`
	Thumbnail       string    `graphql:"thumbnail"`
}

// Extracts are a collection of voices
type Extracts struct {
	Available   []int `graphql:"available"`
	Preselected []int `graphql:"preselected"`
}

// Ratings rate an asset
type Ratings struct {
	Difficulty          int `graphql:"difficulty"`
	EnsemblePlayability int `graphql:"ensemblePlayability"`
}

// Copyright describes the copyright for a piece of music and lyrics
type Copyright struct {
	Music  string `graphql:"music"`
	Lyrics string `graphql:"lyrics"`
}

// Project is an arrangements of assets
type Project struct {
	Title          string         `graphql:"title"`
	ProductionNode string         `graphql:"productionNode"`
	Assets         []ProjectAsset `graphql:"assets"`
}

// ProjectAsset includes an asset in a project
type ProjectAsset struct {
	Asset            Asset    `graphql:"asset"`
	Rating           int      `graphql:"rating"`
	ReferenceCopies  []string `graphql:"referenceCopies"`
	SelectedExtracts []int    `graphql:"selectedExtracts"`
	SortString       string   `graphql:"sortString"`
}
