package mysql

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

// Open connect to mysql service, auto set database connect; dsn: runner:112233@tcp(127.0.0.1:3306)/running?charset=utf8mb4&collation=utf8mb4_unicode_ci
func Open(dsn string) (err error) {
	db, err = sql.Open("mysql", dsn)
	if err != nil {
		return
	}
	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(256)
	db.SetMaxIdleConns(256)
	return
}

// GetDatabase get database connect
func GetDatabase() *sql.DB {
	return db
}

// SetDatabase set database connect
func SetDatabase(database *sql.DB) {
	db = database
}

// Query query
type Query struct {
	db   *sql.DB
	scan func(rows *sql.Rows) (err error)
}

// Scan set scan function
func (s *Query) Scan(anonymous func(rows *sql.Rows) (err error)) *Query {
	s.scan = anonymous
	return s
}

// Query execute query sql
func (s *Query) Query(prepare string, args ...interface{}) (err error) {
	var stmt *sql.Stmt
	stmt, err = s.db.Prepare(prepare)
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
	err = s.scan(rows)
	return
}

// NewQuery create a query
func NewQuery() *Query {
	return &Query{
		db: db,
	}
}

// Exec exec
type Exec struct {
	db   *sql.DB
	tx   *sql.Tx
	scan func(rows *sql.Rows) (err error)
}

// BatchExec batch execute sql
type BatchExec struct {
	Prepare string
	Args    []interface{}
}

// Begin begin a transaction
func (s *Exec) Begin() (err error) {
	s.tx, err = s.db.Begin()
	return
}

// Rollback rollback transaction
func (s *Exec) Rollback() (err error) {
	if s.tx == nil {
		err = fmt.Errorf("please open the transaction first")
		return
	}
	err = s.tx.Rollback()
	return
}

// Commit commit transaction
func (s *Exec) Commit() (err error) {
	if s.tx == nil {
		err = fmt.Errorf("please open the transaction first")
		return
	}
	err = s.tx.Commit()
	return
}

// Scan set scan function
func (s *Exec) Scan(anonymous func(rows *sql.Rows) (err error)) *Exec {
	s.scan = anonymous
	return s
}

// Query execute a query sql, transaction priority
func (s *Exec) Query(prepare string, args ...interface{}) (err error) {
	var stmt *sql.Stmt
	if s.tx != nil {
		stmt, err = s.tx.Prepare(prepare)
	} else {
		stmt, err = s.db.Prepare(prepare)
	}
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
	err = s.scan(rows)
	return
}

// Exec execute a execute sql, transaction priority
func (s *Exec) Exec(prepare string, args ...interface{}) (rowsAffected int64, err error) {
	var stmt *sql.Stmt
	if s.tx != nil {
		stmt, err = s.tx.Prepare(prepare)
	} else {
		stmt, err = s.db.Prepare(prepare)
	}
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

// AddOne insert a piece of data and get the self-increment value of the inserted row, transaction priority
func (s *Exec) AddOne(prepare string, args ...interface{}) (lastId int64, err error) {
	var stmt *sql.Stmt
	if s.tx != nil {
		stmt, err = s.tx.Prepare(prepare)
	} else {
		stmt, err = s.db.Prepare(prepare)
	}
	if err != nil {
		return
	}
	defer stmt.Close()
	var result sql.Result
	result, err = stmt.Exec(args...)
	if err != nil {
		return
	}
	lastId, err = result.LastInsertId()
	return
}

// Transaction closure execution transaction, automatic rollback on error
func (s *Exec) Transaction(times int, anonymous func(exe *Exec) (err error)) (err error) {
	if times <= 0 {
		err = fmt.Errorf("the number of transactions executed by the database has been used up")
		return
	}
	for i := 0; i < times; i++ {
		s.tx, err = s.db.Begin()
		if err != nil {
			continue
		}
		err = anonymous(s)
		if err != nil {
			_ = s.tx.Rollback()
			continue
		}
		_ = s.tx.Commit()
		break
	}
	return
}

// BatchExec batch exec
func (s *Exec) BatchExec(batch ...*BatchExec) (err error) {
	err = s.Transaction(3, func(exe *Exec) (err error) {
		for _, b := range batch {
			_, err = s.Exec(b.Prepare, b.Args...)
			if err != nil {
				break
			}
		}
		return
	})
	return
}

// NewExec create a exec
func NewExec() *Exec {
	return &Exec{
		db: db,
	}
}
