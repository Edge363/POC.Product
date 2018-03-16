package main

//Product is the dynamodb type
type Product struct {
	ID    string  `json:"id"`
	Price float64 `json:"price"`
	Name  string  `json:"name"`
}
