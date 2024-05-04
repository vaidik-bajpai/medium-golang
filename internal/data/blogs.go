package data

import (
	"database/sql"
	"time"
)

type Blog struct {
	ID        int64
	CreatedAt time.Time
	Title     string
	Content   string
	Readtime  int32
	AuthorID  int64
	Version   int32
}

type BlogModel struct {
	DB *sql.DB
}
