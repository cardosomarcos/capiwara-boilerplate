package gql

import (
	"github.com/graphql-go/graphql"
)

var query = graphql.NewObject(
	graphql.ObjectConfig{
		Name:        "query",
		Description: "query",
		Fields: graphql.Fields{
			"user": userQuery,
		},
	},
)
