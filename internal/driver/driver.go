package driver

import (
	"database/sql"
	"time"

	_ "github.com/jackc/pgconn"
	_ "github.com/jackc/pgx/v4"
	_ "github.com/jackc/pgx/v4/stdlib"
)

type Db struct {
	SQL *sql.DB
}

var dbCon = &Db{}

const maxOpenCon = 10
const maxIdleCon = 5
const maxDbLife = 5 * time.Minute

func ConnectSQL(dsn string) (*Db, error) {
	d, err := NewDatabase(dsn)
	if err != nil {
		panic(err)
	}
	d.SetMaxOpenConns(maxOpenCon)
	d.SetMaxIdleConns(maxIdleCon)
	d.SetConnMaxLifetime(maxDbLife)
	dbCon.SQL = d
	err = testDB(d)
	if err != nil {
		return nil, err
	}
	return dbCon, nil
}

func testDB(d *sql.DB) error {
	err := d.Ping()
	if err != nil {
		return err
	}
	return nil

}

func NewDatabase(dsn string) (*sql.DB, error) {
	db, err := sql.Open("pgx", dsn)
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}
