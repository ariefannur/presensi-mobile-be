package queries

import (
	"fmt"
	"presensi-mobile/models"

	"github.com/jmoiron/sqlx"
)

type PresensiQueries struct {
	*sqlx.DB
}

func (p *PresensiQueries) CreatePresensi(pp *models.Presensi) error {
	query := `INSERT INTO "Presensi" (user_id, foto, lat, lng, alamat) VALUES($1, $2, $3, $4, $5)`
	_, err := p.Exec(query, pp.User_Id, pp.Foto, pp.Lat, pp.Lng, pp.Alamat)
	if err != nil {
		return err
	}
	return nil
}

func (p *PresensiQueries) GetPresensi(userId string) ([]models.Presensi, error) {
	query := `SELECT * FROM "Presensi" WHERE user_id = $1`
	presensi := []models.Presensi{}
	err := p.Select(&presensi, query, userId)
	if err != nil {
		return presensi, err
	}
	return presensi, nil
}

func (p *PresensiQueries) CheckPrensensiToday(userId string) (models.Presensi, error) {
	query := `SELECT * FROM "Presensi" WHERE time::date = current_date AND user_id = $1 LIMIT 1`
	presensi := models.Presensi{}
	err := p.QueryRow(query, userId).Scan(&presensi.ID, &presensi.User_Id, &presensi.Foto, &presensi.Time, &presensi.Lat, &presensi.Lng, &presensi.Alamat)
	if err != nil {
		return presensi, err
	}
	fmt.Println(presensi)
	return presensi, nil
}
