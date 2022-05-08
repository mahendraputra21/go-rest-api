package query

import (
	"go-rest-api/app/model"
	"go-rest-api/helper"
	"log"
)

func GetOrderDetailById(id int16) (model.Order, error) {

	//connection to db postgres
	db := helper.CreateConnection()

	//we close the connection at the end of the process
	defer db.Close()

	var order model.Order

	sqlStatement := `SELECT 
						o.id,
						o.product_id,
						(select p.product_name from public.product p where p.id=o.product_id) as product_name,
						o.qty,
						o.price::numeric::int,
						o.grand_total
						FROM public."order" o
						where o.id=$1`

	row := db.QueryRow(sqlStatement, id)

	err := row.Scan(&order.ID,
		&order.ProductId,
		&order.ProductName,
		&order.Qty,
		&order.Price,
		&order.GrandTotal)

	if err != nil {
		log.Fatalf("Cannot get data product by id %v", err)
	}

	return order, err
}
