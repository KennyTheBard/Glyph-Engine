package user

import (
	"database/sql"
	"log"
)

type Service struct {
	db *sql.DB
}

func NewService(db *sql.DB) *Service {
	s := new(Service)
	s.db = db

	return s
}


func (s Service) Insert(name string) {
	_, err := s.db.Exec("insert into users (first_name) values($1)", name)
	if err != nil {
		log.Fatal(err)
	}
}

func (s Service) Get(id int) string {
	var name string
	err := s.db.QueryRow("select first_name from users where id = $1", id).Scan(&name)
	if err != nil {
		log.Fatal(err)
	}
	return name
}
