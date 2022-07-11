package data

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"regexp"
	"time"

	"github.com/go-playground/validator/v10"
)

// Product defines the structure for an API product
type Product struct {
	ID          int     `json:"id"`
	Name        string  `json:"name" validate:"required"`
	Description string  `json:"description"`
	Price       float32 `json:"price" validate:"gt=0"`
	SKU         string  `json:"sku" validate:"required,sku"`
	CreatedOn   string  `json:"_"`
	UpdatedOn   string  `json:""`
	DeletedOn   string  `json:""`
}

// Products is a collection of Product
type Products []*Product

// ToJSON serializes the contents of the collection to JSON
// NewEncoder provides better performance than json.Unmarshal as it does not
// have to buffer the output into an in memory slice of bytes
// this reduces allocations and the overheads of the service
//
// https://golang.org/pkg/encoding/json/#NewEncoder
func (p *Products) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)

	return e.Encode(p)
}

// FromJson deserializes the contents of the collection to JSON
// NewDecoder provides better performance than json.Unmarshal as it does not
// have to buffer the output into an in memory slice of bytes
// this reduces allocations and the overheads of the service
//
// https://golang.org/pkg/encoding/json/#NewDecoder

func (p *Product) FromJSON(r io.Reader) error {
	e := json.NewDecoder(r)
	return e.Decode(p)

}

// validator the input value filed
//
//https://github.com/go-playground/validator
func (p *Product) Validator() error {

	validate := validator.New()
	validate.RegisterValidation("sku", validateSKU)
	return validate.Struct(p)
}

func validateSKU(fl validator.FieldLevel) bool {
	// sku is of format abc-absd-dfsdf
	re := regexp.MustCompile(`[a-z]+-[a-z]+-[a-z]+`)
	matches := re.FindAllString(fl.Field().String(), -1)

	if len(matches) != 1 {
		return false
	}
	return true
}

// GetProducts returns a list of products
func GetProducts() Products {
	return productList
}

func AddProduct(p *Product) {
	p.ID = getNextId()
	productList = append(productList, p)
}

func UpdateProduct(id int, p *Product) error {

	_, pos, err := findPrdouct(id)

	if err != nil {
		return err
	}

	p.ID = id

	log.Print("pro", productList[pos])
	productList[pos] = p

	return nil
}

var ErrProductNotFound = fmt.Errorf("Product Not Found")

func findPrdouct(id int) (*Product, int, error) {

	log.Print("product id inside findProduct-----", id)
	log.Print("product id inside findProduct-----", productList)
	for id, p := range productList {
		log.Print("product id inside findProduct-----loop   ", p.ID)
		if p.ID != id {
			log.Print("product id inside findProduct-----loop  ", p.ID)
			return p, id, nil
		}
	}

	return nil, id, ErrProductNotFound
}

func getNextId() int {

	lp := productList[len(productList)-1]
	return lp.ID + 1
}

// productList is a hard coded list of products for this
// example data source
var productList = Products{
	&Product{
		ID:          1,
		Name:        "Latte",
		Description: "Frothy milky coffee",
		Price:       2.45,
		SKU:         "abc323",
		CreatedOn:   time.Now().UTC().String(),
		UpdatedOn:   time.Now().UTC().String(),
		DeletedOn:   "",
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
