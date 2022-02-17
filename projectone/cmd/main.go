package main

import (
	"encoding/json"
	"fmt"
	"mangtas-code-challenge/projectone/text"

	"net/http"
	"os"
)

type ServerConfig struct {
	Port int `json:"port"`
}

var config ServerConfig

func main() {
	loadServerConfig()

	http.HandleFunc("/", checkRequestMethod(mainRouteHandler, http.MethodGet))
	http.HandleFunc("/search", checkRequestMethod(text.HandleSearch, http.MethodPost))
	fmt.Println("Server started on port:", config.Port)
	http.ListenAndServe(fmt.Sprintf(":%d", config.Port), nil)

}

func mainRouteHandler(w http.ResponseWriter, _ *http.Request) {
	fmt.Fprintf(w, "Hello World!")
}

func loadServerConfig() {
	//load config json
	f, err := os.Open("config.json")
	if err != nil {
		exitf(err.Error())
	}
	defer f.Close()

	decoder := json.NewDecoder(f)
	err = decoder.Decode(&config)
	if err != nil {
		exitf(err.Error())
	}
}

func exitf(s string, args ...interface{}) {
	errorf(s, args...)
	os.Exit(1)
}

func errorf(s string, args ...interface{}) {
	fmt.Fprintf(os.Stderr, s+"\n", args...)
}

//middleware to check request method
func checkRequestMethod(next http.HandlerFunc, allowedMethod string) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != allowedMethod {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}
		next.ServeHTTP(w, r)
	})
}
