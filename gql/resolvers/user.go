package resolvers

import (
	"capiwara-boilerplate/users"

	"github.com/graphql-go/graphql"
)

func GetUser(params graphql.ResolveParams) (interface{}, error) {
	id := params.Context.Value("id").(string)
	user, err := users.GetUser(id)
	if err != nil {
		return nil, err
	}
	return user, nil
}
