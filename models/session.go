package models

import (
	"time"
)

type Session struct {
	ID      int64     `json:"id"`
	User_ID int64     `json:"user_id"`
	Token   string    `json:"token"`
	Time    time.Time `json:"time"`
	Device  string    `json:"device"`
	Status  string    `json:"status"`
}
