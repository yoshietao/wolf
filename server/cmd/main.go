package main

import (
	"database/sql"
	"fmt"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/yoshietao/wolf/server/router"
)

func main() {
	fmt.Println("Hello world.")

	db, err := sql.Open("mysql", "wolf:wolf@/wolf")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()
	fmt.Println("Success!")

	handler := router.Register(db)

	http.ListenAndServe(":9000", handler)
}
