package routes

import (
	"github.com/Aleksander-Sienkiewicz/go-bookstore-proj/pkg/controllers"
	"github.com/gorilla/mux"
)

var RegisterBookStoreRoutes = func(router *mux.Router) {
	router.HandleFunc("/book/", controllers.CreateBook).Methods("POST")
	router.HandleFunc("/book/", controllers.GetBook).Methods("GET")
	router.HandleFunc("/book/{bookId}", controllers.GetBookById).Methods("GET")
	router.HandleFunc("/Book/{BookId}", controllers.UpdateBook).Methods("PUT")
	router.HandleFunc("/Book/{BookId}", controllers.DeleteBook).Methods("DELETE")
}
