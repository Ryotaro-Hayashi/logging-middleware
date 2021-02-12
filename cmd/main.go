package main

import (
	"fmt"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Hello, middleware!")
}

func middleware1(nextFunc http.HandlerFunc) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("[START] middleware1")
		nextFunc.ServeHTTP(w, r)
		fmt.Println("[END] middleware2")
	}
}

func middleware2(nextFunc http.HandlerFunc) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("[START] middleware1")
		nextFunc.ServeHTTP(w, r)
		fmt.Println("[END] middleware2")
	}
}

func middleware3(nextFunc http.HandlerFunc) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("[START] middleware1")
		nextFunc.ServeHTTP(w, r)
		fmt.Println("[END] middleware2")
	}
}

func main() {
	mux := http.NewServeMux()
	// 第2引数は「func(ResponseWriter, *Request)」
	mux.HandleFunc("/hello", middleware1(middleware2(middleware3(helloHandler))))

	http.ListenAndServe(":8080", mux)
}
