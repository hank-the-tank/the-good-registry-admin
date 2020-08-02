# Start
1. Run the app
```
$ go run *.go
```

## DB - Version
1. Update db version
```  
migrate -source file://migrations -database "postgres://hank:password@localhost:5432/the_good_registry?sslmode=disable" up 1   
```
2. Downgrade db version
```
migrate -source file://migrations -database "postgres://hank:password@localhost:5432/the_good_registry?sslmode=disable" down 1 
```


# Reference
1. [gin](https://github.com/gin-gonic/gin)
2. [golang-migrate](https://github.com/golang-migrate/migrate)
3. [psql](https://www.postgresql.org/)