package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/rs/cors"
	"github.com/google/uuid"
	// "github.com/gorilla/handlers"
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
	_, _ = io.WriteString(w, "https://lis.4dev.kz/public/"+handler.Filename)
	_, _ = io.Copy(f, file)
}

func main() {
	log.Println("Server will start at https://lis.4dev.kz/")

	route := mux.NewRouter()

	fs := http.FileServer(http.Dir("./public/"))
	route.PathPrefix("/public/").Handler(http.StripPrefix("/public/", fs))

	route.HandleFunc("/", homeLink).Methods("GET")
	route.HandleFunc("/upload", uploadImage).Methods("POST")


	c := cors.New(cors.Options{
        AllowedOrigins: []string{"*"},
        AllowCredentials: true,
    })

	handler := c.Handler(route)

	log.Fatal(http.ListenAndServe(":5000", handler))
}
