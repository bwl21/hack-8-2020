package api

import (
	"io"
	"net/http"
	"strings"

	"github.com/friendsofgo/graphiql"
	"github.com/samsarahq/thunder/graphql"
	"github.com/samsarahq/thunder/graphql/introspection"
	"github.com/samsarahq/thunder/graphql/schemabuilder"
	log "github.com/sirupsen/logrus"
)

const (
	// pathServeAsset is the path where one can retrieve an asset from
	pathServeAsset = "/serve/asset/"
)

// NewServer creates a new server
func NewServer(baseURL string, assets AssetStore) (*Server, error) {
	s := Server{
		BaseURL: baseURL,
		assets:  assets,
	}

	err := s.buildSchema()
	if err != nil {
		return nil, err
	}

	return &s, nil
}

// Server serves the GraphQL API
type Server struct {
	Schema  *graphql.Schema
	BaseURL string

	assets AssetStore
}

func (s *Server) buildSchema() error {
	builder := schemabuilder.NewSchema()
	s.registerQuery(builder)
	s.registerMutation(builder)
	schema, err := builder.Build()
	if err != nil {
		return err
	}
	introspection.AddIntrospectionToSchema(schema)

	s.Schema = schema
	return nil
}

// SchemaJSON returns the GraphQL schema as JSON
func (s *Server) SchemaJSON() ([]byte, error) {
	builder := schemabuilder.NewSchema()
	s.registerQuery(builder)
	s.registerMutation(builder)

	return introspection.ComputeSchemaJSON(*builder)
}

type assetQueryParams struct {
	Filename string `graphql:"filename"`
}

func (s *Server) registerQuery(builder *schemabuilder.Schema) {
	q := builder.Query()
	q.FieldFunc("assets", func() []*Asset {
		res, err := s.assets.List(nil)
		if err != nil {
			log.WithError(err).Warn("no idea how to handle errors here")
		}
		return res
	})
	q.FieldFunc("projects", func() []*Project {
		return []*Project{
			{
				Assets: []ProjectAsset{
					{
						Asset: Asset{},
					},
				},
			},
		}
	})
	q.FieldFunc("asset", func(args assetQueryParams) *Asset {
		res, err := s.assets.Get(args.Filename)
		if err != nil {
			log.WithError(err).Warn("no idea how to handle errors here")
		}
		return res
	})

	assetObj := builder.Object("Asset", Asset{})
	assetObj.FieldFunc("media", func(asset *Asset) Media {
		return Media{
			Self: s.BaseURL + pathServeAsset + asset.Filename,
		}
	})
}

func (s *Server) registerMutation(builder *schemabuilder.Schema) {
	builder.Mutation()
}

// Serve starts the actual handler
func (s *Server) Serve(mux *http.ServeMux) {
	mux.Handle("/graphql", graphql.HTTPHandler(s.Schema))
	gqlhandler, err := graphiql.NewGraphiqlHandler("/graphql")
	if err != nil {
		// This should never happen
		panic(err)
	}
	mux.Handle("/debug", gqlhandler)
	mux.HandleFunc(pathServeAsset, func(res http.ResponseWriter, req *http.Request) {
		fn := strings.TrimPrefix(req.URL.Path, pathServeAsset)
		fc, err := s.assets.Read(fn)
		if err != nil {
			http.Error(res, err.Error(), http.StatusNotFound)
			return
		}
		defer fc.Close()

		res.Header().Add("Content-Type", "application/text")
		_, err = io.Copy(res, fc)
		if err != nil {
			log.WithError(err).WithField("fn", fn).Warn("cannot serve file")
		}
	})
}
