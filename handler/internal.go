package handler

import (
	"fmt"
	"net/http"
)

func Create(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Internal /Create")
	return
}

func Read(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Internal /Read")
	return
}

func Update(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Internal /Update")
	return
}

func Delete(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Internal /Delete")
	return
}
