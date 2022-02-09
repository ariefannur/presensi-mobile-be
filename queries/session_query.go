package queries

import (
	"fmt"
	"presensi-mobile/models"

	"github.com/jmoiron/sqlx"
)

type SessionQueries struct {
	*sqlx.DB
}

func (q *SessionQueries) CreateSession(userId string, token string, device string) error {
	getQuery := `SELECT * FROM "Sessions" WHERE user_id = $1`
	session := models.Session{}
	errGet := q.Get(&session, getQuery, userId)

	if errGet != nil {
		fmt.Println("insert session")
		query := `INSERT INTO "Sessions" (user_id, token, device, status) VALUES ($1, $2, $3, 'ONLINE')`
		_, err := q.Exec(query, userId, token, device)
		if err != nil {
			return err
		}
	} else {
		fmt.Println("update session")
		query := `UPDATE "Sessions" SET token = $1, device = $2, status = 'ONLINE' WHERE user_id = $3`
		_, err := q.Exec(query, token, device, userId)
		if err != nil {
			return err
		}
	}

	return nil
}

func (q *SessionQueries) Logout(userId string) error {
	query := `UPDATE "Sessions" SET status = 'OFFLINE', token = '' WHERE user_id = $1`
	_, err := q.Exec(query, userId)
	if err != nil {
		return err
	}
	return nil
}

func (q *SessionQueries) GetSession(userId string) (models.Session, error) {
	query := `SELECT * FROM "Sessions" WHERE user_id = $1`

	session := models.Session{}
	err := q.QueryRow(query, userId).Scan(&session.ID, &session.User_ID, &session.Token, &session.Time, &session.Device, &session.Status)
	fmt.Println(err)
	if err != nil {
		return session, err
	}
	return session, nil
}

func (q *SessionQueries) RefreshToken(userId string, token string) error {
	query := `UPDATE "Sessions" SET token = $1 WHERE user_id = $2`
	_, err := q.Exec(query, token, userId)
	if err != nil {
		return err
	}
	return nil
}
