package route

import (
	"github.com/gorilla/mux"
	"github.com/hireza/sirclo-test/berat/app/delivery/rest/handler"
	"github.com/hireza/sirclo-test/berat/packages"
)

func NewWeightRoute(route *mux.Router, mgr packages.Packages) {
	handler := handler.NewWeightHandler(mgr)

	route.Handle("/weights", handler.GetAll()).Methods("GET")
	route.Handle("/weight", handler.GetByDate()).Methods("GET")
	route.Handle("/weight", handler.Create()).Methods("POST")
	route.Handle("/weight", handler.Update()).Methods("PUT")
	route.Handle("/weight", handler.Delete()).Methods("DELETE")
}
