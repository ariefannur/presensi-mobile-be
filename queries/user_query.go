package queries

import (
	"fmt"
	"presensi-mobile/models"

	"github.com/jmoiron/sqlx"
)

type UserQueries struct {
	*sqlx.DB
}

func (q *UserQueries) InsertUsers(user *models.User) error {
	query := `INSERT INTO users VALUES ($1, $2, $3, $4, $5)`
	fmt.Println("SUDAH DISINI")
	_, err := q.Exec(query, user.Id, user.Name, user.Email, user.Password, user.UserType)

	if err != nil {
		return err
	}

	return nil
}
