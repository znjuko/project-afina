package casters

import (
	"project-afina/internal/parcer"
	"project-afina/pkg/models/excel-data"
)

type NameCuster struct {
	next parcer.Caster
}

func (c *NameCuster) Cast(d *excel_data.StoredData, parsedData []string) (err error) {
	d.OfferID = parsedData[0]
	return c.next.Cast(d, parsedData[1:])
}

func (c *NameCuster) Next(caster parcer.Caster) {
	c.next = caster
}

func NewNameCuster() *NameCuster {
	return &NameCuster{}
}
