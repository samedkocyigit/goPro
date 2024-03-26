package routes

import (
	"net/http"

	"github.com/gorilla/mux"

	"goPro/controllers"
)

func setupRoutes(){
	router := mux.NewRouter()

	router.HandleFunc("/users",controllers.createUser).Methods("POST")
	router.HandleFunc("/users",controllers.getUsers).Methods("GET")
	router.HandleFunc("/users/{id}",controllers.getUser).Methods("GET")
	router.HandleFunc("/users/{id}",controllers.updateUser).Methods("PUT")
	router.HandleFunc("/users/{id}",controllers.deleteUser).Methods("DELETE")

	http.ListenAndServe(":8080",router)

}