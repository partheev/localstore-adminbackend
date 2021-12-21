package main

import (
	"adminbackend/models"
	"encoding/json"
	"net/http"
)

func (app *application) addProductHandler(w http.ResponseWriter, r *http.Request) {

	var product models.Product
	err := json.NewDecoder(r.Body).Decode(&product)
	if err != nil {
		app.writeError(w, err)
		return
	}
	err = app.DB.DBModel.InsertProduct(&product)
	if err != nil {
		app.writeError(w, err)
		return
	}
	app.writeJson(w, product, 201)
}

func (app *application) getAllProductsHandler(w http.ResponseWriter, r *http.Request) {

	products, err := app.DB.DBModel.GetAllProducts()
	if err != nil {
		app.writeError(w, err)
		return
	}

	app.writeJson(w, products, http.StatusCreated)
}

func (app *application) updateProductHandler(w http.ResponseWriter, r *http.Request) {
	var product models.Product
	err := json.NewDecoder(r.Body).Decode(&product)
	if err != nil {
		app.writeError(w, err)
		return
	}

	err = app.DB.DBModel.UpdateProduct(&product)

	if err != nil {
		app.writeError(w, err)
		return
	}
	res := resMessage("Product successfully updated.")
	app.writeJson(w, res, http.StatusAccepted)
}

func (app *application) deleteOneProductHandler(w http.ResponseWriter, r *http.Request) {
	type reqBody struct {
		Id int `json:"id"`
	}
	var body reqBody
	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		app.writeError(w, err)
		return
	}

	err = app.DB.DBModel.DeleteOneProduct(body.Id)

	if err != nil {
		app.writeError(w, err)
		return
	}
	res := resMessage("Product successfully deleted.")
	app.writeJson(w, res, http.StatusAccepted)

}

func (app *application) deleteProductsHandler(w http.ResponseWriter, r *http.Request) {
	type reqBody struct {
		ProductIds []int `json:"productIds"`
	}
	var body reqBody

	json.NewDecoder(r.Body).Decode(&body)
	err := app.DB.DBModel.DeleteProducts(body.ProductIds)
	if err != nil {
		app.writeError(w, err)
		return
	}

	res := resMessage("Products deleted successfully.")
	app.writeJson(w, res, http.StatusAccepted)
}
