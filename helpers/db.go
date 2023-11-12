package helpers

import "database/sql"

func CommitOrRollback(tx *sql.Tx) {
	err := recover()
	if err != nil {
		errRollback := tx.Rollback()
		PanifIfError(errRollback)
		panic(err)
	} else {
		errCommit := tx.Commit()
		PanifIfError(errCommit)
	}
}
