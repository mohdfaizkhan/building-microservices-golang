package handlers

import (
	"log"
	"net/http"

	"github.com/mohdfaizkhan/building-microservices-golang/products-api/data"
)

type Products struct {
	l *log.Logger
}

func NewProducts(l *log.Logger) *Products {
	return &Products{l}
}

// ServeHTTP implements the go http.Handler interface
// https://golang.org/pkg/net/http/#Handler
func (p *Products) ServeHTTP(rw http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodGet {
		p.getProducts(rw, r)
	}
	rw.WriteHeader(http.StatusMethodNotAllowed)

}

func (p *Products) getProducts(rw http.ResponseWriter, r *http.Request) {
	lp := data.GetProducts()

	//d, err := json.Marshal(lp)

	err := lp.ToJSON(rw)

	if err != nil {
		http.Error(rw, "Unable to parse the data", http.StatusInternalServerError)
	}

	//rw.Write(d)
}