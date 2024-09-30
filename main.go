package main

import (
	"fmt"
	"io/ioutil"
	"main/posts"
	"main/split"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func getPostsHandler(w http.ResponseWriter, r *http.Request) {
	// Configura el encabezado Content-Type como application/json
	w.Header().Set("Content-Type", "application/json")

	hashtag := r.URL.Query().Get("hashtag")
	fmt.Println(hashtag)

	posts.Posts(hashtag)
	split.Split()

	// Lee el archivo JSON generado
	filePath := "extracted_posts.json"
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error al leer el archivo JSON: %v", err), http.StatusInternalServerError)
		return
	}

	// Establece el encabezado Content-Type como application/json
	w.Header().Set("Content-Type", "application/json")

	// Escribe el contenido del archivo JSON en la respuesta
	w.Write(data)
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/posts", getPostsHandler).Methods("GET")

	// Configura CORS
	corsHandler := handlers.CORS(
		handlers.AllowedOrigins([]string{"*"}),
	)(r)
	fmt.Println("server iniciado")
	http.ListenAndServe(":8000", corsHandler)
}
