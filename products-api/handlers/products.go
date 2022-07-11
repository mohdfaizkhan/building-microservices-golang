package handlers

import (
	"context"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/mohdfaizkhan/building-microservices-golang/products-api/data"
)

// Products is a http.Handler
type Products struct {
	l *log.Logger
}

// NewProducts creates a products handler with the given logger
func NewProducts(l *log.Logger) *Products {
	return &Products{l}
}

// ServeHTTP is the main entry point for the handler and staisfies the http.Handler
// interface
// ServeHTTP implements the go http.Handler interface
// https://golang.org/pkg/net/http/#Handler
// func (p *Products) ServeHTTP(rw http.ResponseWriter, r *http.Request) {

// 	if r.Method == http.MethodGet {
// 		p.getProducts(rw, r)
// 	}

// 	if r.Method == http.MethodPost {
// 		p.addProduct(rw, r)
// 	}

// 	if r.Method == http.MethodPut {
// 		p.l.Printf("PUT Products")
// 		reg := regexp.MustCompile(`/([0-9]+)`)
// 		g := reg.FindAllStringSubmatch(r.URL.Path, -1)

// 		if len(g) != 1 {
// 			http.Error(rw, "Invalid URL", http.StatusInternalServerError)
// 			return
// 		}
// 		if len(g[0]) != 2 {
// 			http.Error(rw, "Invalid URL", http.StatusInternalServerError)
// 			return
// 		}

// 		idString := g[0][1]

// 		id, err := strconv.Atoi(idString)

// 		if err != nil {
// 			http.Error(rw, "Invalid URL", http.StatusInternalServerError)
// 			return
// 		}

// 		p.l.Printf("got Id :", id)

// 		p.updateProducts(id, rw, r)
// 	}

// 	rw.WriteHeader(http.StatusMethodNotAllowed)

// }

// getProducts returns the products from the data store
func (p *Products) GetProducts(rw http.ResponseWriter, r *http.Request) {
	p.l.Printf("Handle GET Products")

	lp := data.GetProducts()

	//d, err := json.Marshal(lp)

	err := lp.ToJSON(rw)

	if err != nil {
		http.Error(rw, "Unable to parse the data", http.StatusInternalServerError)
	}

	//rw.Write(d)
}

func (p *Products) AddProduct(rw http.ResponseWriter, r *http.Request) {
	p.l.Printf("Handle Post Product")

	//prod := &data.Product{}

	//err := prod.FromJSON(r.Body)

	//if err != nil {
	//	http.Error(rw, "Unable to unmarshall the request", http.StatusBadRequest)
	//}

	prod := r.Context().Value(KeyProduct{}).(data.Product)
	p.l.Printf("Prod: %#v", prod)

	data.AddProduct(&prod)
}

func (p *Products) UpdateProducts(rw http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	id, errForConvert := strconv.Atoi(vars["id"])

	if errForConvert != nil {
		http.Error(rw, "Bad request", http.StatusBadRequest)
		return
	}

	p.l.Print("Handle PUT Product", id)

	prod := r.Context().Value(KeyProduct{}).(data.Product)

	//prod := &data.Product{}

	p.l.Print("Product", prod)

	/* 	err := prod.FromJSON(r.Body)

	   	p.l.Print("Product---", prod)

	   	if err != nil {
	   		http.Error(rw, "Unable to unmarshall the json", http.StatusBadRequest)
	   		return
	   	}
	*/
	excError := data.UpdateProduct(id, &prod)

	if excError == data.ErrProductNotFound {
		http.Error(rw, "Product not found", http.StatusNotFound)
		return
	}

	if excError != nil {
		http.Error(rw, "Product not found", http.StatusInternalServerError)
		return
	}

}

type KeyProduct struct {
}

func (p Products) MiddlewareProductValidation(next http.Handler) http.Handler {

	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {

		prod := data.Product{}

		err := prod.FromJSON(r.Body)
		log.Print("data print ", prod)
		if err != nil {
			http.Error(rw, "Unable to unmarshall the json", http.StatusBadRequest)
			return
		}

		ctx := context.WithValue(r.Context(), KeyProduct{}, prod)

		req := r.WithContext(ctx)

		next.ServeHTTP(rw, req)
	})
}
