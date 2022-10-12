package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

func homeLink(w http.ResponseWriter, r *http.Request) {
	_, _ = fmt.Fprintf(w, "Welcome to LFS(Leha file server)!")
}

func uploadImage(w http.ResponseWriter, r *http.Request) {
	file, handler, err := r.FormFile("file")
	handler.Filename = uuid.New().String()
	if err != nil {
		panic(err)
	}
	defer file.Close()
	f, err := os.OpenFile("./public/"+handler.Filename, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	_, _ = io.WriteString(w, "http://192.168.31.180:4554/public/"+handler.Filename)
	_, _ = io.Copy(f, file)
}

func main() {
	log.Println("Server will start at http://192.168.31.180:4554/")

	route := mux.NewRouter()

	fs := http.FileServer(http.Dir("./public/"))
	route.PathPrefix("/public/").Handler(http.StripPrefix("/public/", fs))

	route.HandleFunc("/", homeLink).Methods("GET")
	route.HandleFunc("/upload", uploadImage).Methods("POST")

	log.Fatal(http.ListenAndServe(":8000", route))
}
