package psql

import (
	"context"
	"encoding/json"

	"github.com/doug-martin/goqu"

	"project-afina/pkg/models/storage"
)

type commonStorage interface {
	StoreData(ctx context.Context, req *storage.Request) (err error)
	GetData(ctx context.Context, req *storage.Request) (data [][]byte, err error)
	DeleteData(ctx context.Context, reqs []*storage.Request) (err error)
}

type Storage struct {
	common commonStorage

	storedTable string
}

func (s *Storage) SelectOfferData(ctx context.Context, conf []func(expr goqu.Ex)) (data []*storage.StoredData, err error) {
	ex := goqu.Ex{}
	for _, f := range conf {
		f(ex)
	}

	sql, params, err := goqu.From(s.storedTable).Select(
		&storage.StoredData{},
	).Where(
		ex,
	).ToSQL()
	if err != nil {
		return data, err
	}

	byteData, err := s.common.GetData(ctx, &storage.Request{
		Expression: sql,
		Params:     params,
	})
	if err != nil {
		return data, err
	}

	data = make([]*storage.StoredData, len(byteData))
	for iter := range byteData {
		if err = json.Unmarshal(byteData[iter], data[iter]); err != nil {
			return data, err
		}
	}

	return
}

func (s *Storage) StoreOfferData(
	ctx context.Context, seller_id string, data []*storage.StoredData,
) (diff *storage.ChangedData, err error) {
	delQ := goqu.From(s.storedTable)
	upsQ := goqu.From(s.storedTable)

	return
}
