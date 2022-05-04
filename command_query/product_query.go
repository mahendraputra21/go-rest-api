package command_query

import (
	"go-rest-api/model"
	"log"

	"go-rest-api/helper"
)

func GetProductById(id int16) (model.Product, error) {

	//connection to db postgres
	db := helper.CreateConnection()

	//we close the connection at the end of the process
	defer db.Close()

	var product model.Product

	//create an sql statement where id=?
	sqlStatement := `SELECT 
					 p.id, 
					 p.product_name, 
					 p.brand_id, 
					 b.brand_name, 
					 p.price::numeric::int
					 FROM public.product p 
					 INNER JOIN public.brand b 
					 ON p.brand_id = b.id 
					 WHERE p.id=$1;`

	//execute sql query
	row := db.QueryRow(sqlStatement, id)

	err := row.Scan(&product.ID,
		&product.ProductName,
		&product.BrandId,
		&product.BrandName,
		&product.Price)

	if err != nil {
		log.Fatalf("Cannot get data product by id %v", err)
	}

	return product, err
}

func GetProductByBrandId(id int16) ([]model.Product, error) {

	//connection to db postgres
	db := helper.CreateConnection()

	//we close the connection at the end of the process
	defer db.Close()

	var products []model.Product

	//create an sql statement where id=?
	sqlStatement := `SELECT 
					 p.id, 
					 p.product_name, 
					 p.brand_id, 
					 b.brand_name, 
					 p.price::numeric::int
					 FROM public.product p 
					 INNER JOIN public.brand b 
					 ON p.brand_id = b.id 
					 WHERE p.brand_id=$1;`

	//execute sql query
	rows, err := db.Query(sqlStatement, id)

	if err != nil {
		log.Fatalf("Error when execute the query %v", err)
	}

	//we close the rows at the end of the process
	defer rows.Close()

	for rows.Next() {

		var prd model.Product

		err := rows.Scan(&prd.ID, &prd.ProductName, &prd.BrandId, &prd.BrandName, &prd.Price)

		if err != nil {
			log.Fatalf("Cannot get data product by brand id %v", err)
		}

		products = append(products, prd)
	}

	if err = rows.Err(); err != nil {
		return products, err
	}

	return products, err

}
