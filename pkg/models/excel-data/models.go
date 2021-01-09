package excel_data

import "project-afina/pkg/models/storage"

type ParsedData struct {
	Data           []*storage.StoredData
	ParsingFailure int32
}
