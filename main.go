package main

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type LogMiddleware struct {
	http.Handler
}

func (middleware *LogMiddleware) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	fmt.Println("receive request")
	middleware.Handler.ServeHTTP(writer, request)
}

func main() {
	router := httprouter.New()
	router.PanicHandler = func(w http.ResponseWriter, r *http.Request, error interface{}) {
		fmt.Fprint(w, "Panic : ", error)
	}
	router.NotFound = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Sorry, we can't find the page you looking for")
	})

	router.MethodNotAllowed = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Sorry method is not allowed")
	})

	router.GET("/", func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		// panic("Opps, something went wrong!")
		fmt.Fprint(w, "Hello http router")
	})

	middleware := LogMiddleware{router}

	server := http.Server{
		Handler: &middleware,
		Addr:    "localhost:3000",
	}

	server.ListenAndServe()
}
