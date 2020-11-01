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

		stopCtx, stop := context.WithCancel(context.Background())
		defer stop()

		appSignal := make(chan os.Signal, 3)
		signal.Notify(appSignal, os.Interrupt)

		go func() {
			select {
			case <-appSignal:
				stop()
			}
		}()

		ctx, cancel := context.WithTimeout(stopCtx, 5*time.Second)
		defer cancel()

		DoPing(pool)

		log.Print("Creating tables if not exists")

		_, deploymentTableCreateErr := pool.ExecContext(ctx,
			`CREATE TABLE IF NOT EXISTS Deployments (
                Id varchar(36) PRIMARY KEY NOT NULL,
                OwnerUserId varchar(36) NOT NULL,
                Domain varchar(255) NOT NULL,
                Status ENUM('unreachable', 'starting_up', 'failed', 'running', 'shutting_down') NOT NULL
            );`,
		)
		if deploymentTableCreateErr != nil {
			log.Fatalf("Failed to create deployments table. error: %v", deploymentTableCreateErr)
		}

		_, peripheralTableCreateErr := pool.ExecContext(ctx,
			`CREATE TABLE IF NOT EXISTS Peripherals (
                Id varchar(36) PRIMARY KEY NOT NULL,
                OwnerUserId varchar(36) NOT NULL,
                DeploymentId varchar(36) NOT NULL,
                HardwareId varchar(36) NOT NULL,
                Type ENUM('THERMAL', 'PARTICLE') NOT NULL
            );`,
		)
		if peripheralTableCreateErr != nil {
			log.Fatalf("Failed to create peripherals table. error: %v", peripheralTableCreateErr)
		}

		_, peripheralEventsTableCreateErr := pool.ExecContext(ctx,
			`CREATE TABLE IF NOT EXISTS PeripheralEvents (
                Id varchar(36) PRIMARY KEY NOT NULL,
				PeripheralId varchar(36) NOT NULL,
                DeploymentId varchar(36) NOT NULL,
                Value smallint NOT NULL,
				Timestamp timestamp NOT NULL
            );`,
		)
		if peripheralEventsTableCreateErr != nil {
			log.Fatalf(
				"Failed to create peripheral events table. error: %v",
				peripheralEventsTableCreateErr,
			)
		}

		_, userTableCreateErr := pool.ExecContext(ctx,
			`CREATE TABLE IF NOT EXISTS Users (
                Id varchar(36) PRIMARY KEY NOT NULL,
				Email varchar(255) NOT NULL UNIQUE
            );`,
		)

		if userTableCreateErr != nil {
			log.Fatalf(
				"Failed to create user table. error: %v",
				userTableCreateErr,
			)
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
