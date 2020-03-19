package main

import (
	"database/sql"
	"fmt"

	u "glyph/user"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 54321
	user 	 = "postgres"
	password = "password"
	dbname   = "glyphdb"
)

func main() {

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	s := u.NewService(db)
	s.Insert("Example")
	fmt.Println(s.Get(1))

	// // database connection
	// db, err := sql.Open("pq",
	// 	"user:password@tcp(127.0.0.1:3306)/hello")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// defer db.Close()

	// r := chi.NewRouter()

	// r.Use(middleware.Timeout(60 * time.Second))

	// r.Get("/", func(w http.ResponseWriter, r *http.Request) {
	// 	w.Write([]byte("hi"))
	// })

	// http.ListenAndServe(":8080", r)
}
