package repository

import (
	"database/sql"
	"time"

	"github.com/mmfshirokan/nbrb/internal/model"
)

type MySQl struct {
	db *sql.DB
}

func New(db *sql.DB) *MySQl {
	return &MySQl{
		db: db,
	}
}

func (m *MySQl) Add(crs []model.Currency) error {
	for _, cr := range crs {
		_, err := m.db.Exec("INSERT INTO currency(id, cur_date, abbreviation, scale, name, officialRate) VALUES (?, ?, ?, ?, ?, ?)", cr.ID, cr.Date, cr.Abbreviation, cr.Scale, cr.Name, cr.OfficialRate.String())
		if err != nil {
			return err
		}
	}

	return nil
}

func (m *MySQl) Get(date time.Time) ([]model.Currency, error) {
	day := date.Format(time.DateOnly)

	rows, err := m.db.Query("SELECT id, cur_date, abbreviation, scale, name, officialRate FROM currency WHERE ? >= cur_date AND cur_date >= ?", day+" 23:59:59", day+" 00:00:00")
	if err != nil {
		return nil, err
	}

	result := make([]model.Currency, 0, 32)
	for rows.Next() {
		var cr model.Currency

		err = rows.Scan(&cr.ID, &cr.Date, &cr.Abbreviation, &cr.Scale, &cr.Name, &cr.OfficialRate)
		if err != nil {
			return nil, err
		}

		result = append(result, cr)
	}

	return result, nil
}

func (m *MySQl) GetAll() ([]model.Currency, error) {
	rows, err := m.db.Query("SELECT id, cur_date, abbreviation, scale, name, officialRate FROM currency")
	if err != nil {
		return nil, err
	}

	result := make([]model.Currency, 0, 32)
	for rows.Next() {
		var cr model.Currency

		err = rows.Scan(&cr.ID, &cr.Date, &cr.Abbreviation, &cr.Scale, &cr.Name, &cr.OfficialRate)
		if err != nil {
			return nil, err
		}

		result = append(result, cr)
	}

	return result, nil
}
