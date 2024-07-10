package model

type Product struct {
	ID    int    `json:"id_product"`
	Name  string `json:"name"`
	Price int    `json:"price"`
}
