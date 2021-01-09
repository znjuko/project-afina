package casters

import (
	"strconv"

	"project-afina/internal/parcer"
	"project-afina/pkg/models/excel-data"
)

type QuantityCuster struct {
	next parcer.Caster
}

func (c *QuantityCuster) Cast(d *excel_data.StoredData, parsedData []string) (err error) {
	if d.Quantity, err = strconv.ParseInt(parsedData[0], 10, 64); err != nil {
		return
	}

	return c.next.Cast(d, parsedData[1:])
}

func (c *QuantityCuster) Next(caster parcer.Caster) {
	c.next = caster
}

func NewQuantityCuster() *QuantityCuster {
	return &QuantityCuster{}
}
