package helper

import (
	"database/sql"
	"encoding/json"
	"go-rest-api/model"
	"net/http"
)

func EnableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
}

type NullString struct {
	sql.NullString
}

func (s NullString) MarshallJSON() ([]byte, error) {
	if !s.Valid {
		return []byte("null"), nil
	}
	return json.Marshal(s.String)
}

func (s *NullString) unMarshalJSON(data []byte) error {
	if string(data) == "null" {
		s.String, s.Valid = "", false
		return nil
	}
	s.String, s.Valid = string(data), true
	return nil
}

//isEmpty validates if a string is empty or not.
func IsEmpty(str string) bool {
	if str == "" || len(str) == 0 {
		return true
	} else {
		return false
	}
}

type Response struct {
	ID      int32  `json:"id,omitempty"`
	Message string `json:"message,omitempty"`
}

type BrandResponse struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Data    model.Brand `json:"data"`
}

type ProductResponse struct {
	Status  int           `json:"status"`
	Message string        `json:"message"`
	Data    model.Product `json:"data"`
}

type ProductResponses struct {
	Status  int             `json:"status"`
	Message string          `json:"message"`
	Data    []model.Product `json:"data"`
}

type OrderResponse struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Data    model.Order `json:"data"`
}
