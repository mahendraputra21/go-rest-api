package controller

import (
	"encoding/json"
	"go-rest-api/app/command"
	"go-rest-api/app/model"
	"go-rest-api/helper"
	"log"
	"net/http"
)

func POSTBrandControl(w http.ResponseWriter, r *http.Request) {

	helper.EnableCors(&w)

	//create an empty brand of type model.Brand
	var brand model.Brand

	//decode data json request to brand
	err := json.NewDecoder(r.Body).Decode(&brand)

	if err != nil {
		log.Fatalf("Cannot decode from request body. %v", err)
	}

	//validation for brand name
	if helper.IsEmpty(brand.BrandName) {
		res := helper.BrandResponse{
			Status:  http.StatusBadRequest,
			Message: "Brand Name cannot be empty. Pass a valid string value and try again !",
		}
		json.NewEncoder(w).Encode(res)
		return
	}

	//call AddNewBrand function then insert the new brand
	insertID := command.AddNewBrand(brand)

	//format the response object
	res := helper.Response{
		ID:      insertID,
		Message: "New Brand is Added",
	}

	//send the response
	json.NewEncoder(w).Encode(res)
}
