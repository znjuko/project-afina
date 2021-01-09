package usecase

import (
	"context"

	"github.com/doug-martin/goqu"

	"project-afina/pkg/models/requests"
)

type transformer interface {
	TransformFromGetRequest(req *requests.GetOffersRequest) (conf []func(expr goqu.Ex))
}

type Offer struct {
	dataTransformer
}

func (o *Offer) GetOfferData(
	ctx context.Context, req *requests.GetOffersRequest,
) (resp *requests.GetOffersResponse, err error) {

	return
}
