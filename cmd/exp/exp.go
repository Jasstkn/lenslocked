package main

// TODO: update to v5
import (
	"database/sql"
	"fmt"
	"github.com/Jasstkn/lenslocked/models"

	_ "github.com/jackc/pgx/v4/stdlib"
)

func main() {
	cfg := models.DefaultPostgresConfig()
	db, err := sql.Open("pgx", cfg.String())
	defer db.Close()

	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}
	fmt.Println("Connection to DB is established!")

	us := models.UserService{DB: db}
	user, err := us.Create("bob1@bob.com", "bob123")
	if err != nil {
		panic(err)
	}
	fmt.Println(user)
	// Insert some data...
	// name := "Jon Calhoun"
	// email := "new@calhoun.io"
	// row := db.QueryRow(`
	// 	INSERT INTO users (name, email)
	// 	VALUES ($1, $2) RETURNING id;`, name, email)

	// var id int
	// err = row.Scan(&id)
	// if err != nil {
	// 	panic(err)
	// }

	// fmt.Println("User created. id = ", id)

	// id := 1

	// row := db.QueryRow(`
	// 	SELECT name, email
	// 	FROM users
	// 	WHERE id=$1`, id)

	// var name, email string
	// err = row.Scan(&name, &email)
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Printf("User information: name=%s, email=%s\n", name, email)

	// userID := 1
	// for i := 1; i < 5; i++ {
	// 	amount := i * 100
	// 	desc := fmt.Sprintf("Fake order #%d", i)
	// 	_, err := db.Exec(`
	// 		INSERT INTO orders(user_id, amount, description)
	// 		VALUES($1, $2, $3)`, userID, amount, desc)
	// 	if err != nil {
	// 		panic(err)
	// 	}
	// }
	// fmt.Println("Created fake orders.")

	//type Order struct {
	//	ID          int
	//	UserID      int
	//	Amount      int
	//	Description string
	//}
	//
	//var orders []Order
	//
	//userID := 1
	//rows, err := db.Query(`
	//	SELECT id, amount, description
	//	FROM orders
	//	WHERE user_id=$1`, userID)
	//if err != nil {
	//	panic(err)
	//}
	//defer rows.Close()
	//
	//for rows.Next() {
	//	var order Order
	//	order.UserID = userID
	//	err := rows.Scan(&order.ID, &order.Amount, &order.Description)
	//	if err != nil {
	//		panic(err)
	//	}
	//	orders = append(orders, order)
	//}
	//err = rows.Err()
	//if err != nil {
	//	panic(err)
	//}
	//
	//fmt.Println("Orders: ", orders)
}
