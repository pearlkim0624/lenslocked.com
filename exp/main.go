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
	Orders     []Order
}

type Order struct {
	gorm.Model
	UserID      uint
	Amount      int
	Description string
}

func main() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err := gorm.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()
	db.LogMode(true)
	db.AutoMigrate(&User{}, &Order{})

	var users []User
	if err := db.Preload("Orders").Find(&users).Error; err != nil {
		panic(err)
	}
	fmt.Println(users)

	/* Users & Orders
	var u User
	if err := db.Preload("Orders").First(&u).Error; err != nil {
		panic(err)
	}
	fmt.Println(u)
	fmt.Println(u.Orders)
	*/

	/* Create Order
	var u User
	if err := db.First(&u).Error; err != nil {
		panic(err)
	}
	createOrder(db, u, 1001, "Fake Description #1")
	createOrder(db, u, 9999, "Fake Description #2")
	createOrder(db, u, 100, "Fake Description #3")
	*/

	//fmt.Println(u)
}

/*
func createOrder(db *gorm.DB, user User, amount int, desc string) {
	err := db.Create(&Order{
		UserID:      user.ID,
		Amount:      amount,
		Description: desc,
	}).Error
	if err != nil {
		panic(err)
	}
}
*/
