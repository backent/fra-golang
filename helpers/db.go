package helpers

import (
	"database/sql"
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
