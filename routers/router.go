package routers

import (
	"OneGO/controllers"

	"github.com/gorilla/mux"
)

func InitializeRouter() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/product", controllers.GetProducts).Methods("GET")
	router.HandleFunc("/product/{id}", controllers.GetProduct).Methods("GET")
	router.HandleFunc("/product", controllers.CreateProduct).Methods("POST")
	router.HandleFunc("/product/{id}", controllers.UpdateProduct).Methods("PUT")
	router.HandleFunc("/product/{id}", controllers.DeleteProduct).Methods("DELETE")

	router.HandleFunc("/measure", controllers.GetMeasures).Methods("GET")
	router.HandleFunc("/measure/{id}", controllers.GetMeasure).Methods("GET")
	router.HandleFunc("/measure", controllers.CreateMeasure).Methods("POST")
	router.HandleFunc("/measure/{id}", controllers.UpdateMeasure).Methods("PUT")
	router.HandleFunc("/measure/{id}", controllers.DeleteMeasure).Methods("DELETE")

	return router
}
