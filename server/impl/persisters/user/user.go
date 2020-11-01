package user

import (
	api "api/protobuf"
	connection "api/server/impl/database"
	"context"
	"database/sql"
	"log"
)

func findByEmail(ctx context.Context, db *sql.DB, email string) (*api.User, error) {
	var user api.User

	err := ScanOneUser(db.QueryRowContext(
		ctx,
		`SELECT * FROM Users WHERE Email = ?;`,
		&email,
	).Scan, &user)

	if err != nil {
		log.Printf(
			"Failed to get user by email %v, error %v",
			email,
			err,
		)
	}
	return &user, nil
}

func findById(ctx context.Context, db *sql.DB, userId string) (*api.User, error) {
	var user api.User

	err := ScanOneUser(db.QueryRowContext(
		ctx,
		`SELECT * FROM Users WHERE Id = ?;`,
		&userId,
	).Scan, &user)

	if err != nil {
		log.Printf(
			"Failed to get user with Id %v, error %v",
			userId,
			err,
		)
	}

	return &user, nil
}

func CreateUser(ctx context.Context, newUserEmail string) (*api.User, error) {
	db := connection.GetConnectionPool()

	_, err := db.ExecContext(
		ctx,
		`INSERT INTO Users (Id, Email) VALUES (UUID(), ?);`,
		&newUserEmail,
	)

	if err != nil {
		log.Printf("Failed to insert new user %v, error %v", newUserEmail, err)
	} else {
		return findByEmail(ctx, db, newUserEmail)
	}

	return nil, err
}

func UpdateUserEmail(ctx context.Context, userId string, newEmail string) (*api.User, error) {
	db := connection.GetConnectionPool()

	_, err := db.ExecContext(
		ctx,
		`UPDATE Users
		SET Email = ?
		WHERE Id = ?;`,
		&newEmail,
		&userId,
	)

	if err != nil {
		log.Printf("Failed to update user email %v %v, error %v", userId, newEmail, err)
	} else {
		return findById(ctx, db, userId)
	}

	return nil, err
}

func RemoveUser(ctx context.Context, userId string) (*api.Empty, error) {
	db := connection.GetConnectionPool()

	_, err := db.ExecContext(
		ctx,
		`DELETE FROM Users
		WHERE Id = ?;`,
		&userId,
	)

	return &api.Empty{}, err
}

func GetUser(ctx context.Context, userId *string, email *string) (*api.User, error) {
	db := connection.GetConnectionPool()

	if userId != nil {
		return findById(ctx, db, *userId)
	}

	if email != nil {
		return findByEmail(ctx, db, *email)
	}

	log.Print("A user Id or an email is required in order to find a user")

	return nil, nil
}
