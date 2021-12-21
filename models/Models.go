package models

import "database/sql"

type DBModel struct {
	client *sql.DB
}
type DB struct {
	DBModel *DBModel
}

func NewDBModel(client *sql.DB) DB {

	return DB{
		DBModel: &DBModel{
			client: client,
		},
	}
}

type Product struct {
	Id              int      `json:"id"`
	Name            string   `json:"name"`
	Price           float64  `json:"price"`
	Description     string   `json:"description"`
	Offer_percent   int      `json:"offer_percent"`
	Highlights      []string `json:"highlights"`
	Qty_available   int      `json:"qty_available"`
	Delivery_radius float64  `json:"delivery_radius"`
	Image_url       []string `json:"image_url"`
}
