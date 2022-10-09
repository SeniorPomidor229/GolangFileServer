package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

func homeLink(w http.ResponseWriter, r *http.Request) {
	_, _ = fmt.Fprintf(w, "Welcome to LIS(Liha image server) 2.0!")
}

func uploadImage(w http.ResponseWriter, r *http.Request) {
	file, handler, err := r.FormFile("file")
	handler.Filename = uuid.New().String() + strings.TrimPrefix(handler.Filename, ".")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	f, err := os.OpenFile("./public/"+handler.Filename, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	_, _ = io.WriteString(w, "File url http://localhost:8000/public/"+handler.Filename)
	_, _ = io.Copy(f, file)
}

func main() {
	log.Println("Server will start at http://localhost:8000/")

	route := mux.NewRouter()

	fs := http.FileServer(http.Dir("./public/"))
	route.PathPrefix("/public/").Handler(http.StripPrefix("/public/", fs))

	route.HandleFunc("/", homeLink).Methods("GET")
	route.HandleFunc("/upload", uploadImage).Methods("POST")

	log.Fatal(http.ListenAndServe(":8000", route))
}
