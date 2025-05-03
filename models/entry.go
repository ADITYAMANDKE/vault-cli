package models

import (
	"time"
)

type Entry struct {
	Site       string    `json:"site"`
	Username   string    `json:"username"`
	Password   string    `json:"password"`
	ModifiedAt time.Time `json:"modified_at"`
	Note       string    `json:"note"`
}
