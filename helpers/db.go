package helpers

import (
	"database/sql"
	"strconv"
	"strings"
)

func CommitOrRollback(tx *sql.Tx) {
	err := recover()
	if err != nil {
		errRollback := tx.Rollback()
		PanicIfError(errRollback)
		panic(err)
	} else {
		errCommit := tx.Commit()
		PanicIfError(errCommit)
	}
}

func Placeholders(n int) string {
	ps := make([]string, n)
	for i := 0; i < n; i++ {
		ps[i] = "?"
	}
	return strings.Join(ps, ",")
}

func PrintStringIDRelation(a ...interface{}) string {
	var printtedString string
	for _, item := range a {
		var argString string
		switch v := item.(type) {
		case string:
			argString = v
		case int:
			argString = strconv.Itoa(v)
		default:
		}

		if printtedString == "" {
			printtedString += argString
		} else {
			printtedString += "-" + argString
		}
	}

	return printtedString
}
