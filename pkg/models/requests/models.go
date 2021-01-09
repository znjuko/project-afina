package requests

import (
	"project-afina/pkg/models/storage"
)

type GetOffersRequest struct {
	SellerID string `json:"seller_id,omitempty"`
	OfferID  string `json:"offer_id,omitempty"`
	Search   string `json:"search,omitempty"`
}

type GetOffersResponse struct {
	Offers storage.StoredData `json:"offers,omitempty"`
}

type UpsertOffersRequest struct {
	SellerID     string `json:"seller_id,omitempty"`
	DownloadLink string `json:"download_link,omitempty"`
}

type UpsertOffersResponse struct {
	Created int `json:"created,omitempty"`
	Updated int `json:"updated,omitempty"`
	Deleted int `json:"deleted,omitempty"`
	Failed  int `json:"failed,omitempty"`
}
