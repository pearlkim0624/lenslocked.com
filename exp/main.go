package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "Skan3147**"
	dbname   = "lenslocked_dev"
)

/*
type User struct {
	ID   int
	Name string
}
*/

func main() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	rows, err := db.Query(`SELECT * FROM users INNER JOIN orders ON users.id=orders.user_id`)
	if err != nil {
		panic(err)
	}

	for rows.Next() {
		var userID, orderID, amount int
		var email, name, desc string
		if err := rows.Scan(&userID, &name, &email, &orderID, &userID, &amount, &desc); err != nil {
			panic(err)
		}
		fmt.Println("userID:", userID, "name:", name, "email:", email, "orderID:", orderID, "userID:", userID, "amount:", amount, "desc:", desc)
	}
	if rows.Err() != nil {
		panic(rows.Err())
	}

	/*
		_, err = db.Exec(`
			INSERT INTO users(name, email)
			VALUES($1, $2)`, "Jon Calhoun", "jon@calhoun.io")
	*/

	/*
		var id int

			err = db.QueryRow(`
				INSERT INTO users(name, email)
				VALUES($1, $2) RETURNING id`, "Jon Calhoun", "jon@calhoun.io").Scan(&id)
	*/
	/*
		var name, email string

			row := db.QueryRow(`
				SELECT id, name, email
				FROM users
				WHERE id=$1`, 1)
			err = row.Scan(&id, &name, &email)
			if err != nil {
				if err == sql.ErrNoRows {
					fmt.Println("no rows")
				} else {
					panic(err)
				}
			}
	*/

	/*
		type User struct {
			ID    int
			Name  string
			Email string
		}

		var users []User
		rows, err := db.Query(`
			SELECT id, name, email
			FROM users`)

		if err != nil {
			panic(err)
		}
		defer rows.Close()

		for rows.Next() {
			var user User
			if err := rows.Scan(&user.ID, &user.Name, &user.Email); err != nil {
				panic(err)
			}
			users = append(users, user)
		}
		if rows.Err() != nil {
			panic(err)
		}
		fmt.Println(users)
	*/
}
