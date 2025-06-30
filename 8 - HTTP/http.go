package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/home", func(w http.ResponseWriter, r *http.Request) {
		log.Println("Requisição recebida em /home")
		w.Write([]byte("Ola mundo"))
	})

	log.Println("Iniciando servidor em http://localhost:5000...")

	err := http.ListenAndServe(":5000", nil)
	if err != nil {
		log.Fatalf("Erro ao iniciar o servidor: %v", err)
	}
}
