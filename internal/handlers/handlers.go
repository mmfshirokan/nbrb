package handlers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/mmfshirokan/nbrb/internal/model"
)

type Geter interface {
	Get(date time.Time) ([]model.Currency, error)
	GetAll() ([]model.Currency, error)
}

type Handlers struct {
	sv Geter
}

func New(sv Geter) *Handlers {
	return &Handlers{sv: sv}
}

func (h *Handlers) Get(w http.ResponseWriter, r *http.Request) {
	date, err := time.Parse(time.DateOnly, r.URL.Query().Get("date"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	crs, err := h.sv.Get(date)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := json.NewEncoder(w).Encode(crs); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (h *Handlers) GetAll(w http.ResponseWriter, r *http.Request) {
	crs, err := h.sv.GetAll()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := json.NewEncoder(w).Encode(crs); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
