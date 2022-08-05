package main

import (
	"first_rep/greet"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
)

// var schema = `
// 	CREATE TABLE IF NOT EXISTS pets (
// 		id INT(6) UNSIGNED AUTO_INCREMENT PRIMARY KEY,
// 		name VARCHAR(255),
// 		age INT(6) UNSIGNED,
// 		gender ENUM('male', 'female'),
// 		can_fly BOOLEAN,
// 		created_at DATETIME,
// 		updated_at TIMESTAMP
// 	)
// `

type Pet struct {
	ID        int     `db:"id" json:"id"`
	Name      string  `db:"name" json:"name" query:"name"`
	Age       int     `db:"age" json:"age" query:"age"`
	Gender    string  `db:"gender" json:"gender" query:"gender"`
	CanFly    bool    `db:"can_fly" json:"can_fly" query:"can_fly"`
	CreatedAt *string `db:"created_at" json:"created_at"`
	UpdateAt  *string `db:"updated_at" json:"updated_at"`
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

	app := fiber.New()
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})
	app.Get("animals", func(c *fiber.Ctx) error {
		animals := []Pet{}
		err = db.Select(&animals, "SELECT * FROM pets")
		if err != nil {
			log.Println(err)
		}
		return c.JSON(animals)
	})
	app.Get("animals/add", func(ctx *fiber.Ctx) error {
		pet := Pet{}
		ctx.QueryParser(&pet)
		tx := db.MustBegin()
		tx.MustExec("INSERT INTO pets (name, age, gender, can_fly) VALUES (?, ?, ?, ?)", pet.Name, pet.Age, pet.Gender, pet.CanFly)
		tx.Commit()
		return ctx.JSON(pet)
	})
	app.Listen(":3000")
	// // exec the schema or fail; multi-statement Exec behavior varies between
	// // database drivers;  pq will exec them all, sqlite3 won't, ymmv
	// db.MustExec(schema)

	// // tx := db.MustBegin()
	// // tx.MustExec("INSERT INTO pets (name, age, gender, can_fly) VALUES (?, ?, ?, ?)", "Cat", 5, "male", false)
	// // tx.MustExec("INSERT INTO pets (name, age, gender, can_fly) VALUES (?, ?, ?, ?)", "Bird", 2, "female", true)
	// // tx.MustExec("INSERT INTO pets (name, age, gender, can_fly) VALUES (?, ?, ?, ?)", "Dog", 3, "male", false)
	// // tx.MustExec("INSERT INTO pets (name, age, gender, can_fly) VALUES (?, ?, ?, ?)", "Something", 50, "female", true)
	// // tx.Commit()

	// for _, v := range animals {
	// 	log.Println(v)
	// }
	// fmt.Print("he")
	greet.PrintCircleArea(6)
	greet.CalculateCricleArea(5)

}
