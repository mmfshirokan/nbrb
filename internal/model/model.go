package model

import (
	"github.com/shopspring/decimal"
)

type Currency struct {
	ID           int             `json:"Cur_ID"`
	Date         string          `json:"Date"`
	Abbreviation string          `json:"Cur_Abbreviation"`
	Scale        int             `json:"Cur_Scale"`
	Name         string          `json:"Cur_Name"`
	OfficialRate decimal.Decimal `json:"Cur_OfficialRate"`
}
