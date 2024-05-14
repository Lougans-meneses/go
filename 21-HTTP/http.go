package main

import (
	"log"
	"net/http"
)

func raiz(w http.ResponseWriter, r *http.Request)  {
	w.Write([]byte("Página Raiz"))
}

func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Olá mundo!"))
}

func usuarios(w http.ResponseWriter, r *http.Request)  {
	w.Write([]byte("Carregando página de usuário!"))
}

func main() {
	// HTTP É UM PROTOCOLO DE COMUNICAÇÃO - BASE DA COMUNICAÇÃO WEB

	// CLIENTE (FAZ REQUISIÇÃO) - SERVIDOR (PROCESSA REQUISIÇÃO E ENVIA RESPOSTA)

	// Request - Response

	// Rotas
	// URI - Identificador do Recurso
	// Método - GET, POST, PUT, DELETE

	http.HandleFunc("/home", home)

	http.HandleFunc("/usuario", usuarios)

	http.HandleFunc("/", raiz)

	log.Fatal(http.ListenAndServe(":5000", nil))
}