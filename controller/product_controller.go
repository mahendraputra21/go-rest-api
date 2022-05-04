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

func POSTProductController(w http.ResponseWriter, r *http.Request) {

	helper.EnableCors(&w)

	//create an empty product of type model.Product
	var product model.Product

	//decode data json request to product
	err := json.NewDecoder(r.Body).Decode(&product)

	if err != nil {
		log.Fatalf("Cannot decode from request body guys. %v", err)
	}

	//validation for price
	if product.Price == 0 {
		res := helper.BrandResponse{
			Status:  http.StatusBadRequest,
			Message: "Price cannot be empty. Pass a valid number value and try again !",
		}
		json.NewEncoder(w).Encode(res)
		return
	}

	//call AddNewProduct function then insert the new Product
	insertID := command_query.AddNewProduct(product)

	//format the response object
	res := helper.Response{
		ID:      insertID,
		Message: "New Product is Added",
	}

	//send the response
	json.NewEncoder(w).Encode(res)
}

func GETProductByIdController(w http.ResponseWriter, r *http.Request) {

	helper.EnableCors(&w)

	query := r.URL.Query()
	param := query.Get("id")

	if param == "" {
		log.Fatalf("Param is empty")
	}

	//convert param id from string to int
	id, err := strconv.Atoi(param)

	//call product query
	product, err := command_query.GetProductById(int16(id))

	if err != nil {
		log.Fatalf("Cannot get data product %v", err)
	}

	//format the response object
	res := helper.ProductResponse{
		Status:  http.StatusOK,
		Message: "Fetch data product with id success",
		Data:    product,
	}

	//return json response
	json.NewEncoder(w).Encode(res)
}

func GETProductByBrandId(w http.ResponseWriter, r *http.Request) {

	helper.EnableCors(&w)

	query := r.URL.Query()
	param := query.Get("id")

	if param == "" {
		log.Fatalf("Param is empty")
	}

	//convert param id from string to int
	id, err := strconv.Atoi(param)

	//call product query
	product, err := command_query.GetProductByBrandId(int16(id))

	if err != nil {
		log.Fatalf("Cannot get data product by Brand id %v", err)
	}

	//format the response object
	res := helper.ProductResponses{
		Status:  http.StatusOK,
		Message: "Fetch data product with brand id success",
		Data:    product,
	}

	//return json response
	json.NewEncoder(w).Encode(res)

}
