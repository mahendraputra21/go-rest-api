package command

import (
	"fmt"
	"go-rest-api/app/model"
	"go-rest-api/helper"
	"log"
)

func AddNewBrand(brand model.Brand) int32 {

	//connection to db postgres
	db := helper.CreateConnection()

	//we close the connection at the end of the process
	defer db.Close()

	//create an insert query
	sqlStatement := `INSERT INTO brand (brand_name) VALUES ($1) RETURNING id`

	var id int32

	err := db.QueryRow(sqlStatement, brand.BrandName).Scan(&id)

	if err != nil {
		log.Fatalf("Cannot execute this query. %v", err)
	}

	fmt.Printf("Insert data brand single record %v", id)

	//return
	return id
}
