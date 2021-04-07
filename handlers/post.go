package handlers

import (
	"net/http"
	"reflect"

	"github.com/ksupdev/updev-go-product-api/data"
)

// swagger:route POST /products products createProduct
// Create a new product
//
// responses:
//	200: productResponse
//  422: errorValidation
//  501: errorResponse

// Create handles POST requests to add new products
func (p *Products) Create(rw http.ResponseWriter, r *http.Request) {
	// fetch the product from the context
	p.l.Println("------Handle Create------")
	prod := r.Context().Value(KeyProduct{})
	p.l.Println("------Check data type of value in interface------ ", reflect.TypeOf(prod))

	if _, ok := prod.(*data.Product); ok {
		p.l.Printf("[DEBUG] Inserting product: %#v\n", prod)
		data.AddProduct(*(prod.(*data.Product)))
	} else {
		p.l.Println("[ERROR] Not okay")
	}
}

// Create handles POST requests to add new products
// func (p *Products) Create(rw http.ResponseWriter, r *http.Request) {
// 	// fetch the product from the context
// 	prod := r.Context().Value(KeyProduct{}).(data.Product)

// 	p.l.Printf("[DEBUG] Inserting product: %#v\n", prod)
// 	data.AddProduct(prod)
// }
