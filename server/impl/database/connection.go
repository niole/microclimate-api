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

func GetConnectionPool() *sql.DB {
	if pool == nil {
		var err error

		pool, err = sql.Open("mysql", "niole:pw@tcp(127.0.0.1:3306)/default")
		if err != nil {
			panic(err)
		}
		//defer pool.Close()

		pool.SetConnMaxLifetime(time.Minute * 3)
		pool.SetMaxOpenConns(10)
		pool.SetMaxIdleConns(10)

		ctx, stop := context.WithCancel(context.Background())
		defer stop()

		appSignal := make(chan os.Signal, 3)
		signal.Notify(appSignal, os.Interrupt)

		go func() {
			select {
			case <-appSignal:
				stop()
			}
		}()

		ctx, cancel := context.WithTimeout(ctx, 1*time.Second)
		defer cancel()

		DoPing(pool)

		// create tables

		_, tableCreateErr := pool.ExecContext(ctx,
			`CREATE TABLE IF NOT EXISTS Deployments (
                Id varchar(36) PRIMARY KEY NOT NULL,
                OwnerUserId varchar(36) NOT NULL,
                Domain varchar(255) NOT NULL,
                Status ENUM('unreachable', 'starting_up', 'failed', 'running', 'shutting_down') NOT NULL
            );`,
		)
		if tableCreateErr != nil {
			log.Printf("Failed to create deployments table. error: %v", tableCreateErr)
		}
	}
	return pool
}

func DoPing(pool *sql.DB) {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()
	if err := pool.PingContext(ctx); err != nil {
		log.Fatalf("unable to connect to database: %v", err)
	}
}
