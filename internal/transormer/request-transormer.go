package transormer

import (
	"github.com/doug-martin/goqu"

	"project-afina/pkg/models/requests"
)

type RequestTransformer struct{}

func (t *RequestTransformer) TransformFromGetRequest(req *requests.GetOffersRequest) (conf []func(expr goqu.Ex)) {
	conf = make([]func(expr goqu.Ex), 0)

	if req.SellerID != "" {
		conf = append(conf, func(expr goqu.Ex) {
			expr["seller_id"] = req.SellerID
		})
	}

	if req.OfferID != "" {
		conf = append(conf, func(expr goqu.Ex) {
			expr["offer_id"] = req.OfferID
		})
	}

	if req.Search != "" {
		conf = append(conf, func(expr goqu.Ex) {
			expr["name"] = goqu.Op{"like": req.Search}
		})
	}

	return
}

func NewRequestTransformer() *RequestTransformer {
	return &RequestTransformer{}
}
