package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("Hello world.")

	handler := h

	http.ListenAndServe(":9000", handler)
}
