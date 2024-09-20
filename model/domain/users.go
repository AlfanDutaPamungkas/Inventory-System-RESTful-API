package domain

import "time"

type Users struct {
	Id, Name, Email, Password, Role, Status string
	CreatedAt, UpdatedAt            time.Time
}
