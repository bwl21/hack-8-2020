package api

import (
	"net/http"

	"github.com/friendsofgo/graphiql"
	"github.com/samsarahq/thunder/graphql"
	"github.com/samsarahq/thunder/graphql/introspection"
	"github.com/samsarahq/thunder/graphql/schemabuilder"
)

// NewServer creates a new server
func NewServer() (*Server, error) {
	var s Server

	err := s.buildSchema()
	if err != nil {
		return nil, err
	}

	return &s, nil
}

// Server serves the GraphQL API
type Server struct {
	Schema *graphql.Schema
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

func (s *Server) registerQuery(builder *schemabuilder.Schema) {
	q := builder.Query()
	q.FieldFunc("assets", func() []*Asset {
		return []*Asset{
			{},
		}
	})

	asset := builder.Object("Asset", Asset{})
	asset.FieldFunc("filename", func() string {
		return "hugo"
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
}
