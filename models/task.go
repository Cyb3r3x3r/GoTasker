package models

import "time"

type Task struct {
	ID      int
	Title   string
	Status  string
	Duetime time.Time
	UserID  int
}
