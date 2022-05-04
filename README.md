# Build a RESTful JSON API with GOlang
In this repository, we are creating an sample rest API that will allow customer to :
* create and view brand and product 
* make transaction order 
* get detail order

## Prerequisites
* Have REST Client extension
* Have Postgres DB installed. If not, check it out [here](https://www.postgresql.org/)
* Have Go installed. If not, check it out [here](https://golang.org/doc/install)
* Also after installing, make sure you are working inside your GOPATH

## How to run 
* clone this repository first
* run ```go install github.com/pressly/goose/v3/cmd/goose@latest```
* run ```go build goose.go```
* run ```cd db```
* run ```goose postgres "user=<user> password=<password> dbname=<dbname> sslmode=disable" up```  
* run this command ```go run main.go```

## Endpoint Routes 

| Endpoint | Path  | Method  | Description
| ------- | --- | --- | ----------- |
| Create Brand | /brand | POST | Create brand endpoint |
| Create Product | /product | POST | Create product endpoint |
| Get Single Product | /product? | GET | Retrieve product with querystring, you can use product id, eg:/product?id=123|
| Get All Product by Brand | /product/brand? | GET | Retrieve all products by brand with querystring, you can use brand id, eg:/product/brand?id=123|
| Make an transaction/order | /order | POST | Make a new purchase with how many products that customer buys. Eg: Grand total for Buy 2 products with price @Rp. 500,00 Rp. 1,000,00|
| Transaction Detail | /order? | GET | Retrieve transaction detail with querystring, you can use order/?id=123|


## Results in test.http (using extension REST Client)
### Test Create brand endpoint
  #### Request
  ```
   POST http://localhost:8080/brand
   Content-Type: application/json
   {
    "brand_name": "Samsung"
   }
  ```
   #### Response
   ```
   HTTP/1.1 200 OK
  Access-Control-Allow-Origin: *
  Date: Wed, 04 May 2022 11:11:57 GMT
  Content-Length: 40
  Content-Type: text/plain; charset=utf-8
  Connection: close

  {
  "id": 7,
  "message": "New Brand is Added"
  }
   ```

## References
* https://github.com/pressly/goose
* https://golang.org/pkg/net/http/#NewServeMux
* https://pkg.go.dev/encoding/json
* https://pkg.go.dev/io/ioutil
* https://pkg.go.dev/net/http
* https://pkg.go.dev/log
* https://pkg.go.dev/strings
* https://pkg.go.dev/regexp
* https://pkg.go.dev/context
* https://pkg.go.dev/database/sql

