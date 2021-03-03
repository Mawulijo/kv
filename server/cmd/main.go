package main

import (
	"flag"
	"fmt"
	"github.com/gorilla/mux"
	config "kv"
	"kv/store"
	"log"
	"net/http"
)

var fs = store.NewFileStore("kv-store", ".fileStore")

//var ms = store.NewInMemoryStore("kv-store")

func handleSet(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	vars := mux.Vars(r)
	err := fs.Set(vars["key"], vars["val"])
	if err != nil {
		http.Error(w, fmt.Sprintf("Key already exists: %s", vars["key"]), http.StatusNotFound)
		return
	}
	fmt.Fprintf(w, vars["val"]+" OK")
}

func handleGet(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	vars := mux.Vars(r)
	value, err := fs.Get(vars["key"])
	if err != nil {
		http.Error(w, fmt.Sprintf("Could not get value for non-exixtent key: %s", vars["key"]), http.StatusNotFound)
	}
	fmt.Fprintf(w, value)
}
func handleUpdate(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	vars := mux.Vars(r)
	err := fs.Update(vars["key"], vars["val"])
	if err != nil {
		http.Error(w, fmt.Sprintf("Could not update value for non-exixtent key: %s", vars["key"]), http.StatusNotFound)
	}
	if err == nil {
		fmt.Fprintf(w, fmt.Sprintf("%s updated to %s", vars["key"], vars["val"]))
	}
}
func handleGetAll(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	//vars := mux.Vars(r)
	kv, err := fs.GetAll()
	if err != nil {
		http.Error(w, "Could not get all records", http.StatusNotFound)
	}
	_, err = fmt.Fprintf(w, kv)
	if err != nil {
		http.Error(w, "Something went wrong", http.StatusInternalServerError)
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

func main() {

	boolPtr := flag.Bool("port", false, "Provide this flag to connect to kv server. This ensures that a .config file is provided before the application starts.")
	flag.Parse()

	cfg := config.LoadConfig(*boolPtr)

	r := mux.NewRouter()
	r.HandleFunc("/v1/kv", handlePing)
	r.HandleFunc("/v1/kv/{key:[A-Za-z]+}/{val:[A-Za-z0-9]+}/set", handleSet).Methods("POST")
	r.HandleFunc("/v1/kv/{key:[A-Za-z]+}/{val:[A-Za-z0-9]+}/update", handleUpdate).Methods("POST")
	r.HandleFunc("/v1/kv/{key:[A-Za-z]+}/get", handleGet).Methods("GET")
	r.HandleFunc("/v1/kv/all", handleGetAll).Methods("GET")
	r.HandleFunc("/v1/kv/{key:[A-Za-z]+}/delete", handleDelete).Methods("POST").Name("deleteRecord")
	fmt.Println(fmt.Sprintf("kv is starting on port: %d", cfg.Port))
	fmt.Println("Ready to accept connections")
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", cfg.Port), r))
}
