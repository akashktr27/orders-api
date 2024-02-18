package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	fmt.Println("hello world from new account")

	router := chi.NewRouter()
	router.Use(middleware.Logger)

	router.Get("/hello", basicHandler)

	server := &http.Server{
		Addr: ":3000",
		// Handler: http.HandlerFunc(basicHandler),
		Handler: router,
	}
	err := server.ListenAndServe()

	if err != nil {
		fmt.Println("faild to listen to server", err)
	}

}

func basicHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello world"))
}
