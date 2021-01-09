package transormer

import (
	"errors"

	"github.com/doug-martin/goqu"

	"project-afina/pkg/models/storage"
)

type StorageTransformer struct{}

func (t *StorageTransformer) TransformToDeleteRequest(seller_id string, model *storage.StoredData) (expr goqu.Ex, err error) {
	if model == nil {
		return expr, errors.New("stored data cannot be nil")
	}
	expr = goqu.Ex{}

	expr["seller_id"] = seller_id

	if model.OfferID != "" {
		expr["offer_id"] = model.OfferID
	}

	return
}
