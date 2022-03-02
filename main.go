package main

import (
	"fmt"
	"net/http"

	"github.com/movie42/golang/myapp"
)

func main() {
	fmt.Print("http://localhost:3000")
	http.ListenAndServe(":3000", myapp.NewHttpHandler())
}
