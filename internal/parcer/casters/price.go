package casters

import (
	"strconv"

	"project-afina/internal/parcer"
	"project-afina/pkg/models/excel-data"
)

type PriceCuster struct {
	next parcer.Caster
}

func (c *PriceCuster) Cast(d *excel_data.StoredData, parsedData []string) (err error) {
	if d.Price, err = strconv.ParseFloat(parsedData[0], 64); err != nil {
		return
	}

	return c.next.Cast(d, parsedData[1:])
}

func (c *PriceCuster) Next(caster parcer.Caster) {
	c.next = caster
}

func NewPriceCuster() *PriceCuster {
	return &PriceCuster{}
}
