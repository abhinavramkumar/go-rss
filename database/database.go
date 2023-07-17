package database

import (
	"context"
	"log"
	"os"
	"sync"

	"github.com/jackc/pgx/v5"
)

var dbOnce sync.Once

func GetDB() *pgx.Conn {
	var err error
	var database *pgx.Conn
	dbOnce.Do(func() {
		log.Println(os.Getenv("DATABASE_URL"))
		database, err = pgx.Connect(context.Background(), os.Getenv("DATABASE_URL"))
		if err != nil {
			log.Println(err)
			panic(err)
		}
		// maxOpenConn := 50
		// database.SetMaxOpenConns(maxOpenConn)
		// database.SetMaxIdleConns(50)
		// database.SetConnMaxIdleTime(2 * time.Minute)
		// database.SetConnMaxLifetime(5 * time.Minute)
	})
	return database
}
