package main

import (
	"capiwara-boilerplate/db"
	"capiwara-boilerplate/gql"
	"capiwara-boilerplate/users"
	"context"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"runtime"

	"github.com/gorilla/mux"

	gqlhandler "github.com/graphql-go/graphql-go-handler"
)

func main() {
	router := mux.NewRouter()

	if err := db.NewSession("", ""); err != nil {
		return
	}
	runtime.GOMAXPROCS(runtime.NumCPU() - 1)
	handler := gqlhandler.New(&gqlhandler.Config{
		Schema: &gql.Schema,
	})
	router.HandleFunc("/login", loginAuth)
	router.Handle("/graphql", requireAuth(handler))
	log.Println("Server started at http://localhost:3000/graphql")
	log.Fatal(http.ListenAndServe(":3000", nil))
}

func requireAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("Authorization")
		res, err := users.Decode(token)
		if err != nil {
			//TODO return forbidden
			log.Printf("Permission denied: %v", err)
			return
		}
		if res.Id == "" {
			//TODO return error
			res.Id = "5b085bb0c28dd7370cb2cdcc"
		}
		ctx := context.WithValue(r.Context(), "id", res.Id)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func loginAuth(w http.ResponseWriter, r *http.Request) {
	var login users.Login
	body, err := ioutil.ReadAll(r.Body)

	defer r.Body.Close()
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	err = json.Unmarshal(body, &login)

	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	token, err := login.Auth()
	if err == nil {
		w.Header().Set("Authorization", token)
	}
}
