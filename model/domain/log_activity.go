package domain

import "time"

type LogActivity struct {
	Id      int `json:"id"`
	AdminId string `json:"admin_id"`
	Message string `json:"message"`
	Time    time.Time `json:"time"`
}
