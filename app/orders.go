package app

import (
	"database/sql"
	"log"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type order struct {
	ID            int       `json: "id"`
	Volume        int       `json: "vloume"`
	Amount        int       `json: "amount"`
	Created       time.Time `json: "created"`
	CustomerEmail string    `json: "customer_email"`
	CustomerName  string    `json: "customer_name"`
}

var connStr = "postgres://hank:password@localhost:5432/the_good_registry?sslmode=disable"
var id int
var amount int
var volume int
var created time.Time
var email string
var customerName string

var orders []order

// GetOrders is a function for an endpoint
func GetOrders(c *gin.Context) {
	db, err := sql.Open("postgres", connStr)
	rows, err := db.Query(`SELECT o.id, o.amount, o.volume, o.created, c.email, c.preferred_name  FROM orders o JOIN customers c ON c.id = o.customer_id ORDER BY created;`)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&id, &amount, &volume, &created, &email, &customerName)
		if err != nil {
			log.Fatal(err)
		}
		row := order{ID: id, Amount: amount, Volume: volume, Created: created, CustomerName: customerName, CustomerEmail: email}
		c.Bind(&row)
		c.JSON(200, gin.H{
			strconv.Itoa(id): row,
		})
		orders = append(orders, row)
	}
	// c.JSON(200, gin.H{
	// 	"order":   "Hank",
	// 	"order_1": "Hank",
	// 	"order_2": "Hank",
	// 	"order_3": "Hank",
	// 	"order_4": "Hank",
	// })
}
