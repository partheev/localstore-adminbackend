package models

import (
	"context"
	"time"

	"github.com/lib/pq"
)

/*
	Returns list of all products in database
*/
func (m *DBModel) GetAllProducts() ([]Product, error) {
	stmt := `SELECT id,name ,price,description,offer_percent,
	highlights,qty_available,delivery_radius,image_url FROM products`
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*2)
	defer cancel()
	rows, err := m.client.QueryContext(ctx, stmt)
	if err != nil {
		return nil, err
	}

	var products []Product
	defer rows.Close()
	for rows.Next() {
		var p Product
		err := rows.Scan(
			&p.Id,
			&p.Name,
			&p.Price,
			&p.Description,
			&p.Offer_percent,
			pq.Array(&p.Highlights),
			&p.Qty_available,
			&p.Delivery_radius,
			pq.Array(&p.Image_url),
		)
		if err != nil {
			return nil, err
		}
		products = append(products, p)
	}

	return products, nil

}
func (db *DBModel) InsertProduct(product *Product) error {
	stmt := `INSERT INTO
		products (	
			name,price,description,offer_percent,
			highlights,qty_available,delivery_radius,image_url
			) VALUES($1,$2,$3,$4,$5,$6,$7,$8)`

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*2)
	defer cancel()
	_, err := db.client.ExecContext(ctx, stmt,
		product.Name,
		product.Price,
		product.Description,
		product.Offer_percent,
		pq.Array(product.Highlights),
		product.Qty_available,
		product.Delivery_radius,
		pq.Array(product.Image_url),
	)
	if err != nil {
		return err
	}

	return nil
}
func (db *DBModel) UpdateProduct(p *Product) error {
	stmt := `UPDATE products set name = $1,price = $2, description = $3,
			offer_percent = $4, highlights = $5,qty_available = $6, delivery_radius = $7,
			image_url = $8 	WHERE id = $9`
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*2)
	defer cancel()
	_, err := db.client.ExecContext(ctx, stmt,
		p.Name,
		p.Price,
		p.Description,
		p.Offer_percent,
		pq.Array(p.Highlights),
		p.Qty_available,
		p.Delivery_radius,
		pq.Array(p.Image_url),
		p.Id,
	)
	if err != nil {
		return err
	}
	return nil
}

func (db *DBModel) DeleteOneProduct(id int) error {
	stmt := `DELETE FROM products where id = $1`
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*2)
	defer cancel()

	a, err := db.client.ExecContext(ctx, stmt, id)

	if err != nil {
		return err
	}
	rowsEffects, err := a.RowsAffected()
	if rowsEffects != 1 || err != nil {
		return ThrowDbError("Failed to delete product.")
	}

	return nil
}

func (db *DBModel) DeleteProducts(idArray []int) error {
	stmt := `DELETE FROM products WHERE id = ANY($1)`
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*2)
	defer cancel()

	result, err := db.client.ExecContext(ctx, stmt, pq.Array(idArray))
	customError := ThrowDbError("Unable to delete products.Please Try again.")
	if err != nil {
		return customError
	}
	idsLength := int64(len(idArray))
	rowsAffects, err := result.RowsAffected()
	if idsLength != rowsAffects || err != nil {
		//we should revoke the operation before deploying into production
		return customError
	}

	return nil
}
