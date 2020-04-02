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

func (s Service) Create(title, text string, storyId int) error {
	result, err := s.db.Exec("INSERT INTO scenes (title, text, story_id) "+
		"VALUES($1, $2, $3)", title, text, storyId)
	if err != nil {
		return err
	}

	rows, _ := result.RowsAffected()
	if rows == 0 {
		return errors.New("Scene could not be saved")
	}

	return nil
}

func (s Service) GetById(id int) (gin.H, error) {
	var title, text string
	var storyId int
	err := s.db.QueryRow("SELECT title, text, story_id FROM scenes WHERE id = $1", id).
		Scan(&title, &text, &storyId)
	if err != nil {
		return nil, err
	}

	return gin.H{
		"title":    title,
		"text":     text,
		"story_id": storyId,
	}, nil
}

func (s Service) GetAll() ([]gin.H, error) {
	rows, err := s.db.Query("SELECT id, title, text, story_id FROM scenes")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	all := make([]gin.H, 0)
	for rows.Next() {
		var title, text string
		var id, storyId int
		err = rows.Scan(&id, &title, &text, &storyId)
		if err != nil {
			return nil, err
		}

		all = append(all, gin.H{
			"id":       id,
			"title":    title,
			"text":     text,
			"story_id": storyId,
		})
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return all, nil
}

func (s Service) Update(id int, title, text string) error {
	result, err := s.db.Exec("UPDATE scenes SET title = $2, text = $3 "+
		"WHERE id = $1", id, title, text)
	if err != nil {
		return err
	}

	rows, _ := result.RowsAffected()
	if rows == 0 {
		return errors.New("No scene found with given id: " + strconv.Itoa(id))
	}

	return nil
}

func (s Service) Delete(id int) error {
	result, err := s.db.Exec("DELETE FROM scenes WHERE id = $1", id)
	if err != nil {
		return err
	}

	rows, _ := result.RowsAffected()
	if rows == 0 {
		return errors.New("No scene found with given id: " + strconv.Itoa(id))
	}

	return nil
}
