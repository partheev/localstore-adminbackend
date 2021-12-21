package main

import (
	"adminbackend/config"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (app *application) Routes() http.Handler {
	router := httprouter.New()
	productsManagement(app, router)
	return router
}

func productsManagement(app *application, router *httprouter.Router) {
	router.HandlerFunc(http.MethodPut, config.UPDATEPRODUCTROUTE, app.updateProductHandler)
	router.HandlerFunc(http.MethodPost, config.ADDPRODUCTROUTE, app.addProductHandler)
	router.HandlerFunc(http.MethodGet, config.GETPRODUCTSROUTE, app.getAllProductsHandler)
	router.HandlerFunc(http.MethodDelete, config.DELETEONEPRODUCTROUTE, app.deleteOneProductHandler)
	router.HandlerFunc(http.MethodDelete, config.DELETEPRODUCTSROUTE, app.deleteProductsHandler)
}
