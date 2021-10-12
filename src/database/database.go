package database

import (
	"database/sql"
	"errors"
	"os"

	// loading postgresql driver
	"github.com/google/uuid"
	_ "github.com/lib/pq"
)

type UserT struct {
    UserID int
    Name string
	Email string
	Password string
}

var database *sql.DB

func Connect() error {
    var err error

    if database != nil {
        database.Close()
        database = nil
    }

    param := os.ExpandEnv(
        `host=${DB_HOST}
         port=${DB_PORT}
         user=${DB_USER}
         password=${DB_PWD}
         dbname=${DB_NAME}
         sslmode=${DB_SSLMODE}`)

    database, err = sql.Open("postgres", param)
    if err != nil {
        return DBError{ "Connect: connecting to database failed", err }
    }

    err = database.Ping()
    if err != nil {
        return DBError{ "Connect: pinging database failed", err }
    }

    return nil
}

func Initialize() error {
    if database == nil {
        return errors.New("Initialize: not connected to database")
    }

    // Verify connection to database
    err := database.Ping()
    if err != nil {
        database.Close()
        return DBError{ "Initialize: pinging database failed", err }
    }

	_, err = database.Exec(
        `CREATE TABLE IF NOT EXISTS users (
       	UserID serial PRIMARY KEY,
        Name varchar,
		Email varchar,
        Password varchar)`)
    if err != nil {
        database.Close()
        return DBError{ "Initialize: creating users table failed", err }
    }

	_, err = database.Exec(
        `CREATE TABLE IF NOT EXISTS attendance (
       	Code uuid PRIMARY KEY,
        UserID integer REFERENCES users (UserID) ON DELETE CASCADE,
        CheckedIn boolean)`)
    if err != nil {
        database.Close()
        return DBError{ "Initialize: creating attendance table failed", err }
    }

	return nil
}

func Close() error {
    if database == nil {
        return errors.New("Close: not connected to database")
    }

    database.Close()

    return nil
}

func Register(user UserT) (string, error) {
    if database == nil {
		return "", errors.New("Register: not connected to database")
	}

    // Verify connection to database
	err := database.Ping()
	if err != nil {
		database.Close()
		return "", DBError{ "Register: pinging database failed", err }
	}

    u, err := uuid.NewRandom()
	code := u.String()
    if err != nil || code == "" {
		return "", errors.New("Register: could not create uuid")
	}

    // add user to users table
	err = database.QueryRow(
		`INSERT INTO users (Name, Email, Password) VALUES ($1, $2, $3) RETURNING UserID`,
    user.Name, user.Email, user.Password).Scan(&user.UserID)
	if err != nil {
		return "",DBError{ "Register: inserting user into users table failed", err }
	}

    // add user to attendance table
	err = database.QueryRow(
		`INSERT INTO attendance (Code, UserID, CheckedIn) VALUES ($1, $2, $3)`,
    code, user.UserID, false).Err()
	if err != nil {
		return "",DBError{ "Register: inserting user into attendance table failed", err }
	}

    return code, nil
}
