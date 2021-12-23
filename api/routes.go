package main

import (
	"adminbackend/config"
	"net/http"

	"github.com/gorilla/mux"
)

func (app *application) Routes() http.Handler {
	router := mux.NewRouter()
	productsManagement(app, router)
	return router
}

func productsManagement(app *application, router *mux.Router) {
	router.HandleFunc(config.UPDATEPRODUCTROUTE, app.updateProductHandler).Methods(http.MethodPut)
	router.HandleFunc(config.ADDPRODUCTROUTE, app.addProductHandler).Methods(http.MethodPost)
	router.HandleFunc(config.GETPRODUCTSROUTE, app.getAllProductsHandler).Methods(http.MethodGet)
	router.HandleFunc(config.DELETEONEPRODUCTROUTE, app.deleteOneProductHandler).Methods(http.MethodDelete)
	router.HandleFunc(config.DELETEPRODUCTSROUTE, app.deleteProductsHandler).Methods(http.MethodDelete)
}
