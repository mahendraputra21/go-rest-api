# Build a RESTful JSON API with GO and PostgreSQL
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
* please set the .env file to connect your postgredb ```POSTGRES_URL="dbname=<dbName> user=<user> password=<Password> host=<host> sslmode=disable"```
* set the path to load .env file
  ```
  func CreateConnection() *sql.DB {
	//load .env file
	err := godotenv.Load(os.ExpandEnv("go-rest-api/.env")) //please set path for the .env if you got the error message
  ```
* run ```go install github.com/pressly/goose/v3/cmd/goose@latest```
* run ```go build goose.go```
* run ```cd db```
* run ```goose postgres "user=<user> password=<password> dbname=<dbname> sslmode=disable" up```  
* run this command ```go run main.go```

## Endpoint Routes 

| Endpoint | Path  | Method  | 
| ------- | --- | --- | 
| Create Brand | /brand | POST | 
| Create Product | /product | POST |
| Get Single Product | /product? | GET | 
| Get All Product by Brand | /product/brand? | GET | 
| Make an transaction/order | /order | POST |
| Transaction Detail | /order? | GET | 


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

