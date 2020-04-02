package story

import (
	"database/sql"
	"errors"
	"log"
	"strconv"

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

func (s Service) Create(title, description string, authorId int) error {
	_, err := s.db.Exec("INSERT INTO stories (title, description, author_id) "+
		"VALUES($1, $2, $3)", title, description, authorId)
	if err != nil {
		return err
	}

	return err
}

func (s Service) GetById(id int) (gin.H, error) {
	var title, description string
	var authorId int
	err := s.db.QueryRow("SELECT title, description, author_id FROM stories WHERE id = $1", id).
		Scan(&title, &description, &authorId)
	if err != nil {
		return nil, err
	}

	return gin.H{
		"title":       title,
		"description": description,
		"author_id":   authorId,
	}, nil
}

func (s Service) GetAll() ([]gin.H, error) {
	rows, err := s.db.Query("SELECT title, description, author_id FROM stories")
	if err != nil {
		return nil, err
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
		return nil, err
	}

	return all, nil
}

func (s Service) Update(id int, title, description string, authorId int) error {
	result, err := s.db.Exec("UPDATE stories SET title = $2, description = $3, author_id = $4 "+
		"WHERE id = $1", id, title, description, authorId)
	if err != nil {
		return err
	}

	rows, _ := result.RowsAffected()
	if rows == 0 {
		return errors.New("No story found with given id: " + strconv.Itoa(id))
	}

	return nil
}

func (s Service) Delete(id int) error {
	result, err := s.db.Exec("DELETE FROM stories WHERE id = $1", id)
	if err != nil {
		return err
	}

	rows, _ := result.RowsAffected()
	if rows == 0 {
		return errors.New("No story found with given id: " + strconv.Itoa(id))
	}

	return nil
}
