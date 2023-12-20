package gvalidation

import (
	"testing"

	"github.com/SyaibanAhmadRamadhan/gocatch/gtypedata/gmap"
)

type Address struct {
	City    string `json:"city"    validate:"required"`
	Country string `json:"country" validate:"required"`
}

var validate *Validation

func init() {
	validate = New()
	validate.SetIdTranslation()
}

type Request struct {
	Name    string  `json:"name" validate:"required,min=10"`
	Address Address `json:"address" validate:"required"`
}

func TestValidation(t *testing.T) {
	req := Request{
		Name: "",
		Address: Address{
			City:    "",
			Country: "",
		},
	}

	data := validate.StructM(req)
	gmap.Println(data)
}

func TestValidation2(t *testing.T) {
	req := Request{
		Name: "2",
		Address: Address{
			City:    "",
			Country: "",
		},
	}

	data := validate.StructM(req)
	gmap.Println(data)
}
