package queries

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"presensi-mobile/models"

	"github.com/jmoiron/sqlx"
)

type UserQueries struct {
	*sqlx.DB
}

func GetMD5Hash(text string) string {
	hasher := md5.New()
	hasher.Write([]byte(text))
	return hex.EncodeToString(hasher.Sum(nil))
}

func (q *UserQueries) InsertUsers(user *models.Users) error {
	query := `INSERT INTO "Users" VALUES ($1, $2, $3, $4, $5)`
	pwd := GetMD5Hash(user.Password)
	_, err := q.Exec(query, user.Id, user.Name, user.Email, pwd, user.User_Type)

	if err != nil {
		return err
	}

	return nil
}

func (q *UserQueries) GetUsers() ([]models.Users, error) {
	query := `SELECT * FROM "Users" `
	users := []models.Users{}
	err := q.Select(&users, query)
	fmt.Println(err)
	if err != nil {
		return users, err
	}
	return users, nil
}

func (q *UserQueries) GetUsersById(id string) (models.Users, error) {
	query := `SELECT * FROM "Users" WHERE id = $1`
	users := models.Users{}
	err := q.Get(&users, query, id)
	fmt.Println(err)
	if err != nil {
		return users, err
	}
	return users, nil
}

func (q *UserQueries) Login(email string, password string) (models.Users, error) {
	query := `SELECT * FROM "Users" WHERE email = $1 AND password = $2 LIMIT 1`
	user := models.Users{}
	err := q.QueryRow(query, email, GetMD5Hash(password)).Scan(&user.Id, &user.Name, &user.Email, &user.Password, &user.User_Type)

	if err != nil {
		return user, err
	}

	return user, nil
}

func (q *UserQueries) ChangePassword(newPassword string, email string) error {
	query := `UPDATE "Users" SET password = $1 WHERE email = $2`
	_, err := q.Exec(query, GetMD5Hash(newPassword), email)
	fmt.Println(err)
	if err != nil {
		return err
	}

	return nil
}
