// mysql database call

package mysql

import (
	"database/sql"
	"errors"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

// db database object
var db *sql.DB

// ErrorTransactionNotOpened transaction not opened
var ErrorTransactionNotOpened = errors.New("mysql: please open the transaction first")

// Open connect to mysql service, auto set database connect; dsn: runner:112233@tcp(127.0.0.1:3306)/running?charset=utf8mb4&collation=utf8mb4_unicode_ci
func Open(dsn string) (err error) {
	db, err = sql.Open("mysql", dsn)
	if err != nil {
		return
	}
	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(512)
	db.SetMaxIdleConns(128)
	return
}

func Db0(database *sql.DB) {
	db = database
}

func Db1() *sql.DB {
	return db
}

func Query(anonymous func(rows *sql.Rows) (err error), prepare string, args ...interface{}) error {
	return NewExec().OneStepQuery(anonymous, prepare, args...)
}

func Execute(prepare string, args ...interface{}) (int64, error) {
	return NewExec().OneStepExecute(prepare, args...)
}

func AddOne(prepare string, args ...interface{}) (int64, error) {
	return NewExec().OneStepAddOne(prepare, args...)
}

// Exec mysql database sql statement execute object
type Exec struct {
	db      *sql.DB                          // database connection object
	tx      *sql.Tx                          // database transaction object
	prepare string                           // sql statement to be executed
	args    []interface{}                    // executed sql parameters
	scan    func(rows *sql.Rows) (err error) // scan query results
}

func NewExec() *Exec {
	return &Exec{
		db: db,
	}
}

func (s *Exec) Begin() (err error) {
	s.tx, err = s.db.Begin()
	return
}

func (s *Exec) Rollback() (err error) {
	if s.tx == nil {
		err = ErrorTransactionNotOpened
		return
	}
	err = s.tx.Rollback()
	s.tx = nil
	return
}

func (s *Exec) Commit() (err error) {
	if s.tx == nil {
		err = ErrorTransactionNotOpened
		return
	}
	err = s.tx.Commit()
	s.tx = nil
	return
}

func (s *Exec) Scan(anonymous func(rows *sql.Rows) (err error)) *Exec {
	s.scan = anonymous
	return s
}

func (s *Exec) Prepare(prepare string) *Exec {
	s.prepare = prepare
	return s
}

func (s *Exec) Args(args ...interface{}) *Exec {
	s.args = args
	return s
}

func (s *Exec) Stmt() (stmt *sql.Stmt, err error) {
	if s.tx != nil {
		stmt, err = s.tx.Prepare(s.prepare)
	} else {
		stmt, err = s.db.Prepare(s.prepare)
	}
	return
}

func (s *Exec) FetchSql() (prepare string, args []interface{}) {
	prepare, args = s.prepare, s.args
	return
}

func (s *Exec) Query() (err error) {
	var stmt *sql.Stmt
	stmt, err = s.Stmt()
	if err != nil {
		return
	}
	defer stmt.Close()
	var rows *sql.Rows
	rows, err = stmt.Query(s.args...)
	if err != nil {
		return
	}
	defer rows.Close()
	err = s.scan(rows)
	return
}

func (s *Exec) Execute() (rowsAffected int64, err error) {
	var stmt *sql.Stmt
	stmt, err = s.Stmt()
	if err != nil {
		return
	}
	defer stmt.Close()
	var result sql.Result
	result, err = stmt.Exec(s.args...)
	if err != nil {
		return
	}
	rowsAffected, err = result.RowsAffected()
	return
}

func (s *Exec) AddOne() (lastId int64, err error) {
	var stmt *sql.Stmt
	stmt, err = s.Stmt()
	if err != nil {
		return
	}
	defer stmt.Close()
	var result sql.Result
	result, err = stmt.Exec(s.args...)
	if err != nil {
		return
	}
	lastId, err = result.LastInsertId()
	return
}

func (s *Exec) OneStepQuery(anonymous func(rows *sql.Rows) (err error), prepare string, args ...interface{}) (err error) {
	err = s.Scan(anonymous).Prepare(prepare).Args(args...).Query()
	return
}

func (s *Exec) OneStepExecute(prepare string, args ...interface{}) (int64, error) {
	return s.Prepare(prepare).Args(args...).Execute()
}

func (s *Exec) OneStepAddOne(prepare string, args ...interface{}) (int64, error) {
	return s.Prepare(prepare).Args(args...).AddOne()
}

// Transaction closure execute transaction, automatic rollback on error
func (s *Exec) Transaction(times int, anonymous func(exe *Exec) (err error)) (err error) {
	if times <= 0 {
		err = fmt.Errorf("mysql: the number of transactions executed by the database has been used up")
		return
	}
	for i := 0; i < times; i++ {
		err = s.Begin()
		if err != nil {
			continue
		}
		err = anonymous(s)
		if err != nil {
			_ = s.Rollback()
			continue
		}
		_ = s.Commit()
		break
	}
	return
}
