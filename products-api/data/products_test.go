package data

import (
	"log"
	"testing"
)

func TestCheckValidator(t *testing.T) {

	p := &Product{
		Name:  "faiz",
		Price: 1.00,
		SKU:   "abs-cde-fgh",
	}

	log.Print("product ", p)

	err := p.Validator()

	if err != nil {
		t.Fatal(err)
	}
}
