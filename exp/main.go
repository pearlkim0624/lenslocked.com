package main

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "Skan3147**"
	dbname   = "lenslocked_dev"
)

/*
type Model struct {
	ID        uint `gorm:"primary_key"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:"index"`
}
*/

type User struct {
	gorm.Model // embeded
	Name       string
	Email      string `gorm:"not null;unique_index"`
	Color      string
}

func main() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err := gorm.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()
	db.LogMode(true)
	db.AutoMigrate(&User{})

	var u User
	/*
		newDB := db.Where("id = ?", 3).Where("color = ?", "yellow").First(&u)
		if newDB.Error != nil {
			panic(newDB.Error)
		}
	*/
	if err := db.Where("id = ?", 3).Where("color = ?", "yellow").First(&u).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			fmt.Println("No user found!")
		default:
			panic(err)
		}
	}
	fmt.Println(u)
}
