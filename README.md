# About
This is a project that aims to build an admin for [The Good Registry](https://thegoodregistry.com/) to manage their orders efficiently. 

### How it works
1. Fetch email from Gmail to get the gift codes 
2. List all the orders and manage the orders
3. A tool to generate gift cards for the customer who prefers to print the gift cards as a gift to the receiver

# Start
1. Run the app
```
$ go run *.go
```

## DB - Version
1. Update db version
```  
$ migrate -source file://migrations -database "postgres://hank:password@localhost:5432/the_good_registry?sslmode=disable" up 1   
```
2. Downgrade db version
```
$ migrate -source file://migrations -database "postgres://hank:password@localhost:5432/the_good_registry?sslmode=disable" down 1 
```


# Reference
1. [gin](https://github.com/gin-gonic/gin)
2. [golang-migrate](https://github.com/golang-migrate/migrate)
3. [psql](https://www.postgresql.org/)