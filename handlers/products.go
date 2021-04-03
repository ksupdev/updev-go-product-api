// Package classification of Product API
//
// Documentation for Product API
//
//	Schemes: http
//	BasePath: /
//	Version: 1.0.0
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
// swagger:meta
package handlers

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/ksupdev/updev-go-product-api/data"
)

// A list of products
// swagger:response productsResponse
type productsResponseWrapprt struct {
	// All current products
	// in: body
	Body []data.Product
}

type Products struct {
	l *log.Logger
}

func NewProducts(l *log.Logger) *Products {
	return &Products{l}
}

// swagger:route GET /products products listProducts
// Return a list of products from the database
// responses:
//	200: productsResponse

// GetProduct returns the products from the data store
func (p *Products) GetProducts(rw http.ResponseWriter, r *http.Request) {
	lp := data.GetProducts()
	err := lp.ToJson(rw)
	if err != nil {
		http.Error(rw, "Unable to marshal json", http.StatusInternalServerError)
	}
}

func (p *Products) AddProduct(rw http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle POST Product")
	prod := r.Context().Value(KeyProduct{}).(data.Product)
	data.AddProduct(&prod)
}

func (p *Products) UpdateProducts(rw http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(rw, "Unable to convert ID", http.StatusBadRequest)
		return
	}

	p.l.Println("Handle PUT Product ", id)
	prod := r.Context().Value(KeyProduct{}).(data.Product)
	err = data.UpdateProducts(id, &prod)
	if err == data.ErrProductNotFound {
		http.Error(rw, "Product not found", http.StatusNotFound)
		return
	}

	if err != nil {
		http.Error(rw, "Product not found", http.StatusInternalServerError)
		return
	}
}

type KeyProduct struct{}

func (p *Products) MiddlewareProductValidation(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		prod := data.Product{}

		err := prod.FronJSON(r.Body)
		if err != nil {
			p.l.Println("[ERROR] deserializibg product", err)
			http.Error(rw, "Unable to unmarshal json", http.StatusBadRequest)
			return
		}

		// validate the product
		err = prod.Validate()
		if err != nil {
			p.l.Println("[ERROR] validate product", err)
			http.Error(
				rw,
				fmt.Sprintf("Unable to validate product : %s", err),
				http.StatusBadRequest)
			return
		}

		ctx := context.WithValue(r.Context(), KeyProduct{}, prod)
		req := r.WithContext(ctx)

		next.ServeHTTP(rw, req)

	})
}

// LoggingMiddleware - adds middleware around endpoints
func (p *Products) LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		p.l.Println("Log.. request ", r.Header)
		next.ServeHTTP(w, r)
	})
}
