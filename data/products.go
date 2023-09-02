package data

import (
	"encoding/json"
	"fmt"
	"io"
	"time"
)

type Product struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float32 `json:"price"`
	SKU         string  `json:"sku"`
	CreatedOn   string  `json:"-"`
	UpdatedOn   string  `json:"-"`
	DeletedOn   string  `json:"-"`
}

type ProductList []*Product

func (product *Product) FromJson(reader io.Reader) error {
	return json.NewDecoder(reader).Decode(product)
}

func (products *ProductList) ToJSON(writer io.Writer) error {
	return json.NewEncoder(writer).Encode(products)
}

// GetProducts returns a list of products
func GetProducts() ProductList {
	return productList1
}

func UpdateSingleProduct(id int, updatedProduct *Product) error {
	_, productIndex, err := FindProductById(id)
	if err != nil {
		return err
	}
	updatedProduct.ID = id
	productList1[productIndex] = updatedProduct
	return nil
}

var ErrorProductNotFound = fmt.Errorf("Product not found")

func FindProductById(id int) (*Product, int, error) {
	for currentIndex, currentProduct := range productList1 {
		if currentProduct.ID == id {
			return currentProduct, currentIndex, nil
		}
	}
	return nil, -1, ErrorProductNotFound
}

func CreateSingleProduct(newProduct *Product) {
	newProduct.ID = GetNextId()
	productList1 = append(productList1, newProduct)
}

func GetNextId() int {
	lastProduct := productList1[len(productList1)-1]
	return lastProduct.ID + 1
}

// productList is a hard coded list of products for this
// example data source
var productList1 = []*Product{
	&Product{
		ID:          1,
		Name:        "Latte",
		Description: "Frothy milky coffee",
		Price:       2.45,
		SKU:         "abc323",
		CreatedOn:   time.Now().UTC().String(),
		UpdatedOn:   time.Now().UTC().String(),
	},
	&Product{
		ID:          2,
		Name:        "Espresso",
		Description: "Short and strong coffee without milk",
		Price:       1.99,
		SKU:         "fjd34",
		CreatedOn:   time.Now().UTC().String(),
		UpdatedOn:   time.Now().UTC().String(),
	},
}
