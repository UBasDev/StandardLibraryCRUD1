package handlers

import (
	"log"
	"net/http"
	"regexp"
	"strconv"

	"example.com/goproject6/data"
)

type Products struct {
	log *log.Logger
}

func NewProducts(log *log.Logger) *Products {
	return &Products{
		log: log,
	}
}

func (products *Products) ServeHTTP(responseWriter1 http.ResponseWriter, request *http.Request) {
	responseWriter1.Header().Set("Content-Type", "application/json")
	if request.Method == http.MethodGet {
		products.getAllProducts(responseWriter1, request)
		return
	}
	if request.Method == http.MethodPost {
		products.createSingleProduct(responseWriter1, request)
		return
	}
	if request.Method == http.MethodPut {
		reg := regexp.MustCompile(`/([0-9]+)`)
		g := reg.FindAllStringSubmatch(request.URL.Path, -1)
		if len(g) != 1 {
			products.log.Println("Invalid URI more than one id")
			http.Error(responseWriter1, "Invalid URI", http.StatusBadRequest)
			return
		}
		if len(g[0]) != 2 {
			products.log.Println("Invalid URI more than one capture group")
			http.Error(responseWriter1, "Invalid URI", http.StatusBadRequest)
			return
		}
		idString := g[0][1]
		id, err := strconv.Atoi(idString)
		if err != nil {
			products.log.Println("Invalid URI unable to convert to number", id)
			http.Error(responseWriter1, "Cannot convert the ID parameter", http.StatusBadRequest)
			return
		}
		products.updateSingleProduct(id, responseWriter1, request)
		return
	}
	responseWriter1.WriteHeader(http.StatusMethodNotAllowed)
}

func (products *Products) getAllProducts(responseWriter1 http.ResponseWriter, request *http.Request) {

	products.log.Println("Handle GET Products")
	productList1 := data.GetProducts()
	err := productList1.ToJSON(responseWriter1)
	if err != nil {
		http.Error(responseWriter1, "Unable to marshal json", http.StatusInternalServerError)
	}

}

func (products *Products) createSingleProduct(responseWriter1 http.ResponseWriter, request *http.Request) {

	products.log.Println("Handle POST Products")
	productToCreate := &data.Product{}
	err := productToCreate.FromJson(request.Body)
	if err != nil {
		http.Error(responseWriter1, "Unable to unmarshal json", http.StatusInternalServerError)
	}
	data.CreateSingleProduct(productToCreate)

}
func (products *Products) updateSingleProduct(id int, responseWriter1 http.ResponseWriter, request *http.Request) {

	products.log.Println("Handle PUT Products")
	productToUpdate := &data.Product{}
	err := productToUpdate.FromJson(request.Body)
	if err != nil {
		http.Error(responseWriter1, "Unable to unmarshal json", http.StatusInternalServerError)
	}
	err = data.UpdateSingleProduct(id, productToUpdate)
	if err == data.ErrorProductNotFound {
		http.Error(responseWriter1, "Product not found", http.StatusBadRequest)
		return
	}
	if err != nil {
		http.Error(responseWriter1, "Product not found", http.StatusInternalServerError)
		return
	}
}
