package models

type Book struct {
	ID int64 `gorm:"primary_key" json:"id"`
	BookName string `gorm:"type:varchar(300)" json:"book_name"`
	Author string `gorm:"type:varchar(100)" json:"author"`
}