package models

// User
type App struct {
	Id    string `json:"id"`
	Name  string `gorm:"unique" json:"name"`
	Url   string `gorm:"unique" json:"url"`
	Users []User `gorm:"many2many:user_app" json:"users"`
}
