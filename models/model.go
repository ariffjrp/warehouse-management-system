package models

import (
	"gorm.io/gorm"
)

type Inbound struct {
	gorm.Model
	Id        uint64 `json:"id"`
	Products  string `json:"products"`
	Date      string `json:"date"`
	Qty       string `json:"qty"`
	Location  string `json:"location"`
	Building  string `json:"building"`
	Room      string `json:"room"`
	Floor     string `json:"floor"`
	Area      string `json:"area"`
	Rack      string `json:"rack"`
	Racklevel string `json:"racklevel"`
}

type ErrorResponse struct {
	Error string `json:"error"`
}

type SuccessResponse struct {
	Message string `json:"message"`
}
