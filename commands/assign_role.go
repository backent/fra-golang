package main

import (
	"database/sql"
	"flag"

	"github.com/backent/fra-golang/helpers"
	"github.com/backent/fra-golang/libs"
	"github.com/joho/godotenv"
)

func main() {

	var nik, role string
	scanArg(&nik, &role)

	err := godotenv.Load()
	helpers.PanicIfError(err)

	sql := libs.NewDatabase()
	tx, err := sql.Begin()
	helpers.PanicIfError(err)

	defer helpers.CommitOrRollback(tx)

	checkIfUserExists(tx, nik)

	setRoleToNik(tx, nik, role)
}

func scanArg(nik *string, role *string) {
	flag.StringVar(nik, "nik", "", "nik to change the role")
	flag.StringVar(role, "role", "", "role that need to assign to role")

	flag.Parse()
	if *nik == "" || *role == "" {
		panic("required nik and role to be present")
	}
}

func checkIfUserExists(tx *sql.Tx, nik string) {
	query := "SELECT id FROM users WHERE nik = ?"
	rows, err := tx.Query(query, nik)
	helpers.PanicIfError(err)
	defer rows.Close()

	if !rows.Next() {
		panic("nik not found")
	}

}

func setRoleToNik(tx *sql.Tx, nik string, role string) {
	query := "UPDATE users SET role = ? WHERE nik = ? "
	_, err := tx.Exec(query, role, nik)
	helpers.PanicIfError(err)
}
