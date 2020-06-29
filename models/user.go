package models

import "gin/db"

type User struct {
	ID   string
	Name string
}

func init() {
	// fmt.Println(123123)
	// db := db.GetDB()
	// fmt.Println(db, 1)
	// if !db.HasTable(&User{}) {
	// 	db.CreateTable(&User{})
	// }
}

func (u User) Get() *User {
	db := db.GetDB()
	if !db.HasTable(&User{}) {
		db.CreateTable(&User{})
	}
	// fmt.Println(db, 2)
	user := new(User)

	// db.First(user)
	return user
}