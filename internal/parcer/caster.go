package parcer

type Caster interface {
	Cast(d *StoredData, parsedData []string) (err error)
	Next(caster Caster)
}
