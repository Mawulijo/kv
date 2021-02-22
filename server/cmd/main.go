package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"kv/store"
	"log"
	"net/http"
)

var fs = store.NewFileStore("kv-store", ".fileStore")
//var ms = store.NewInMemoryStore("kv-store")


func handleSet(w http.ResponseWriter, r *http.Request){
	w.WriteHeader(http.StatusOK)
	//queryParams := r.URL.Query() // query-based matching
	vars := mux.Vars(r) // path-based matching
	err := fs.Set(vars["key"], vars["val"])
	//err = fs.Set( queryParams["key"][0],  queryParams["val"][0])
	if err != nil {
		http.Error(w,"Could not save record", http.StatusInternalServerError)
	}
}

func handleGet(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	vars := mux.Vars(r)
	value, err := fs.Get(vars["key"])
	if err != nil {
		http.Error(w, fmt.Sprintf("No value for the key: %s", vars["key"]), http.StatusNotFound)
	}
	_, err = fmt.Fprintf(w, value)
	if err != nil {
		http.Error(w,"Something went wrong", http.StatusInternalServerError)
	}
}

// Not working yet
func handleDelete(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	vars := mux.Vars(r)
	err := fs.Delete(vars["key"])
	if err != nil {
		http.Error(w, fmt.Sprintf("Could not delete a non-existent key: %s", vars["key"]), http.StatusNotFound)
	}
}
func handlePing(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("PONG"))
}

func main(){

	r := mux.NewRouter()
	r.HandleFunc("/v1/kv", handlePing)
	r.HandleFunc("/v1/kv/{key:[A-Za-z]+}/{val:[0-9]+}", handleSet) // path-based matching
	//r.HandleFunc("/v1/kv/, handleSet) // query-based matching
	r.HandleFunc("/v1/kv/{key:[A-Za-z]+}", handleGet).Methods("GET")
	r.HandleFunc("/v1/kv/{key:[A-Za-z]+}", handleDelete).Methods("POST").Name("deleteRecord")
	fmt.Println("kv is starting")
	fmt.Println("Ready to accept connections")
	log.Fatal(http.ListenAndServe(":1024", r))
}
