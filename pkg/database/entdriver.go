package database

import (
	"database/sql"
	"entgo.io/ent/dialect"
	entsql "entgo.io/ent/dialect/sql"
	"github.com/eiixy/monorepo/internal/pkg/config"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
	_ "github.com/mattn/go-sqlite3"
	"strconv"
	"time"
)

func NewEntDriver(database config.Database) (dialect.Driver, error) {
	driver := dialect.MySQL
	if database.Driver != "" {
		driver = database.Driver
	}
	db, err := sql.Open(driver, database.Dsn)
	if err != nil {
		return nil, err
	}
	if database.MaxOpenConns != "" {
		moc, err := strconv.Atoi(database.MaxOpenConns)
		if err != nil {
			return nil, err
		}
		db.SetMaxOpenConns(moc)
	}
	if database.MaxIdleConns != "" {
		mic, err := strconv.Atoi(database.MaxIdleConns)
		if err != nil {
			return nil, err
		}
		db.SetMaxIdleConns(mic)
	}
	if database.ConnMaxIdleTime != "" {
		cmt, err := strconv.Atoi(database.ConnMaxIdleTime)
		if err != nil {
			return nil, err
		}
		db.SetConnMaxIdleTime(time.Duration(cmt) * time.Second)
	}
	if database.ConnMaxLifetime != "" {
		cml, err := strconv.Atoi(database.ConnMaxLifetime)
		if err != nil {
			return nil, err
		}
		db.SetConnMaxLifetime(time.Duration(cml) * time.Second)
	}
	drv := entsql.OpenDB(driver, db)
	return drv, nil
}
