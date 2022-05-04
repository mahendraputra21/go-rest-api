package unit_test

import (
	"bytes"
	"go-rest-api/controller"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPOSTBrandControl(t *testing.T) {

	var jsonStr = []byte(`{"id":6,"brand_name":"Brand Name from Unit test"}`)

	req, err := http.NewRequest("POST", "/brand", bytes.NewBuffer(jsonStr))

	if err != nil {
		t.Fatal(err)
	}

	req.Header.Set("Content-Type", "application/json")

	rr := httptest.NewRecorder()

	handler := http.HandlerFunc(controller.POSTBrandControl)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler return wrong status code: got %v want %v", status, http.StatusOK)
	}

	var actual = rr.Body.String()
	var expected string = `{"id":6,"message":"New Brand is Added"}`

	require.JSONEq(t, expected, actual)
}

func TestPOSTProductControl(t *testing.T) {

	var jsonStr = []byte(`{"id":5, "product_name":"Product from unit test","brand_id":5}`)

	req, err := http.NewRequest("POST", "/product", bytes.NewBuffer(jsonStr))

	if err != nil {
		t.Fatal(err)
	}

	req.Header.Set("Content-Type", "application/json")

	rr := httptest.NewRecorder()

	handler := http.HandlerFunc(controller.POSTProductControl)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler return wrong status code: got %v want %v", status, http.StatusOK)
	}

	var actual string = rr.Body.String()
	var expected string = `{"id":5,"message":"New Product is Added"}`

	require.JSONEq(t, expected, actual)
}

func TestPOSTOrderControl(t *testing.T) {

	var jsonStr = []byte(`{"id":3, "product_id":2, "price":2000, "qty":4}`)

	req, err := http.NewRequest("POST", "/order", bytes.NewBuffer(jsonStr))

	if err != nil {
		t.Fatal(err)
	}

	req.Header.Set("Content-Type", "application/json")

	rr := httptest.NewRecorder()

	handler := http.HandlerFunc(controller.POSTOrderControl)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler return wrong status code: got %v want %v", status, http.StatusOK)
	}

	var actual string = rr.Body.String()
	var expected string = `{"id":3,"message":"New Order is Added"}`

	require.JSONEq(t, expected, actual)
}

func TestGETProductById(t *testing.T) {

	var param string = "1"
	req, err := http.NewRequest("GET", "/product", nil)

	if err != nil {
		t.Fatal(err)
	}

	q := req.URL.Query()
	q.Add("id", param)
	req.URL.RawQuery = q.Encode()
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(controller.GETProductByIdControl)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	var actual string = rr.Body.String()
	var expected string = `{
		"status": 200,
		"message": "Fetch data product with id success",
		"data": {
		  "id": 1,
		  "product_name": "smart watch Cassio",
		  "brand_id": 1,
		  "brand_name": "Cassio",
		  "price": 1000
		}
	  }`

	require.JSONEq(t, expected, actual)
}

func TestGETProductByBrandId(t *testing.T) {

	var param string = "2"
	req, err := http.NewRequest("GET", "/product/brand", nil)

	if err != nil {
		t.Fatal(err)
	}

	q := req.URL.Query()
	q.Add("id", param)
	req.URL.RawQuery = q.Encode()
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(controller.GETProductByBrandIdControl)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	var actual string = rr.Body.String()
	var expected string = `{
		"status": 200,
		"message": "Fetch data product with brand id success",
		"data": [
		  {
			"id": 2,
			"product_name": "smart watch Fossil",
			"brand_id": 2,
			"brand_name": "Fossil",
			"price": 500
		  }
		]
	  }`

	require.JSONEq(t, expected, actual)
}

func TestGETOrderDetailById(t *testing.T) {

	var param string = "1"
	req, err := http.NewRequest("GET", "/order", nil)

	if err != nil {
		t.Fatal(err)
	}

	q := req.URL.Query()
	q.Add("id", param)
	req.URL.RawQuery = q.Encode()
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(controller.GetOrderDetailByIdControl)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	var actual string = rr.Body.String()
	var expected string = `{
		"status": 200,
		"message": "Fetch data Order detail is success",
		"data": {
		  "id": 1,
		  "product_id": 1,
		  "product_name": "smart watch Cassio",
		  "price": 1000,
		  "qty": 2,
		  "grand_total": 2000
		}
	  }`

	require.JSONEq(t, expected, actual)
}
