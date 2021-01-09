package casters

import (
	"project-afina/internal/parcer"
)


type OfferCuster struct {
	next parcer.Caster
}

func (c *OfferCuster) Cast(d *parcer.StoredData, parsedData []string) (err error) {
	d.OfferID = parsedData[0]
	return c.next.Cast(d, parsedData[1:])
}

func (c *OfferCuster) Next(caster parcer.Caster) {
	c.next = caster
}

func NewOfferCuster() *OfferCuster {
	return &OfferCuster{}
}
