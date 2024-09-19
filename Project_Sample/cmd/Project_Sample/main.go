package main

import (
	"fmt"
	"net/http"

	"github.com/golisrikanth1989/5gc_code_prac/Project_Sample/internal/routes"
)

func main() {

	router := routes.NewRouter()

	port := 8000
	addr := fmt.Sprintf("%d", port)
	fmt.Printf("Server Listening on http://localhost%s\n", addr)
	err := http.ListenAndServe(addr, router)
	if err != nil {
		panic(err)
	}

}
