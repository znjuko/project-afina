package casters

import (
	"project-afina/internal/parcer"

	"errors"
)

type ColumnsChecker struct {
	minColumns int
	next       parcer.Caster
}

func (c *ColumnsChecker) Cast(d *parcer.StoredData, parsedData []string) (err error) {
	if len(parsedData) < c.minColumns {
		return errors.New("this rows is too short")
	}

	return c.next.Cast(d, parsedData)
}

func (c *ColumnsChecker) Next(caster parcer.Caster) {
	c.next = caster
}

func NewColumnsChecker(minColumns int) *ColumnsChecker {
	return &ColumnsChecker{minColumns: minColumns}
}
