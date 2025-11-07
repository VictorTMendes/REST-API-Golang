package model

type Product struct {
	Id    int     `json:"Id"`
	Name  string  `json:"product_name"`
	Price float64 `json:"price"`
}
