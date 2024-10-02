package service

import (
	"time"

	"github.com/mmfshirokan/nbrb/internal/model"
)

type RpController interface {
	Add(crs []model.Currency) error
	Get(date time.Time) ([]model.Currency, error)
	GetAll() ([]model.Currency, error)
}

type Service struct {
	rp RpController
}

func New(rp RpController) *Service {
	return &Service{
		rp: rp,
	}
}

func (s *Service) Add(crs []model.Currency) error {
	return s.rp.Add(crs)
}

func (s *Service) Get(date time.Time) ([]model.Currency, error) {
	return s.rp.Get(date)
}

func (s *Service) GetAll() ([]model.Currency, error) {
	return s.rp.GetAll()
}
