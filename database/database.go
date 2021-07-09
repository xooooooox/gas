package database

import (
	"database/sql"
	"errors"
)

// Query query one sql
func Query(db *sql.DB, fc func(rows *sql.Rows) error, prepare string, args ...interface{}) (err error) {
	if db == nil {
		err = errors.New("`db` is nil")
		return
	}
	if fc == nil {
		err = errors.New("`fc` is nil")
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
		err = errors.New("`db` is nil")
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

// Tx begin a transaction
func Tx(db *sql.DB, fc func(tx *sql.Tx) error) (err error) {
	if db == nil {
		err = errors.New("`db` is nil")
		return
	}
	if fc == nil {
		err = errors.New("`fc` is nil")
		return
	}
	var tx *sql.Tx
	tx, err = db.Begin()
	if err != nil {
		return
	}
	err = fc(tx)
	if err != nil {
		_ = tx.Rollback()
		return
	}
	_ = tx.Commit()
	return
}

// TxQuery query one sql in transaction
func TxQuery(tx *sql.Tx, fc func(rows *sql.Rows) error, prepare string, args ...interface{}) (err error) {
	if tx == nil {
		err = errors.New("`tx` is nil")
		return
	}
	if fc == nil {
		err = errors.New("`fc` is nil")
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

// TxExec execute one sql in transaction
func TxExec(tx *sql.Tx, prepare string, args ...interface{}) (rowsAffected int64, err error) {
	if tx == nil {
		err = errors.New("`tx` is nil")
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