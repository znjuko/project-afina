package casters

import (
	"strconv"

	"project-afina/internal/parcer"
)

type AvailableCuster struct {
	next parcer.Caster
}

func (c *AvailableCuster) Cast(d *parcer.StoredData, parsedData []string) (err error) {
	if d.Available, err = strconv.ParseBool(parsedData[0]); err != nil {
		return
	}

	return
}

func (c *AvailableCuster) Next(caster parcer.Caster) {
	c.next = caster
}

func NewAvailableCuster() *AvailableCuster{
	return &AvailableCuster{}
}
