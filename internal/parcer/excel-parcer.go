package parcer

import (
	"fmt"

	"github.com/360EntSecGroup-Skylar/excelize/v2"
)

type Parser interface {
	ParseFile(filepath string) (data *ParsedData, err error)
}

type ExcelDataParser struct {
	castersChain Caster
}

func (p *ExcelDataParser) ParseFile(filepath string) (data *ParsedData, err error) {
	table, err := excelize.OpenFile(filepath)
	if err != nil {
		fmt.Println(err)
		return data, err
	}

	data = new(ParsedData)
	// TODO : parallel !!!!!
	for _, sheetName := range table.GetSheetMap() {
		rows, err := table.Rows(sheetName)
		if err != nil {
			return data, err
		}

		for rows.Next() {
			row, err := rows.Columns()
			if err != nil {
				return data, err
			}

			stored := new(StoredData)
			if err = p.castersChain.Cast(stored, row); err != nil {
				data.ParsingFailure++
				continue
			}
			data.Data = append(data.Data, stored)
		}

	}

	return
}

func NewExcelDataParser(castersChain Caster) *ExcelDataParser {
	return &ExcelDataParser{castersChain: castersChain}
}
