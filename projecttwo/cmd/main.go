package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"text/template"
)

type ServerConfig struct {
	Port   int    `json:"port"`
	APIURL string `json:"api_url"`
}

type homeData struct {
	Title      string
	Data       map[string]int64
	Message    string
	TextSearch string
}

var config ServerConfig

func main() {
	loadServerConfig()
	//tmpl.DefinedTemplates()

	// load and render views/index.html
	http.HandleFunc("/", renderHome)
	port := fmt.Sprintf(":%d", config.Port)
	fmt.Println("Server started on port " + port)
	http.ListenAndServe(port, nil)

}

func renderHome(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "text/html")
	vd := homeData{
		Title: "Mangtas Code Challenge",
	}
	tmpl, err := parseHTMLFile()

	if err != nil {
		vd.Data = nil
		vd.Message = "Something went wrong"
		log.Printf("Error :%s", err.Error())
		executeTemplate(tmpl, w, vd)
		return
	}

	if r.Method == "POST" {
		renderSearch(w, r, tmpl)
		return
	}

	executeTemplate(tmpl, w, vd)
}

func executeTemplate(tmpl *template.Template, w http.ResponseWriter, vd homeData) {
	err := tmpl.ExecuteTemplate(w, "indexHTML", vd)
	if err != nil {
		log.Printf("Error :%s", err.Error())

		return
	}
	return
}

func parseHTMLFile() (*template.Template, error) {
	filePrefix, _ := filepath.Abs("./")
	tmpl, err := template.ParseFiles(filePrefix + "/views/index.html")
	if err != nil {
		log.Printf("Error :%s", err.Error())
		return nil, err
	}
	return tmpl, nil
}

func renderSearch(w http.ResponseWriter, r *http.Request, tmpl *template.Template) {
	w.Header().Set("Content-Type", "text/html")
	vd := homeData{
		Title:      "Mangtas Code Challenge",
		Data:       nil,
		Message:    "",
		TextSearch: "",
	}

	text := r.FormValue("textSearch")
	vd.TextSearch = text
	if len(text) == 0 {
		log.Println("no text passed")
		vd.Message = "no text passed"
		executeTemplate(tmpl, w, vd)
		return
	}

	response, err := http.PostForm(config.APIURL+"/search", url.Values{"text": {text}})
	if err != nil {
		log.Printf("Error :%s", err.Error())
		vd.Message = "Something went wrong"
		executeTemplate(tmpl, w, vd)
		return
	}
	defer response.Body.Close()

	var result map[string]int64
	err = json.NewDecoder(response.Body).Decode(&result)
	if err != nil {
		log.Printf("Error :%s", err.Error())
		vd.Message = "Something went wrong"
		executeTemplate(tmpl, w, vd)

		return
	}
	vd.Data = result
	executeTemplate(tmpl, w, vd)

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
