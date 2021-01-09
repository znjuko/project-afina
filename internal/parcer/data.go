package parcer

type StoredData struct {
	OfferID   string
	Name      string
	Price     float64
	Quantity  int64
	Available bool
}

type ParsedData struct {
	Data           []*StoredData
	ParsingFailure int32
}
