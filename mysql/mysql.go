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
	return NewSql().OneStepQuery(anonymous, prepare, args...)
}

func Exec(prepare string, args ...interface{}) (int64, error) {
	return NewSql().OneStepExec(prepare, args...)
}

func AddOne(prepare string, args ...interface{}) (int64, error) {
	return NewSql().OneStepAddOne(prepare, args...)
}

type Sql struct {
	db      *sql.DB                          // database connection object
	tx      *sql.Tx                          // database transaction object
	prepare string                           // sql statement to be executed
	args    []interface{}                    // executed sql parameters
	scan    func(rows *sql.Rows) (err error) // scan query results
}

func NewSql() *Sql {
	return &Sql{
		db: db,
	}
}

func (s *Sql) Begin() (err error) {
	s.tx, err = s.db.Begin()
	return
}

func (s *Sql) Rollback() (err error) {
	if s.tx == nil {
		err = ErrorTransactionNotOpened
		return
	}
	err = s.tx.Rollback()
	s.tx = nil
	return
}

func (s *Sql) Commit() (err error) {
	if s.tx == nil {
		err = ErrorTransactionNotOpened
		return
	}
	err = s.tx.Commit()
	s.tx = nil
	return
}

func (s *Sql) Scan(anonymous func(rows *sql.Rows) (err error)) *Sql {
	s.scan = anonymous
	return s
}

func (s *Sql) Prepare(prepare string) *Sql {
	s.prepare = prepare
	return s
}

func (s *Sql) Args(args ...interface{}) *Sql {
	s.args = args
	return s
}

func (s *Sql) Stmt() (stmt *sql.Stmt, err error) {
	if s.tx != nil {
		stmt, err = s.tx.Prepare(s.prepare)
	} else {
		stmt, err = s.db.Prepare(s.prepare)
	}
	return
}

func (s *Sql) FetchSql() (prepare string, args []interface{}) {
	prepare, args = s.prepare, s.args
	return
}

func (s *Sql) Query() (err error) {
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

func (s *Sql) Exec() (rowsAffected int64, err error) {
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

func (s *Sql) AddOne() (lastId int64, err error) {
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

func (s *Sql) OneStepQuery(anonymous func(rows *sql.Rows) (err error), prepare string, args ...interface{}) (err error) {
	err = s.Scan(anonymous).Prepare(prepare).Args(args...).Query()
	return
}

func (s *Sql) OneStepExec(prepare string, args ...interface{}) (int64, error) {
	return s.Prepare(prepare).Args(args...).Exec()
}

func (s *Sql) OneStepAddOne(prepare string, args ...interface{}) (int64, error) {
	return s.Prepare(prepare).Args(args...).AddOne()
}

// Transaction closure execute transaction, automatic rollback on error
func (s *Sql) Transaction(times int, anonymous func(exe *Sql) (err error)) (err error) {
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
