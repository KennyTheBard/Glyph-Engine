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

func (s Service) Create(title, text string, storyId int) {
	_, err := s.db.Exec("INSERT INTO scenes (title, text, story_id) "+
		"VALUES($1, $2, $3)", title, text, storyId)
	if err != nil {
		log.Fatal(err)
	}
}

func (s Service) GetById(id int) gin.H {
	var title, text string
	var storyId int
	err := s.db.QueryRow("SELECT title, text, story_id FROM scenes WHERE id = $1", id).
		Scan(&title, &text, &storyId)
	if err != nil {
		log.Fatal(err)
	}

	return gin.H{
		"title":    title,
		"text":     text,
		"story_id": storyId,
	}
}

func (s Service) GetAll() []gin.H {
	rows, err := s.db.Query("SELECT title, text, story_id FROM scenes")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	all := make([]gin.H, 0)
	for rows.Next() {
		var title, text string
		var storyId int
		err = rows.Scan(&title, &text, &storyId)
		if err != nil {
			log.Fatal(err)
		}

		all = append(all, gin.H{
			"title":    title,
			"text":     text,
			"story_id": storyId,
		})
	}

	if err = rows.Err(); err != nil {
		log.Fatal(err)
	}

	return all
}

func (s Service) Update(id int, title, text string, storyId int) {
	_, err := s.db.Exec("UPDATE scenes SET title = $2, text = $3, story_id = $4 "+
		"WHERE id = $1", id, text, description, storyId)
	if err != nil {
		log.Fatal(err)
	}
}

func (s Service) Delete(id int) {
	_, err := s.db.Exec("DELETE FROM scenes WHERE id = $1", id)
	if err != nil {
		log.Fatal(err)
	}
}
