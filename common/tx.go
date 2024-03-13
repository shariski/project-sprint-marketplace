package common

import (
	"database/sql"
	"project-sprint-marketplace/exception"
)

func CommitOrRollback(tx *sql.Tx) {
	err := recover()
	if err != nil {
		errorRollback := tx.Rollback()
		exception.PanicLogging(errorRollback)
		panic(err)
	} else {
		errorCommit := tx.Commit()
		exception.PanicLogging(errorCommit)
	}
}