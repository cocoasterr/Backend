package models

type User struct{
	Id int64 `gorm:"primaryKey" json:"id"`
	Username string `gorm:"varchar(50)" json:"username"`
	Email string `gorm:"varchar(50)" json:"email"`
	Password string `gorm:"varchar(50)" json:"password"`
}