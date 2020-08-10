package zupfmanager

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"

	"github.com/bwl21/zupfmanager/pkg/api"
	"github.com/bwl21/zupfmanager/pkg/zupfmanager/assetconfig"
	log "github.com/sirupsen/logrus"
)

// LocalAssetStore discovers assets in a filesystem
type LocalAssetStore struct {
	Locations []string
}

// List lists all assets matching the search in the store.
func (l *LocalAssetStore) List(search *regexp.Regexp) ([]*api.Asset, error) {
	idx := make(map[string]*api.Asset)
	for _, loc := range l.Locations {
		fs, err := filepath.Glob(filepath.Join(loc, "*.abc"))
		if err != nil {
			return nil, err
		}

		for _, fn := range fs {
			fnb := filepath.Base(fn)
			if search != nil && !search.MatchString(fnb) {
				continue
			}
			if _, exists := idx[fnb]; exists {
				continue
			}

			asset, err := l.loadAsset(fn)
			if err != nil {
				log.WithError(err).WithField("fn", fn).Warn("cannot load asset")
				continue
			}
			idx[fnb] = asset
		}
	}

	res := make([]*api.Asset, 0, len(idx))
	for _, a := range idx {
		res = append(res, a)
	}
	return res, nil
}

// Get provides access to an asset via its path name
func (l *LocalAssetStore) Get(filename string) (*api.Asset, error) {
	var path string
	for _, p := range l.Locations {
		fn := filepath.Join(p, filename)
		if _, err := os.Stat(fn); err == nil {
			path = fn
			break
		}
	}
	if path == "" {
		return nil, os.ErrNotExist
	}

	return l.loadAsset(path)
}

func (l *LocalAssetStore) Read(filename string) (io.ReadCloser, error) {
	var path string
	for _, p := range l.Locations {
		fn := filepath.Join(p, filename)
		if _, err := os.Stat(fn); err == nil {
			path = fn
			break
		}
	}
	if path == "" {
		return nil, os.ErrNotExist
	}

	return os.Open(path)
}

func (l *LocalAssetStore) loadAsset(fn string) (*api.Asset, error) {
	f, err := os.Open(fn)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	res, err := parseAsset(f)
	if err != nil {
		return nil, err
	}
	res.Path = fn
	res.Filename = filepath.Base(fn)

	return res, nil
}

func parseAsset(fc io.Reader) (*api.Asset, error) {
	var res api.Asset

	lr := bufio.NewReader(fc)
	for line, _, err := lr.ReadLine(); err == nil; line, _, err = lr.ReadLine() {
		if err != nil && err != io.EOF {
			return nil, err
		}
		if len(line) < 2 {
			continue
		}

		line := string(line)
		switch line[0:2] {
		case "X:":
			res.ID = line[2:]
			continue
		case "F:":
			res.Filename = line[2:]
			continue
		case "T:":
			res.Title = line[2:]
			continue
		}
		if strings.HasPrefix(line, "S:Genre ") {
			res.Genre = strings.TrimPrefix(line, "S:Genre ")
			continue
		}
		if strings.HasPrefix(line, "%%%%zupfnoter.config") {
			log.Debug("found json section")
			break
		}
	}

	var cfg assetconfig.Config
	err := json.NewDecoder(lr).Decode(&cfg)
	if err != nil {
		return nil, fmt.Errorf("cannot decode config: %w", err)
	}

	for k := range cfg.Extracts {
		x, err := strconv.Atoi(k)
		if err != nil {
			continue
		}
		res.Extracts.Available = append(res.Extracts.Available, x)
	}
	res.Extracts.Preselected = cfg.Produce
	res.Copyright.Music = cfg.Extracts["0"].Notes["T02_copyright_music"].Text

	return &res, nil
}
