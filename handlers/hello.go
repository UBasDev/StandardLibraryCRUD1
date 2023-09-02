package handlers

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
)

type Person struct {
	Name string `json:"name"`
	Age  int8   `json:"age"`
}

type TypeHelloService struct {
	log *log.Logger
}

type IHelloService interface {
	ServeHTTP(responseWriter http.ResponseWriter, request *http.Request) //(*[]Person, error)
}

func NewHelloService(log *log.Logger) IHelloService {
	return &TypeHelloService{
		log: log,
	}
}

func (typeHelloService *TypeHelloService) ServeHTTP(responseWriter http.ResponseWriter, request *http.Request) { //(*[]Person, error)

	defer request.Body.Close()
	responseWriter.Header().Add(`Content-Type`, `application/json`)

	typeHelloService.log.Println("Hello world!")
	responseWriter.WriteHeader(http.StatusAccepted)

	var requestBody, err = io.ReadAll(request.Body)
	if err != nil {
		typeHelloService.log.Println()
	}
	json.NewEncoder(responseWriter).Encode(requestBody)

}
