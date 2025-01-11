package models

type User struct {
	Id   int64  `gorm:"primaryKey;autoIncrement" json:"id"`
	Name string `gorm:"type:varchar(255)" json:"name"`
	Age  int    `json:"age"`
}
