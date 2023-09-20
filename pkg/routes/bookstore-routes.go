package routes

import (
	"mysql-gorm-gorillamux/pkg/controllers"

	"github.com/gorilla/mux"
)

var RegisterBookStoreRoutes = func(router *mux.Router) {
	router.HandleFunc("/book/", controllers.CreateBook).Methods("POST")
	router.HandleFunc("/book/", controllers.GetBook).Methods("GET")
	router.HandleFunc("/book/{boockId}", controllers.GetBookById).Methods("GET")
	router.HandleFunc("/book/{boockId}", controllers.UpdateBook).Methods("PUT")
	router.HandleFunc("/book/{boockId}", controllers.DeleteBook).Methods("DELETE")
}
