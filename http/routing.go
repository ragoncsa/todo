package http

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func InitServer() *mux.Router {
	return mux.NewRouter()
}

func StartServer(router *mux.Router) {
	fmt.Println("Server at 8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
