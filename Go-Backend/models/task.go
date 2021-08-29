package models

type Task struct {
	// gorm.Model
	ID     uint
	Title  string
	Body   string
	UserId uint
	User   User `gorm:"foreignKey:UserId"`
}
