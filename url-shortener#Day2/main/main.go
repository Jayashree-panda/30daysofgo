package main

import(
	"fmt"
	"net/http"
	"github.com/Jayashree-panda/30daysofgo/url_shortener#day2"
)

func main(){
	mux := DefaultMux()
	
	pathsToUrls := map[string]string {
		"/urlshort-godoc": "https://godoc.org/github.com/gophercises/urlshort",
		"/yaml-godoc":     "https://godoc.org/gopkg.in/yaml.v2",
	}

	mapHandler := urlshort.MapHandler(pathsToUrls, mux)

}

func DefaultMux() *http.ServeMux{
	mux := http.NewServeMux()
	mux.HandleFunc('/',hello)
	return mux
}

func hello(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "Hello world")
}