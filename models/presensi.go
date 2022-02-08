package models

import "time"

type Presensi struct {
	ID      int64     `json:"id"`
	User_Id int64     `json:"user_id"`
	Foto    string    `json:"foto"`
	Time    time.Time `json:"time"`
	Lat     string    `json:"lat"`
	Lng     string    `json:"lng"`
	Alamat  string    `json:"alamat"`
}
