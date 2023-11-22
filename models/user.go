package models

import "time"

// User
type User struct {
	Id        string    `json:"id"`
	Name      string    `gorm:"not null" json:"name"`
	Surname   string    `gorm:"not null" json:"surname"`
	About     string    `json:"about"`
	Email     string    `gorm:"unique" json:"email"`
	Mobile    string    `gorm:"unique" json:"mobile"`
	Password  []byte    `gorm:"not null" json:"-"`
	Image     string    `json:"image"`
	CreatedAt time.Time `json:"created_at"`
	DeleteAt  time.Time `json:"delete_at"`
}

// Save valids tokens
type ValidToken struct {
	Token string `db:"token"`
}
