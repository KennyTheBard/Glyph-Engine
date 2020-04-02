package story

import (
	"database/sql"
	"errors"
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

func (s Service) Create(title, text string, sceneId, nextSceneId int) error {
	result, err := s.db.Exec("INSERT INTO choices (title, text, scene_id, next_scene_id) "+
		"VALUES($1, $2, $3, $4)", title, text, sceneId, nextSceneId)
	if err != nil {
		return err
	}

	rows, _ := result.RowsAffected()
	if rows == 0 {
		return errors.New("Choice could not be saved")
	}

	return nil
}

func (s Service) GetById(id int) (gin.H, error) {
	var title, text string
	var sceneId, nextSceneId int
	err := s.db.QueryRow("SELECT title, text, scene_id, next_scene_id FROM choices WHERE id = $1", id).
		Scan(&title, &text, &sceneId, &nextSceneId)
	if err != nil {
		return nil, err
	}

	return gin.H{
		"title":         title,
		"text":          text,
		"scene_id":      sceneId,
		"next_scene_id": nextSceneId,
	}, nil
}

func (s Service) GetAll() ([]gin.H, error) {
	rows, err := s.db.Query("SELECT id, title, text, scene_id, next_scene_id FROM choices")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	all := make([]gin.H, 0)
	for rows.Next() {
		var title, text string
		var id, sceneId, nextSceneId int
		err = rows.Scan(&id, &title, &text, &sceneId, &nextSceneId)
		if err != nil {
			return nil, err
		}

		all = append(all, gin.H{
			"id":            id,
			"title":         title,
			"text":          text,
			"scene_id":      sceneId,
			"next_scene_id": nextSceneId,
		})
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return all, nil
}

func (s Service) Update(id int, title, text string, sceneId, nextSceneId int) error {
	result, err := s.db.Exec("UPDATE choices SET title = $2, text = $3, story_id = $4, next_scene_id = $5 "+
		"WHERE id = $1", id, title, text, sceneId, nextSceneId)
	if err != nil {
		return err
	}

	rows, _ := result.RowsAffected()
	if rows == 0 {
		return errors.New("No choice found with given id: " + strconv.Itoa(id))
	}

	return nil
}

func (s Service) Delete(id int) error {
	result, err := s.db.Exec("DELETE FROM choices WHERE id = $1", id)
	if err != nil {
		return err
	}

	rows, _ := result.RowsAffected()
	if rows == 0 {
		return errors.New("No choice found with given id: " + strconv.Itoa(id))
	}

	return nil
}
