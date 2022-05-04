package controller

import (
	"encoding/json"
	"go-rest-api/command_query"
	"go-rest-api/helper"
	"go-rest-api/model"
	"log"
	"net/http"
	"strconv"
)

func POSTOrderControl(w http.ResponseWriter, r *http.Request) {

	helper.EnableCors(&w)

	//create an empty order of type model.order
	var order model.Order

	//decode data json request to order
	err := json.NewDecoder(r.Body).Decode(&order)

	if err != nil {
		log.Fatalf("Cannot decode from request body. %v", err)
	}

	//call AddNewOrder function then insert the new order
	insertID := command_query.AddNewOrder(order)

	//format the response object
	res := helper.Response{
		ID:      insertID,
		Message: "New Order is Added",
	}

	//send the response
	json.NewEncoder(w).Encode(res)
}

func GetOrderDetailByIdControl(w http.ResponseWriter, r *http.Request) {

	helper.EnableCors(&w)

	query := r.URL.Query()
	param := query.Get("id")

	if param == "" {
		log.Fatalf("Param is empty")
	}

	//convert param id from string to int
	id, err := strconv.Atoi(param)

	//call product query
	order, err := command_query.GetOrderDetailById(int16(id))

	if err != nil {
		log.Fatalf("Cannot get data Order detail %v", err)
	}

	//format the response object
	res := helper.OrderResponse{
		Status:  http.StatusOK,
		Message: "Fetch data Order detail is success",
		Data:    order,
	}

	//return json response
	json.NewEncoder(w).Encode(res)
}
