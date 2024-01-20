package main

import (
	"encoding/json"
	"fmt"
	httpSwagger "github.com/swaggo/http-swagger"
	"net/http"
)

// @title Swagger ms-rating
// @version 1.0
// @description service ms-rating

// @host localhost:8080
// @BasePath /

// @.apikey BearerAuth
// @in header
// @name X-Vkusvill-Token
// @securityDefinitions.apikey NumberAuth
// @in header
// @name number
func main() {
	mux := new(http.ServeMux)
	mux.HandleFunc("/home", home)
	mux.HandleFunc("/test", test)
	mux.HandleFunc("/swagger/*", httpSwagger.Handler(
		httpSwagger.URL("http://localhost:8080/swagger/doc.json")))
	http.ListenAndServe("localhost:8080", mux)
}

func home(w http.ResponseWriter, r *http.Request) {
	resp, err := homeHandler(r)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}
	w.Write(resp)
	return
}

func test(w http.ResponseWriter, r *http.Request) {
	resp, err := testHandler(r)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}
	w.Write(resp)
	return
}

func homeHandler(r *http.Request) ([]byte, error) {
	resp := fmt.Sprintf("[METHOD] : [%s] -[SHEMA] : [%s] - [HOST] : [%s] - [PATH] : [%s] - [RawQuery] : [%s]",
		r.Method, r.URL.Scheme, r.URL.Host, r.URL.Path, r.URL.RawQuery)
	respJSON, err := json.Marshal(resp)
	if err != nil {
		return nil, err
	}
	return respJSON, nil
}

func testHandler(r *http.Request) ([]byte, error) {
	resp := fmt.Sprintf("[METHOD] : [%s] -[SHEMA] : [%s] - [HOST] : [%s] - [PATH] : [%s] - [RawQuery] : [%s]",
		r.Method, r.URL.Scheme, r.URL.Host, r.URL.Path, r.URL.RawQuery)
	respJSON, err := json.Marshal(resp)
	if err != nil {
		return nil, err
	}
	return respJSON, nil
}
