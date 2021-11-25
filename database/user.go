package database

import (
	"fmt"
	"log"
)

type User struct {
	ID int64 `json:"id" gorm:"column:id;primary_key;AUTO_INCREMENT"`
	UserName string `json:"username" gorm:"column:username"`
	PassWord string `json:"pass_word" gorm:"column:password"`
}

func (u User) TableName() string {
	return "user"
}

func (u *User) SelectAll() []*User {
	d := DB()

	res := make([]*User, 0)

	result := d.Find(&res)

	fmt.Println(result.RowsAffected)
	log.Println(res)
	return res
}