package main

import (
	"capiwara-boilerplate/db"
	"capiwara-boilerplate/gql"
	"context"
	"log"
	"net/http"
	"runtime"

	gqlhandler "github.com/graphql-go/graphql-go-handler"
)

func main() {
	if err := db.NewSession("", ""); err != nil {
		return
	}
	runtime.GOMAXPROCS(runtime.NumCPU() - 1)
	handler := gqlhandler.New(&gqlhandler.Config{
		Schema: &gql.Schema,
	})
	http.Handle("/graphql", requireAuth(handler))
	log.Println("Server started at http://localhost:3000/graphql")
	log.Fatal(http.ListenAndServe(":3000", nil))
}

func requireAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("Authorization")
		if token == "" {
			token = "5b085bb0c28dd7370cb2cdcc"
		}
		ctx := context.WithValue(r.Context(), "id", token)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
