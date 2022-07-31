package main

import (
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var schema = `
	CREATE TABLE IF NOT EXISTS pets (
		id INT(6) UNSIGNED AUTO_INCREMENT PRIMARY KEY,
		name VARCHAR(255),
		age INT(6) UNSIGNED,
		gender ENUM('male', 'female'),
		can_fly BOOLEAN,
		created_at DATETIME,
		updated_at TIMESTAMP
	)
`

type Pet struct {
	ID        int     `db:"id"`
	Name      string  `db:"name"`
	Age       int     `db:"age"`
	Gender    string  `db:"gender"`
	CanFly    bool    `db:"can_fly"`
	CreatedAt *string `db:"created_at"`
	UpdateAt  *string `db:"updated_at"`
}

func main() {
	config := fmt.Sprintf(
		"%s:%s@tcp(%s)/%s?parseTime=True",
		"root", "root", "127.0.0.2", "animals",
	)
	db, err := sqlx.Connect("mysql", config)
	if err != nil {
		log.Fatalln(err)
	}

	// exec the schema or fail; multi-statement Exec behavior varies between
	// database drivers;  pq will exec them all, sqlite3 won't, ymmv
	db.MustExec(schema)

	// tx := db.MustBegin()
	// tx.MustExec("INSERT INTO pets (name, age, gender, can_fly) VALUES (?, ?, ?, ?)", "Cat", 5, "male", false)
	// tx.MustExec("INSERT INTO pets (name, age, gender, can_fly) VALUES (?, ?, ?, ?)", "Bird", 2, "female", true)
	// tx.MustExec("INSERT INTO pets (name, age, gender, can_fly) VALUES (?, ?, ?, ?)", "Dog", 3, "male", false)
	// tx.MustExec("INSERT INTO pets (name, age, gender, can_fly) VALUES (?, ?, ?, ?)", "Something", 50, "female", true)
	// tx.Commit()
	animals := []Pet{}
	err = db.Select(&animals, "SELECT * FROM pets")
	if err != nil {
		log.Println(err)
	}
	for _, v := range animals {
		log.Println(v)
	}
	fmt.Print("hello")
}
