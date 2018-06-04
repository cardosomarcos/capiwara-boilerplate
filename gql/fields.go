package gql

import (
	"capiwara-boilerplate/gql/resolvers"

	"github.com/graphql-go/graphql"
)

var userQuery = &graphql.Field{
	Type:        user,
	Description: "get user",
	Resolve:     resolvers.GetUser,
}
