package db

import (
	"database/sql"
	"log"
	"time"

	sq "github.com/Masterminds/squirrel"
	_ "github.com/go-sql-driver/mysql"
)

type Adaper struct {
	db *sql.DB
}

func NewAdapter(driverName, dataSourceName string) *Adaper {
	// connect
	db, err := sql.Open(driverName, dataSourceName)
	if err != nil {
		log.Fatalf("db connection failure: %v", err)
	}

	// test connection
	err = db.Ping()
	if err != nil {
		log.Fatalf("db ping failure: %v", err)
	}

	return &Adaper{db}
}

func (da Adaper) Close() {
	err := da.db.Close()
	if err != nil {
		log.Fatalf("db close failure: %v", err)
	}
}

func (da Adaper) AddToHistory(answer int32, operation string) error {
	query, args, err := sq.
		Insert("arith_history").
		Columns("date", "answer", "operation").
		Values(time.Now(), answer, operation).ToSql()
	if err != nil {
		return err
	}

	_, err = da.db.Exec(query, args...)
	if err != nil {
		return err
	}

	return nil
}
