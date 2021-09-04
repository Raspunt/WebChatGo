package database

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// type Users struct {
// 	ID   uint
// 	NAME string
// }

type Messages struct {
	ID       uint
	USERNAME string
	TEXT     string
}

var db, err = gorm.Open(postgres.Open(ConString), &gorm.Config{})

func CreateUserTable() {

	if err != nil {
		panic(err)
	}

	db.Migrator().CreateTable(&Messages{})

}

func CreateMessage(id uint, hostname string, text string) {

	if err != nil {
		panic(err)
	}

	message := Messages{id, hostname, text}

	fmt.Println(message)

	db.AutoMigrate(&Messages{})
	db.Create(&message)

}

func UserIsExists(name string) {

	mathes := db.Find(&User, name)

	fmt.Println(mathes)
}

func GetAllMessages() []Messages {
	var res []Messages

	db.Find(&res)
	fmt.Println(res)

	return res

}

func TableExists(inter interface{}) bool {
	fmt.Println(db.Migrator().HasTable(&inter))
	return db.Migrator().HasTable(&inter)
}
