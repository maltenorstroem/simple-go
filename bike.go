package main

import "google.golang.org/genproto/googleapis/type/money"

type Bike struct {
	Id            int         `json:"id"`
	Desc          string      `json:"desc"`
	Frame         string      `json:"frame"`
	Gearing       string      `json:"gearing"`
	CustomerPrice money.Money `json:"customerPrice"`
	SoldOut       bool        `json:"soldOut"`
}

type Bikes []Bike
