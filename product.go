package main

//Product is the dynamodb type
type Product struct {
	Id    string  `json:"id"`
	Price float64 `json:"price"`
	Name  string  `json:"name"`
}
