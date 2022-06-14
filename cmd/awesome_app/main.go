package main

import (
	"context"
	"encoding/json"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/gorilla/schema"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/komron-m/projx/config"
	"github.com/komron-m/projx/db/sqlc/queries"
	"github.com/spf13/viper"
	"log"
	"net/http"
)

func main() {
	v, err := initViper()
	if err != nil {
		log.Fatal(err)
	}

	q, err := newQueries(v)
	if err != nil {
		log.Fatal(err)
	}

	r := mux.NewRouter()
	r.Use(func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json")
			next.ServeHTTP(w, r)
		})
	})

	decoder := schema.NewDecoder()
	decoder.SetAliasTag("json")

	r.Handle("/users/create", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		req := new(queries.CreateUserParams)
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		user, err := q.CreateUser(r.Context(), req)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		json.NewEncoder(w).Encode(user)
	})).Methods(http.MethodPost)

	r.Handle("/users/list", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		r.URL.Query()
		req := new(queries.GetUsersParams)
		if err := decoder.Decode(req, r.URL.Query()); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		users, err := q.GetUsers(r.Context(), req)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		json.NewEncoder(w).Encode(users)

	})).Methods(http.MethodGet)

	r.Handle("/users/{id}", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		vars := mux.Vars(r)
		userId := vars["id"]
		id, _ := uuid.Parse(userId)

		switch r.Method {
		case http.MethodGet:
			user, err := q.GetUser(ctx, id)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			json.NewEncoder(w).Encode(user)

		case http.MethodPut:
			req := new(queries.UpdateUserParams)
			if err := json.NewDecoder(r.Body).Decode(req); err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}
			req.ID = id
			user, err := q.UpdateUser(ctx, req)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			json.NewEncoder(w).Encode(user)

		case http.MethodDelete:
			if err := q.DeleteUser(ctx, id); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
		}
	})).Methods(http.MethodGet, http.MethodPut, http.MethodDelete)

	if err := http.ListenAndServe(":4000", r); err != nil {
		log.Fatal(err)
	}
}

func initViper() (*viper.Viper, error) {
	return config.NewViper()
}

func newQueries(v *viper.Viper) (*queries.Queries, error) {
	ctx := context.Background()

	dbConfig, err := config.NewDBConfig(v)
	if err != nil {
		log.Fatal(err)
	}

	pgxConn, err := pgxpool.Connect(ctx, dbConfig.ConnectionString())
	if err != nil {
		log.Fatal(err)
	}

	return queries.New(pgxConn), nil
}
