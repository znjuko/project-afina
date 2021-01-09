package parcer

import "project-afina/pkg/models/excel-data"

type Caster interface {
	Cast(d *excel_data.StoredData, parsedData []string) (err error)
	Next(caster Caster)
}
