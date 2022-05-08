package command

import (
	"fmt"
	"go-rest-api/app/model"
	"go-rest-api/helper"
	"log"
)

func AddNewProduct(product model.Product) int32 {

	//connection to db postgres
	db := helper.CreateConnection()

	//we close the connection at the end of the process
	defer db.Close()

	//create an insert query
	sqlStatement := `INSERT INTO product (product_name, brand_id, price) VALUES ($1, $2, $3) RETURNING id`

	var id int32

	err := db.QueryRow(sqlStatement, product.ProductName, product.BrandId, product.Price).Scan(&id)

	if err != nil {
		log.Fatalf("Cannot execute this query. %v", err)
	}

	fmt.Printf("Insert data product single record %v", id)

	return id
}
