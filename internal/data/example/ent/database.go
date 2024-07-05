// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
)

// Database is the client that holds all ent builders.
type Database struct {
	client *Client
}

// NewDatabase creates a new database configured with the given options.
func NewDatabase(opts ...Option) *Database {
	return &Database{client: NewClient(opts...)}
}

// InTx runs the given function f within a transaction.
func (db *Database) InTx(ctx context.Context, f func(context.Context) error) error {
	tx := TxFromContext(ctx)
	if tx != nil {
		return f(ctx)
	}

	tx, err := db.client.Tx(ctx)
	if err != nil {
		return fmt.Errorf("starting transaction: %w", err)
	}
	defer func() {
		if v := recover(); v != nil {
			_ = tx.Rollback()
			panic(v)
		}
	}()
	if err = f(NewTxContext(ctx, tx)); err != nil {
		if err2 := tx.Rollback(); err2 != nil {
			return fmt.Errorf("rolling back transaction: %v (original error: %w)", err2, err)
		}
		return err
	}
	return tx.Commit()
}

func (db *Database) loadClient(ctx context.Context) *Client {
	tx := TxFromContext(ctx)
	if tx != nil {
		return tx.Client()
	}
	return db.client
}

// Exec executes a query that doesn't return rows. For example, in SQL, INSERT or UPDATE.
func (db *Database) Exec(ctx context.Context, query string, args ...interface{}) (*sql.Result, error) {
	var res sql.Result
	err := db.loadClient(ctx).driver.Exec(ctx, query, args, &res)
	if err != nil {
		return nil, err
	}
	return &res, nil
}

// Query executes a query that returns rows, typically a SELECT in SQL.
func (db *Database) Query(ctx context.Context, query string, args ...interface{}) (*sql.Rows, error) {
	var rows sql.Rows
	err := db.loadClient(ctx).driver.Query(ctx, query, args, &rows)
	if err != nil {
		return nil, err
	}
	return &rows, nil
}

// Close closes the database connection and prevents new queries from starting.
func (db *Database) Close() error {
	return db.client.Close()
}

// Account is the client for interacting with the Account builders.
func (db *Database) Account(ctx context.Context) *AccountClient {
	return db.loadClient(ctx).Account
}

// OperationLog is the client for interacting with the OperationLog builders.
func (db *Database) OperationLog(ctx context.Context) *OperationLogClient {
	return db.loadClient(ctx).OperationLog
}

// Permission is the client for interacting with the Permission builders.
func (db *Database) Permission(ctx context.Context) *PermissionClient {
	return db.loadClient(ctx).Permission
}

// Role is the client for interacting with the Role builders.
func (db *Database) Role(ctx context.Context) *RoleClient {
	return db.loadClient(ctx).Role
}

// User is the client for interacting with the User builders.
func (db *Database) User(ctx context.Context) *UserClient {
	return db.loadClient(ctx).User
}
