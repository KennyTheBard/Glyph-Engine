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

func (s Service) Create(userId, storyId, sceneId int) error {
	result, err := s.db.Exec("INSERT INTO user_progress (user_id, story_id, scene_id) "+
		"VALUES($1, $2, $3)", userId, storyId, sceneId)
	if err != nil {
		return err
	}

	rows, _ := result.RowsAffected()
	if rows == 0 {
		return errors.New("User progress could not be saved")
	}

	return nil
}

func (s Service) GetByUserIdAndStoryId(userId, storyId int) (gin.H, error) {
	var sceneId int
	err := s.db.QueryRow("SELECT scene_id FROM user_progress WHERE "+
		"user_id = $1 and story_id = $2", userId, storyId).
		Scan(&sceneId)
	if err != nil {
		return nil, err
	}

	return gin.H{
		"user_id":  userId,
		"story_id": storyId,
		"scene_id": sceneId,
	}, nil
}

func (s Service) GetAllByUserId(userId int) ([]gin.H, error) {
	rows, err := s.db.Query("SELECT story_id, scene_id FROM user_progress WHERE "+
		"user_id = $1", userId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	all := make([]gin.H, 0)
	for rows.Next() {
		var storyId, sceneId int
		err = rows.Scan(&storyId, &sceneId)
		if err != nil {
			return nil, err
		}

		all = append(all, gin.H{
			"user_id":  userId,
			"story_id": storyId,
			"scene_id": sceneId,
		})
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return all, nil
}

func (s Service) GetAllByStoryId(storyId int) ([]gin.H, error) {
	rows, err := s.db.Query("SELECT user_id, scene_id FROM user_progress WHERE "+
		"story_id = $1", storyId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	all := make([]gin.H, 0)
	for rows.Next() {
		var userId, sceneId int
		err = rows.Scan(&userId, &sceneId)
		if err != nil {
			return nil, err
		}

		all = append(all, gin.H{
			"user_id":  userId,
			"story_id": storyId,
			"scene_id": sceneId,
		})
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return all, nil
}

func (s Service) GetAll() ([]gin.H, error) {
	rows, err := s.db.Query("SELECT user_id, story_id, scene_id FROM user_progress")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	all := make([]gin.H, 0)
	for rows.Next() {
		var userId, storyId, sceneId int
		err = rows.Scan(&userId, &storyId, &sceneId)
		if err != nil {
			return nil, err
		}

		all = append(all, gin.H{
			"user_id":  userId,
			"story_id": storyId,
			"scene_id": sceneId,
		})
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return all, nil
}

func (s Service) Update(userId, storyId, sceneId int) error {
	result, err := s.db.Exec("UPDATE user_progress SET scene_id = $3 WHERE "+
		"user_id = $1 AND story_id = $2", userId, storyId, sceneId)
	if err != nil {
		return err
	}

	rows, _ := result.RowsAffected()
	if rows == 0 {
		return errors.New("No progress found with given (user_id, story_id): (" +
			strconv.Itoa(userId) + ", " + strconv.Itoa(storyId) + ")")
	}

	return nil
}

func (s Service) Delete(userId, storyId int) error {
	result, err := s.db.Exec("DELETE FROM choices WHERE "+
		"user_id = $1 AND story_id = $2", userId, storyId)
	if err != nil {
		return err
	}

	rows, _ := result.RowsAffected()
	if rows == 0 {
		return errors.New("No progress found with given (user_id, story_id): (" +
			strconv.Itoa(userId) + ", " + strconv.Itoa(storyId) + ")")
	}

	return nil
}
