package main

import (
	"fmt"
	"net/http"

	"github.com/yoshietao/wolf/server/router"
)

func main() {
	fmt.Println("Hello world.")

	handler := router.Register(nil)

	http.ListenAndServe(":9000", handler)
}
