package database

import (
	"database/sql"
	"errors"
)

// Query query one sql
func Query(db *sql.DB, fc func(rows *sql.Rows) error, prepare string, args ...interface{}) (err error) {
	if db == nil {
		err = errors.New("unavailable of database connection")
		return
	}
	if fc == nil {
		err = errors.New("unavailable of scanning closure")
		return
	}
	var stmt *sql.Stmt
	stmt, err = db.Prepare(prepare)
	if err != nil {
		return
	}
	defer stmt.Close()
	var rows *sql.Rows
	rows, err = stmt.Query(args...)
	if err != nil {
		return
	}
	defer rows.Close()
	err = fc(rows)
	return
}

// Exec execute one sql
func Exec(db *sql.DB, prepare string, args ...interface{}) (rowsAffected int64, err error) {
	if db == nil {
		err = errors.New("unavailable of database connection")
		return
	}
	var stmt *sql.Stmt
	stmt, err = db.Prepare(prepare)
	if err != nil {
		return
	}
	defer stmt.Close()
	var result sql.Result
	result, err = stmt.Exec(args...)
	if err != nil {
		return
	}
	rowsAffected, err = result.RowsAffected()
	return
}

// Ask begin a transaction
func Ask(db *sql.DB, fc func(tx *sql.Tx) error) (err error) {
	if db == nil {
		err = errors.New("unavailable of database connection")
		return
	}
	if fc == nil {
		err = errors.New("unavailable of transaction closure")
		return
	}
	var tx *sql.Tx
	tx, err = db.Begin()
	if err != nil {
		return
	}
	err = fc(tx)
	if err != nil {
		tx.Rollback()
		return
	}
	tx.Commit()
	return
}

// AskQuery query one sql in transaction
func AskQuery(tx *sql.Tx, fc func(rows *sql.Rows) error, prepare string, args ...interface{}) (err error) {
	if tx == nil {
		err = errors.New("unavailable of database transaction")
		return
	}
	if fc == nil {
		err = errors.New("unavailable of scanning closure")
		return
	}
	var stmt *sql.Stmt
	stmt, err = tx.Prepare(prepare)
	if err != nil {
		return
	}
	defer stmt.Close()
	var rows *sql.Rows
	rows, err = stmt.Query(args...)
	if err != nil {
		return
	}
	defer rows.Close()
	err = fc(rows)
	return
}

// AskExec execute one sql in transaction
func AskExec(tx *sql.Tx, prepare string, args ...interface{}) (rowsAffected int64, err error) {
	if tx == nil {
		err = errors.New("unavailable of database transaction")
		return
	}
	var stmt *sql.Stmt
	stmt, err = tx.Prepare(prepare)
	if err != nil {
		return
	}
	defer stmt.Close()
	var result sql.Result
	result, err = stmt.Exec(args...)
	if err != nil {
		return
	}
	rowsAffected, err = result.RowsAffected()
	return
}
