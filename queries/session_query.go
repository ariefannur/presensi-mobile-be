package queries

import (
	"presensi-mobile/models"

	"github.com/jmoiron/sqlx"
)

type SessionQueries struct {
	*sqlx.DB
}

func (q *SessionQueries) CreateSession(userId string, token string, device string) error {
	// get first
	// if any data update data
	query := `INSERT INTO "Sessions" (user_id, token, device, status) VALUES ($1, $2, $3, 'ONLINE)`
	_, err := q.Exec(query, userId, token, device)

	if err != nil {
		return err
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
	err := q.Get(&session, query, userId)
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
