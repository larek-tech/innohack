package postgres

import (
	"context"
	"errors"
	"time"

	"github.com/georgysavva/scany/v2/pgxscan"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/rs/zerolog/log"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
)

type ctxKey string

const (
	txKey ctxKey = "tx"
)

// Postgres wraps pgxpool.Pool and adds tracer to all operations.
type Postgres struct {
	pool   *pgxpool.Pool
	tracer trace.Tracer
}

// MustNew creates a new Postgres instance or panics if failed.
func MustNew(cfg *Config, tracer trace.Tracer) *Postgres {
	var (
		err         error
		pool        *pgxpool.Pool
		ctx, cancel = context.WithTimeout(context.Background(), time.Duration(cfg.RetryTimeout)*time.Second)
	)
	defer cancel()

	for ctx.Err() == nil {
		time.Sleep(time.Second)
		pool, err = pgxpool.New(ctx, cfg.URL())
		if err != nil {
			log.Err(err).Msg("failed to open new pool")
			continue
		}

		if err = pool.Ping(ctx); err != nil {
			log.Err(err).Msg("can't access postgres")
			continue
		}

		return &Postgres{pool: pool, tracer: tracer}
	}
	log.Panic().Msg("failed to connect to postgres")
	return nil
}

// GetPool returns the underlying pgxpool.Pool.
func (pg Postgres) GetPool() *pgxpool.Pool {
	return pg.pool
}

// Close closes the underlying db connection.
func (pg Postgres) Close() {
	pg.pool.Close()
}

// BeginSerializable starts a new transaction with serializable isolation level.
func (pg Postgres) BeginSerializable(ctx context.Context) (context.Context, error) {
	ctx, span := pg.tracer.Start(ctx, "Postgres.BeginSerializable")
	defer span.End()

	tx, err := pg.pool.BeginTx(ctx, pgx.TxOptions{
		IsoLevel:   pgx.Serializable,
		AccessMode: pgx.ReadWrite,
	})
	if err != nil {
		return nil, err
	}

	ctx = context.WithValue(ctx, txKey, tx)
	return ctx, nil
}

// CommitTx commits the transaction.
func (pg Postgres) CommitTx(ctx context.Context) error {
	ctx, span := pg.tracer.Start(ctx, "Postgres.CommitTx")
	defer span.End()

	tx, ok := ctx.Value(txKey).(pgx.Tx)
	if !ok {
		return errors.New("transaction not found in context")
	}

	return tx.Commit(ctx)
}

// RollbackTx rolls back the transaction.
func (pg Postgres) RollbackTx(ctx context.Context) error {
	ctx, span := pg.tracer.Start(ctx, "Postgres.RollbackTx")
	defer span.End()

	tx, ok := ctx.Value(txKey).(pgx.Tx)
	if !ok {
		return errors.New("transaction not found in context")
	}

	return tx.Rollback(ctx)
}

// Query executes a query that returns a single row.
func (pg Postgres) Query(ctx context.Context, dest any, query string, args ...any) error {
	ctx, span := pg.tracer.Start(
		ctx,
		"Postgres.Query",
		trace.WithAttributes(attribute.String("query", query)),
	)
	defer span.End()
	return pgxscan.Get(ctx, pg.pool, dest, query, args...)
}

// QuerySlice executes a query that returns multiple rows.
func (pg Postgres) QuerySlice(ctx context.Context, dest any, query string, args ...any) error {
	ctx, span := pg.tracer.Start(
		ctx,
		"Postgres.QuerySlice",
		trace.WithAttributes(attribute.String("query", query)),
	)
	defer span.End()
	return pgxscan.Select(ctx, pg.pool, dest, query, args...)
}

// Exec executes a query that doesn't return any rows.
func (pg Postgres) Exec(ctx context.Context, query string, args ...any) (pgconn.CommandTag, error) {
	ctx, span := pg.tracer.Start(
		ctx,
		"Postgres.Exec",
		trace.WithAttributes(attribute.String("query", query)),
	)
	defer span.End()

	tag, err := pg.pool.Exec(ctx, query, args...)
	return tag, err
}

// QueryTx executes a query that returns a single row in a transaction.
func (pg Postgres) QueryTx(ctx context.Context, dest any, query string, args ...any) error {
	ctx, span := pg.tracer.Start(
		ctx,
		"Postgres.QueryTx",
		trace.WithAttributes(attribute.String("query", query)),
	)
	defer span.End()

	tx, err := getTxFromCtx(ctx)
	if err != nil {
		return err
	}

	return pgxscan.Get(ctx, tx, dest, query, args...)
}

// QuerySliceTx executes a query that returns multiple rows in a transaction.
func (pg Postgres) QuerySliceTx(ctx context.Context, dest any, query string, args ...any) error { //nolint:dupl // because of pgx api
	ctx, span := pg.tracer.Start(
		ctx,
		"Postgres.QuerySliceTx",
		trace.WithAttributes(attribute.String("query", query)),
	)
	defer span.End()

	tx, err := getTxFromCtx(ctx)
	if err != nil {
		return err
	}

	return pgxscan.Select(ctx, tx, dest, query, args...)
}

// ExecTx executes a query that doesn't return any rows in a transaction.
func (pg Postgres) ExecTx(ctx context.Context, query string, args ...any) (pgconn.CommandTag, error) {
	ctx, span := pg.tracer.Start(
		ctx,
		"Postgres.ExecTx",
		trace.WithAttributes(attribute.String("query", query)),
	)
	defer span.End()

	tx, err := getTxFromCtx(ctx)
	if err != nil {
		return pgconn.CommandTag{}, err
	}

	tag, err := tx.Exec(ctx, query, args...)
	return tag, err
}

func getTxFromCtx(ctx context.Context) (pgx.Tx, error) {
	tx, ok := ctx.Value(txKey).(pgx.Tx)
	if !ok {
		return nil, errors.New("transaction not found in context")
	}
	return tx, nil
}
