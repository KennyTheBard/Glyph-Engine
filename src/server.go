package main

import (
	"flag"
	"fmt"
	"strconv"

	"net/http"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"

	config "./config"
	data "./src/data"
	security "./src/security"
	admin "./src/web/admin"
	user "./src/web/user"
)

var router *gin.Engine

func main() {

	// database connection
	db, err := sql.Open("pq",
		"user:password@tcp(127.0.0.1:3306)/hello")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	r := chi.NewRouter()

	r.Use(middleware.Timeout(60 * time.Second))

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hi"))
	})

	http.ListenAndServe(":8080", r)
}
