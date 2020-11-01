package database

import (
	"context"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"os"
	"os/signal"
	"time"
)

var (
	pool *sql.DB
)

// should just initialize tables on startup
func Get(initBlock func(context.Context, *sql.DB) error) *sql.DB {
	stopCtx, stop := context.WithCancel(context.Background())
	defer stop()

	if pool == nil {
		var err error

		pool, err = sql.Open("mysql", "niole:pw@tcp(127.0.0.1:3306)/default")
		if err != nil {
			panic(err)
		}
		//defer pool.Close()

		pool.SetConnMaxLifetime(time.Minute * 3)
		pool.SetMaxOpenConns(3)
		pool.SetMaxIdleConns(3)

		appSignal := make(chan os.Signal, 3)
		signal.Notify(appSignal, os.Interrupt)

		go func() {
			select {
			case <-appSignal:
				stop()
			}
		}()
	}

	DoPing(stopCtx, pool)

	// TODO this gets the table init code out, but it happens too often
	initBlock(stopCtx, pool)

	return pool
}

func DoPing(ctx context.Context, pool *sql.DB) {
	if err := pool.PingContext(ctx); err != nil {
		log.Fatalf("unable to connect to database: %v", err)
	}
}
