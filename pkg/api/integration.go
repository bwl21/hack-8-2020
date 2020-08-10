package api

import (
	"io"
	"regexp"
)

// AssetStore provides access to assets
type AssetStore interface {
	// List lists all assets whoose filenames match the search.
	// If search is nil all assets in the store match.
	List(serch *regexp.Regexp) ([]*Asset, error)

	// Get provides access to a single asset identified by its filename
	Get(filename string) (*Asset, error)

	// Read provides acess to the assets content
	Read(filename string) (io.ReadCloser, error)
}
