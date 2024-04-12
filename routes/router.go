package routes

import (
	"net/http"

	"github.com/gorilla/mux"

	"goProject/controllers"
)

func SetupRoutes() {
	router := mux.NewRouter()

	router.HandleFunc("/users/login", controllers.Login).Methods("POST")
	router.HandleFunc("/users/logout", controllers.Logout).Methods("POST")

	router.HandleFunc("/users", controllers.CreateUser).Methods("POST")
	router.HandleFunc("/users", controllers.GetUsers).Methods("GET")
	router.HandleFunc("/users/{id}", controllers.GetUser).Methods("GET")
	router.HandleFunc("/users/{id}", controllers.UpdateUser).Methods("PUT")
	router.HandleFunc("/users/{id}", controllers.DeleteUser).Methods("DELETE")

	http.Handle("/", router)

}
