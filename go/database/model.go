package model

import (
	"time"
)

// User struct for logged users
type User struct {
	UID       int
	UserName  string
	LastLogin time.Time
}
