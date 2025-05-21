package main

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/julienschmidt/httprouter"
	"github.com/stretchr/testify/assert"
)

func TestRouter(t *testing.T) {
	router := httprouter.New()
	router.GET("/", func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		fmt.Fprint(w, "Hello world")
	})

	request := httptest.NewRequest("GET", "http://localhost:3000/", nil)
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)
	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)
	fmt.Println(body)
	assert.Equal(t, "Hello world", string(body))
}

func TestRouterParams(t *testing.T) {
	router := httprouter.New()
	router.GET("/products/:id", func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		text := "Product " + p.ByName("id")
		fmt.Fprint(w, text)
	})

	request := httptest.NewRequest("GET", "http://localhost:3000/products/1", nil)
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)
	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)
	fmt.Println(string(body))
	assert.Equal(t, "Product 1", string(body))
}

func TestRouterPatternNamedParams(t *testing.T) {
	router := httprouter.New()
	router.GET("/products/:id/items/:itemId", func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		text := "Product " + p.ByName("id") + " Items " + p.ByName("itemId")
		fmt.Fprint(w, text)
	})

	request := httptest.NewRequest("GET", "http://localhost:3000/products/1/items/3", nil)
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)
	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)
	fmt.Println(string(body))
	assert.Equal(t, "Product 1 Items 3", string(body))
}

func TestRouterPatternCatchAllParams(t *testing.T) {
	router := httprouter.New()
	router.GET("/images/*image", func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		text := "Images " + p.ByName("image")
		fmt.Fprint(w, text)
	})

	request := httptest.NewRequest("GET", "http://localhost:3000/images/small/profile.png", nil)
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)
	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)
	fmt.Println(string(body))
	assert.Equal(t, "Images /small/profile.png", string(body))
}
