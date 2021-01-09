package storage

type Request struct {
	Expression string
	Params     []interface{}
}

type StoredData struct {
	SellerID  string  `json:"seller_id,omitempty" db:"seller_id,omitempty"`
	OfferID   string  `json:"offer_id,omitempty" db:"offer_id,omitempty"`
	Name      string  `json:"name,omitempty" db:"name,omitempty"`
	Price     float64 `json:"price,omitempty" db:"price,omitempty"`
	Quantity  int64   `json:"quantity,omitempty" db:"quantity,omitempty"`
	Available bool    `json:"available,omitempty" db:"available,omitempty"`
}

type ChangedData struct {
	Deleted int `json:"???,omitempty" db:"???,omitempty"`
	Updated int `json:"???,omitempty" db:"???,omitempty"`
	Created int `json:"???,omitempty" db:"???,omitempty"`
}
