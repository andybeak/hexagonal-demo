package handlers

import "github.com/gorilla/mux"

func ProvideRouter(
	userHandler *UserHttpHandler,
) *mux.Router {
	r := mux.NewRouter()
	router := r.PathPrefix("/v1").Subrouter()
	router.HandleFunc("/users", userHandler.createUser).Methods("POST")
	router.HandleFunc("/users/{id}", userHandler.getUserById).Methods("GET")
	return router
}
