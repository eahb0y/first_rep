package main

import (
	_ "github.com/go-sql-driver/mysql"
)

//database

var customer_plan = `
	CREATE TABLE IF NOT EXISTS customer_plan(
		id INT(2) UNSIGNED AUTO_INCREMENT PRIMARY KEY,
		customers_id INT(2),
		plan_id INT(2)
	)
`

var tables = `
	CREATE TABLE IF NOT EXISTS tables(
		id INT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
		number INT(2) UNSIGNED AUTO_INCREMENT PRIMARY KEY,
		table_type_id INT(2),
		room_id INT(2)
	)
`
var customer_table = `
	CREATE TABLE IF NOT EXISTS customer_table(
		customer_id INT(2),
		table_id INT(2)
	)
`

var table_types = `
	CREATE TABLE IF NOT EXISTS table_types(
		id INT(2) UNSIGNED AUTO_INCREMENT PRIMARY KEY,
		type VARCHAR(255)
	)
`

var customer_room = `
	CREATE TABLE IF NOT EXISTS customer_room(
		customer_id INT(2),
		room_id INT(2)
	)
`

var rooms = `
	CREATE TABLE IF NOT EXISTS rooms(
		id INT(2),
		name VARCHAR(255),
		working_pleace_count INT(2)
	)
`

var plans = `
	CREATE TABLE IF NOT EXISTS plans(
		id INT(2) UNSIGNED AUTO_INCREMENT PRIMARY KEY,
		name VARCHAR(255),
		price INT(6)
	)
`

var customers = `
CREATE TABLE IF NOT EXISTS customers(
	id INT(2) UNSIGNED AUTO_INCREMENT PRIMARY KEY,
	first_name VARCHAR(255),
	last_name VARCHAR(255),
	passport_id INT(6)
)
`

//database struct
type pets struct {
	ID        int     `db:"id" json:"id"`
	Name      string  `db:"name" json:"Name" query:"Name"`
	Age       int     `db:"age" json:"age" query:"age"`
	Gender    string  `db:"gender" json:"gender" query:"gender"`
	CanFly    bool    `db:"can_fly" json:"can_fly" query:"can_fly"`
	CreatedAt *string `db:"created_at" json:"created_at"`
	UpdateAt  *string `db:"updated_at" json:"updated_at"`
	Pet_Id    int     `db:"pet_id" json:"pet_id" query:"pet_id"`
}

type pets_type struct {
	ID        int    `db:"id" json:"id"`
	Type      string `db:"type" json:"type" query:"type"`
	Column_id int    `db:"column_id" json:"column_id" query:"column_id"`
}

//main function
func main() {
	//declare databese

	// config := fmt.Sprintf(
	// 	"%s:%s@tcp(%s)/%s?parseTime=True",
	// 	"root", "root", "127.0.0.2", "animals",
	// )
	// db, err := sqlx.Connect("mysql", config)
	// if err != nil {
	// 	log.Fatalln(err)
	// }

	//using fiber for API

	// app := fiber.New()
	// app.Get("/", func(c *fiber.Ctx) error {
	// 	return c.SendString("Hello, World!")
	// })
	// app.Get("animals", func(c *fiber.Ctx) error {
	// 	animals := []Pet{}
	// 	err = db.Select(&animals, "SELECT * FROM pets")
	// 	if err != nil {
	// 		log.Println(err)
	// 	}
	// 	return c.JSON(animals)
	// })
	// app.Get("animals/add", func(ctx *fiber.Ctx) error {
	// 	pet := Pet{}
	// 	ctx.QueryParser(&pet)
	// 	tx := db.MustBegin()
	// 	tx.MustExec("INSERT INTO pets (name, age, gender, can_fly) VALUES (?, ?, ?, ?)", pet.Name, pet.Age, pet.Gender, pet.CanFly)
	// 	tx.Commit()
	// 	return ctx.JSON(pet)
	// })
	// app.Listen(":3000")

	//enter to db values
	// exec the schema or fail; multi-statement Exec behavior varies between
	// // database drivers;  pq will exec them all, sqlite3 won't, ymmv
	// db.MustExec()
	// db.MustExec(customers)
	// tx := db.MustBegin()
	// db.MustExec("INSERT INTO customers(first_name, last_name, passport_id) VALUES (?, ?, ?)", "Robert", "Downey", 21038723)
	// db.MustExec("INSERT INTO customers(first_name, last_name, passport_id) VALUES (?, ?, ?)", "Robert", "Downey", 21038723)
	// tx.Commit()

}
