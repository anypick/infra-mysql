package basesql

import (
	"context"
	"database/sql"
)

/**
定义mysql事务的处理方式
*/
type Runner struct {
	Db *sql.DB
	Tx *sql.Tx
}

// 事务执行, 一般在service层调用
func DbTxRunner(fn func(runner *Runner) error) error {
	db := GetDb()
	tx, err := db.Begin()
	if err != nil {
		return err
	}
	runner := &Runner{Db: db, Tx: tx}
	if err = fn(runner); err != nil {
		tx.Rollback()
		return err
	}
	err = tx.Commit()
	if err != nil {
		panic(err)
	}
	return nil
}

// 针对纯查询的Service
func DbWithQuery(fn func(runner *Runner) error) error {
	db := GetDb()
	runner := &Runner{Db: db, Tx: nil}
	if err := fn(runner); err != nil {
		return err
	}
	return nil
}

const (
	TX = "TX"
)

func WithValueContext(parent context.Context, runner *Runner) context.Context {
	return context.WithValue(parent, TX, runner)
}

// 多事务处理方案
func ExecuteContext(ctx context.Context, fn func(runner *Runner) error) error {

	if ctx == nil{
		// 执行单个事务
		return DbTxRunner(fn)
	}
	runner := ctx.Value(TX).(*Runner)
	if runner.Tx == nil {
		return DbWithQuery(fn)
	}
	return fn(runner)
}
