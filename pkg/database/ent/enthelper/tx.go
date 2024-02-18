package enthelper

import (
	"context"
	"github.com/pkg/errors"
)

type clientImp[T txImp] interface {
	Tx(ctx context.Context) (T, error)
}

type txImp interface {
	Rollback() error
	Commit() error
}

func WithTx[Q clientImp[T], T txImp](ctx context.Context, client Q, fn func(tx T) error) error {
	tx, err := client.Tx(ctx)
	if err != nil {
		return err
	}
	defer func() {
		if v := recover(); v != nil {
			_ = tx.Rollback()
			panic(v)
		}
	}()
	if err = fn(tx); err != nil {
		if rollbackErr := tx.Rollback(); rollbackErr != nil {
			err = errors.Wrap(err, rollbackErr.Error())
		}
		return err
	}
	if err = tx.Commit(); err != nil {
		return errors.Wrap(err, "committing transaction")
	}
	return nil
}
