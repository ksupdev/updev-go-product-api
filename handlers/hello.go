package handlers

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Hello struct {
}

func (h *Hello) ServerHTTP(rw http.ResponseWriter, r *http.Request) {
	log.Println("Hello world")
	d, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(rw, "Opps", http.StatusBadRequest)
	}

	//log.Printf("Data %s\n", d)

	fmt.Fprintf(rw, " Hello %s ", d)
}
