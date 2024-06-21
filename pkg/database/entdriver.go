package database

import (
	"database/sql"
	"entgo.io/ent/dialect"
	entsql "entgo.io/ent/dialect/sql"
	"github.com/XSAM/otelsql"
	"github.com/eiixy/monorepo/internal/pkg/config"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
	semconv "go.opentelemetry.io/otel/semconv/v1.25.0"
	"time"
)

func NewEntDriver(cfg config.Database) (dialect.Driver, error) {
	driver := dialect.MySQL
	if cfg.Driver != "" {
		driver = cfg.Driver
	}
	db, err := sql.Open(driver, cfg.Dsn)
	if err != nil {
		return nil, err
	}
	return newEntDriver(driver, db, cfg)
}

func NewEntDriverWithOtel(cfg config.Database, opts ...otelsql.Option) (dialect.Driver, error) {
	driver := dialect.MySQL
	if cfg.Driver != "" {
		driver = cfg.Driver
	}
	opts = append(opts, otelsql.WithAttributes(
		semconv.DBSystemKey.String(driver),
	))
	db, err := otelsql.Open(driver, cfg.Dsn, opts...)
	if err != nil {
		return nil, err
	}
	err = otelsql.RegisterDBStatsMetrics(db, otelsql.WithAttributes(
		semconv.DBSystemMySQL,
	))
	if err != nil {
		return nil, err
	}
	return newEntDriver(driver, db, cfg)
}

func newEntDriver(driver string, db *sql.DB, cfg config.Database) (dialect.Driver, error) {
	if cfg.MaxOpenConns != 0 {
		db.SetMaxOpenConns(cfg.MaxOpenConns)
	}
	if cfg.MaxIdleConns != 0 {
		db.SetMaxIdleConns(cfg.MaxIdleConns)
	}
	if cfg.ConnMaxIdleTime != "" {
		duration, err := time.ParseDuration(cfg.ConnMaxIdleTime)
		if err != nil {
			return nil, err
		}
		db.SetConnMaxIdleTime(duration)
	}
	if cfg.ConnMaxLifetime != "" {
		duration, err := time.ParseDuration(cfg.ConnMaxLifetime)
		if err != nil {
			return nil, err
		}
		db.SetConnMaxLifetime(duration)
	}
	return entsql.OpenDB(driver, db), nil
}
