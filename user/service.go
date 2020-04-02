package user

import (
	"database/sql"
	"errors"
	"log"

	"glyph/security"

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

func (s Service) Register(username, password string) error {
	hashedPass, err := security.HashPassword(password)
	if err != nil {
		log.Fatal(err)
	}

	_, err = s.db.Exec("INSERT INTO users (username, password_hash) VALUES($1, $2)",
		username, hashedPass)

	if err != nil {
		return err
	}

	return nil
}

func (s Service) Login(username, password string) error {
	var id int
	var hashedPassword string
	err := s.db.QueryRow("SELECT id, password_hash FROM users WHERE username = $1", username).
		Scan(&id, &hashedPassword)

	if err != nil || !security.CheckPasswordHash(password, hashedPassword) {
		return errors.New("Username or password incorrect")
	}

	return nil
}

func (s Service) GetById(id int) gin.H {
	var username string
	err := s.db.QueryRow("SELECT username FROM users WHERE id = $1", id).
		Scan(&username)
	if err != nil {
		log.Fatal(err)
	}

	return map[string]interface{}{
		"username": username,
	}
}

func (s Service) GetAll() []gin.H {
	rows, err := s.db.Query("SELECT username FROM users")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	all := make([]gin.H, 0)
	for rows.Next() {
		var username string
		err = rows.Scan(&username)
		if err != nil {
			log.Fatal(err)
		}

		all = append(all, map[string]interface{}{
			"username": username,
		})
	}

	if err = rows.Err(); err != nil {
		log.Fatal(err)
	}

	return all
}

func (s Service) Update(id int, username string) {
	_, err := s.db.Exec("UPDATE users SET username = $2 last_name WHERE id = $1",
		id, username)
	if err != nil {
		log.Fatal(err)
	}
}

func (s Service) Delete(id int) {
	_, err := s.db.Exec("DELETE FROM users WHERE id = $1", id)
	if err != nil {
		log.Fatal(err)
	}
}
