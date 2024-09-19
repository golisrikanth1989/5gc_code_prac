package routes
import {
	"fmt"
	"net/http"
}

func NewRouter() http.Handler{
	mux:= http.NewServerMux()

	mux.HandleFunc("/",indexHandler)
	mux.HandleFunc("/api/data",apiDataHandler)

	return mux
}


func indexHandler(w http.ResponseWriter,r  *http.Request){
	fmt.Fprintln(w,"Welcome to Home Page!")
}

func apiDataHandler(w http.ResponseWriter,r *http.Request){
	data := "Some Data from API"
	fmt.Fprintln(w,data)
}