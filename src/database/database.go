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
    UserID int      `json:",omitempty"`
    Name string
	Email string
	Password string `json:",omitempty"`
}

type AttendantT struct {
    Code string
    UserID int
    CheckedIn bool
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
        return InternalServerError{ "Connect: connecting to database failed", err }
    }

    err = database.Ping()
    if err != nil {
        return InternalServerError{ "Connect: pinging database failed", err }
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
        return InternalServerError{ "Initialize: pinging database failed", err }
    }

	_, err = database.Exec(
        `CREATE TABLE IF NOT EXISTS users (
       	UserID serial PRIMARY KEY,
        Name varchar,
		Email varchar,
        Password varchar)`)
    if err != nil {
        database.Close()
        return InternalServerError{ "Initialize: creating users table failed", err }
    }

	_, err = database.Exec(
        `CREATE TABLE IF NOT EXISTS attendance (
       	Code uuid PRIMARY KEY,
        UserID integer REFERENCES users (UserID) ON DELETE CASCADE,
        CheckedIn boolean)`)
    if err != nil {
        database.Close()
        return InternalServerError{ "Initialize: creating attendance table failed", err }
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
		return "", InternalServerError{ "Register: pinging database failed", err }
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
		return "",InternalServerError{ "Register: inserting user into users table failed", err }
	}

    // add user to attendance table
	err = database.QueryRow(
		`INSERT INTO attendance (Code, UserID, CheckedIn) VALUES ($1, $2, $3)`,
    code, user.UserID, false).Err()
	if err != nil {
		return "",InternalServerError{ "Register: inserting user into attendance table failed", err }
	}

    return code, nil
}

func GetAttendants() ([]UserT, error) {
    if database == nil {
		return nil, errors.New("GetAttendants: not connected to database")
	}

    // Verify connection to database
	err := database.Ping()
	if err != nil {
		database.Close()
		return nil, InternalServerError{ "GetAttendants: pinging database failed", err }
	}

    rows, err := database.Query(
        `SELECT U.Name, U.Email FROM users U
        WHERE U.UserID in (
            SELECT A.UserID from attendance A
            WHERE A.CheckedIn = True
        )`)
	if err != nil {
		return nil, InternalServerError{ "GetAttendants: querying attendance from database failed", err }
	}

    var attendants []UserT
    for rows.Next() {
        var user UserT

        err := rows.Scan(&user.Name, &user.Email)
		if err != nil {
			return nil, InternalServerError{ "GetAttendants: parsing attendance failed", err }
		}

		attendants = append(attendants, user)
	}

    return attendants, nil
}

/*
 * SetAttendant checks in one user by its uuid
 *
 * Possible http error types:
 * StatusInternalServerError, StatusBadRequest
 */
func SetAttendant(code string) error {
    if database == nil {
		return errors.New("SetAttendant: not connected to database")
	}

    // Verify connection to database
	err := database.Ping()
	if err != nil {
		database.Close()
		return InternalServerError{ "SetAttendant: pinging database failed", err }
	}

    // Check validity of code(uuid) format
    _, err = uuid.Parse(code)
	if err != nil {
		return InternalServerError{ "SetAttendant: malformed code(uuid)", err }
	}

    res, err := database.Exec(`UPDATE attendance SET CheckedIn = True WHERE Code = $1`, code)
    if err != nil {
        return InternalServerError{ "SetAttendant: updating CheckedIn status in database failed", err }
    }

    n, err := res.RowsAffected();
    if err != nil {
        return InternalServerError{ "SetAttendant: Reading number of affected rows failed, update status unknown", err }
    }
    if n == 0 {
        return InternalServerError{ "SetAttendant: No user with specified uuid found", nil }
    }

    return nil
}
