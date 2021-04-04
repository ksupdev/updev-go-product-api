package handlers

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/ksupdev/updev-go-product-api/data"
)

// swagger:route DELETE /products/{id} products deleteProduct
// Return a list of products from the database
// responses:
//	201: noContent

// DeleteProduct delete products from database
func (p *Products) DeleteProduct(rw http.ResponseWriter, r *http.Request) {
	// this will always convert because of the router
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	p.l.Println("Handler DELETE Product", id)

	err := data.DeleteProduct(id)
	if err == data.ErrProductNotFound {
		http.Error(rw, "Product not found", http.StatusNotFound)
		return
	}

	if err != nil {
		http.Error(rw, "Product not found", http.StatusInternalServerError)
		return
	}

}
