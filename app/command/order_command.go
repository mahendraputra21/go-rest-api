package command

import (
	"fmt"
	"go-rest-api/app/model"
	"go-rest-api/helper"
	"log"
)

func AddNewOrder(order model.Order) int32 {

	//connection to db postgres
	db := helper.CreateConnection()

	//we close connection at the end of the process
	defer db.Close()

	//create an insert query
	sqlStatement := `INSERT INTO "order" (product_id, price, qty, grand_total) VALUES ($1, $2, $3, $4) RETURNING id`

	var id int32

	//formula grand total
	order.GrandTotal = order.Price * int32(order.Qty)

	err := db.QueryRow(sqlStatement, order.ProductId, order.Price, order.Qty, order.GrandTotal).Scan(&id)

	if err != nil {
		log.Fatalf("Cannot execute this query. %v", err)
	}

	fmt.Printf("Insert data order single record %v", id)

	//return
	return id
}
