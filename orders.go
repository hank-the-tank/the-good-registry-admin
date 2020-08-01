package main

import (
	"database/sql"
	"log"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type order struct {
	ID      int       `json: "id"`
	Name    string    `json:"customer_name"`
	Volume  int       `json:"vloume"`
	Email   string    `json:"email"`
	Created time.Time `json:"created"`
}

var connStr = "postgres://hank:password@localhost:5432/the_good_registry?sslmode=disable"
var id int
var customerName string
var email string
var volume int
var created time.Time

var orders []order

func getOrders(c *gin.Context) {
	db, err := sql.Open("postgres", connStr)
	rows, err := db.Query(`SELECT id, customer_name, email, volume, created  FROM orders ORDER BY created`)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&id, &customerName, &email, &volume, &created)
		if err != nil {
			log.Fatal(err)
		}
		row := order{ID: id, Name: customerName, Email: email, Volume: volume, Created: created}
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
