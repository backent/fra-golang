package libs

import (
	"database/sql"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/backent/fra-golang/helpers"
	_ "github.com/go-sql-driver/mysql"
)

func NewDatabase() *sql.DB {
	MYSQL_HOST := os.Getenv("MYSQL_HOST")
	MYSQL_PORT := os.Getenv("MYSQL_PORT")
	MYSQL_USER := os.Getenv("MYSQL_USER")
	MYSQL_PASSWORD := os.Getenv("MYSQL_PASSWORD")
	MYSQL_DATABASE := os.Getenv("MYSQL_DATABASE")

	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", MYSQL_USER, MYSQL_PASSWORD, MYSQL_HOST, MYSQL_PORT, MYSQL_DATABASE)

	db, err := sql.Open("mysql", dataSourceName)
	if err != nil {
		panic(err)
	}

	DB_CONN_MAX_LIFETIME_IN_SEC, err := strconv.Atoi(os.Getenv("DB_CONN_MAX_LIFETIME_IN_SEC"))
	helpers.PanicIfError(err)

	DB_MAX_OPEN_CONNECTIONS, err := strconv.Atoi(os.Getenv("DB_MAX_OPEN_CONNECTIONS"))
	helpers.PanicIfError(err)

	DB_MAX_IDLE_CONNECTIONS, err := strconv.Atoi(os.Getenv("DB_MAX_IDLE_CONNECTIONS"))
	helpers.PanicIfError(err)

	db.SetConnMaxLifetime(time.Second * time.Duration(DB_CONN_MAX_LIFETIME_IN_SEC))
	db.SetMaxOpenConns(DB_MAX_OPEN_CONNECTIONS)
	db.SetMaxIdleConns(DB_MAX_IDLE_CONNECTIONS)

	return db
}
