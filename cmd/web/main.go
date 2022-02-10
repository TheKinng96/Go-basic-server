package main

import (
	"fmt"
	"net/http"

	"github.com/TheKinng96/Go-basic-server/pkg/controllers"
)

const portNumber = ":8000"

func main() {
	http.HandleFunc("/", controllers.Home)
	http.HandleFunc("/about", controllers.About)

	fmt.Println(fmt.Sprintf("Starting application on port %s", portNumber))
	_ = http.ListenAndServe(portNumber, nil)
}
