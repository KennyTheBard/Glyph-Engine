package story

import (
	"database/sql"
	"log"

	"github.com/gin-gonic/gin"
)

type Service struct {
	db *sql.DB
}

func NewService(db *sql.DB) *Service {
	s := new(Service)
	s.db = db

	return s
}

func (s Service) Create(title, description string, authorId int) {
	_, err := s.db.Exec("INSERT INTO stories (title, description, author_id) "+
		"VALUES($1, $2, $3)", title, description, authorId)
	if err != nil {
		log.Fatal(err)
	}
}

func (s Service) GetById(id int) gin.H {
	var title, description string
	var authorId int
	err := s.db.QueryRow("SELECT title, description, author_id FROM stories WHERE id = $1", id).
		Scan(&title, &description, &authorId)
	if err != nil {
		log.Fatal(err)
	}

	return gin.H{
		"title":       title,
		"description": description,
		"author_id":   authorId,
	}
}

func (s Service) GetAll() []gin.H {
	rows, err := s.db.Query("SELECT title, description, author_id FROM stories")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	all := make([]gin.H, 0)
	for rows.Next() {
		var title, description string
		var authorId int
		err = rows.Scan(&title, &description, &authorId)
		if err != nil {
			log.Fatal(err)
		}

		all = append(all, gin.H{
			"title":       title,
			"description": description,
			"author_id":   authorId,
		})
	}

	if err = rows.Err(); err != nil {
		log.Fatal(err)
	}

	return all
}

func (s Service) Update(id int, title, description string, authorId int) {
	_, err := s.db.Exec("UPDATE stories SET title = $2, description = $3, author_id = $4 "+
		"WHERE id = $1", id, title, description, authorId)
	if err != nil {
		log.Fatal(err)
	}
}

func (s Service) Delete(id int) {
	_, err := s.db.Exec("DELETE FROM stories WHERE id = $1", id)
	if err != nil {
		log.Fatal(err)
	}
}
