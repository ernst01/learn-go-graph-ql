package songlist

import (
	"github.com/graphql-go/graphql"
)

var (
	// SongListSchema represents the GraphQL schema for songlist
	SongListSchema graphql.Schema
)

// InitGraphQL initializes graphql
func (s *Server) InitGraphQL() {
	queryType := graphql.NewObject(graphql.ObjectConfig{
		Name: "Query",
		Fields: graphql.Fields{
			"name": &graphql.Field{
				Type:        graphql.String,
				Description: "The name of the songlist.",
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					return "SoNgLiSt", nil
				},
			},
		},
	})
	SongListSchema, _ = graphql.NewSchema(graphql.SchemaConfig{
		Query: queryType,
	})
}
