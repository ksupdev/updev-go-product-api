package handlers

import (
	"net/http"

	"github.com/ksupdev/updev-go-product-api/data"
)

// swagger:route GET /products products listProducts
// Return a list of products from the database
// responses:
//	200: productsResponse

// GetProduct returns the products from the data store
func (p *Products) GetProducts(rw http.ResponseWriter, r *http.Request) {

	p.l.Println("Handle get product")

	// fetch the products from the database
	lp := data.GetProducts()

	// serialize the list to json
	err := lp.ToJson(rw)
	if err != nil {
		http.Error(rw, "Unable to marshal json", http.StatusInternalServerError)
	}
}
