package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

const port string = ":8000"

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Servidor Activo en index")
	})

	log.Println("Servidor corriendo en el puerto ", port)
	log.Fatalln(http.ListenAndServe(port, router))

}
