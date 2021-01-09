package common

import (
	"context"
	"errors"

	pgx "github.com/jackc/pgx/v4/pgxpool"

	"project-afina/pkg/models/storage"
)

type Psql struct {
	connPool pgx.Pool
}

func (s *Psql) StoreData(ctx context.Context, req *storage.Request) (err error) {
	if req == nil {
		return errors.New("request cannot be empty")
	}

	conn, err := s.connPool.Acquire(ctx)
	if err != nil {
		return err
	}

	tx, err := conn.Begin(ctx)
	if err != nil {
		return err
	}
	defer func() {
		if err != nil {
			tx.Rollback(ctx)
		} else {
			tx.Commit(ctx)
		}
	}()

	if _, err = tx.Exec(ctx, req.Expression, req.Params); err != nil {
		return
	}

	return
}

func (s *Psql) GetData(
	ctx context.Context, req *storage.Request,
) (data [][]byte, err error) {
	if req == nil {
		return data, errors.New("request cannot be empty")
	}

	conn, err := s.connPool.Acquire(ctx)
	if err != nil {
		return data, err
	}

	tx, err := conn.Begin(ctx)
	if err != nil {
		return data, err
	}
	defer func() {
		if err != nil {
			tx.Rollback(ctx)
		} else {
			tx.Commit(ctx)
		}
	}()

	rows, err := tx.Query(ctx, req.Expression, req.Params)
	if err != nil {
		return data, err
	}

	return rows.RawValues(), nil
}

func (s *Psql) DeleteData(ctx context.Context, reqs []*storage.Request) (err error) {
	conn, err := s.connPool.Acquire(ctx)
	if err != nil {
		return err
	}

	tx, err := conn.Begin(ctx)
	if err != nil {
		return err
	}
	defer func() {
		if err != nil {
			tx.Rollback(ctx)
		} else {
			tx.Commit(ctx)
		}
	}()

	for _, req := range reqs {
		if req == nil {
			continue
		}

		if _, err = tx.Exec(ctx, req.Expression, req.Params); err != nil {
			return err
		}
	}

	return
}

func NewStorage(connPool pgx.Pool) *Psql {
	return &Psql{connPool: connPool}
}
