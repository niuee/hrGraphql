package dbHandle

import (
	"database/sql"
	"fmt"
	"os"
	"sync"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

var dbConn *sql.DB
var lock = &sync.Mutex{}
var err = godotenv.Load()

func GetDBConn() *sql.DB {
	if dbConn == nil {
		lock.Lock()
		defer lock.Unlock()
		if dbConn == nil {
			if err != nil {
				fmt.Println("Can't Load Environment")
				panic(err)
			}
			host := os.Getenv("DB_HOST")
			port := os.Getenv("DB_PORT")
			user := os.Getenv("DB_USER")
			password := os.Getenv("DB_PASSWORD")
			dbname := os.Getenv("DB_NAME")
			psqlInfo := fmt.Sprintf("host=%s port=%s user=%s "+"password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
			var err error
			dbConn, err = sql.Open("postgres", psqlInfo)
			if err != nil {
				panic(err)
			}
		}
	}

	return dbConn
}
