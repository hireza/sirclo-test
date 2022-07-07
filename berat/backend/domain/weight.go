package domain

import (
	"context"
	"errors"
	"net/http"

	timeHelper "github.com/hireza/sirclo-test/berat/helper/universal"
)

type WeightRepository interface {
	CheckByDate(ctx context.Context, date string) (bool, error)
	GetAll(ctx context.Context) ([]*Weight, error)
	GetByDate(ctx context.Context, date string) (*Weight, error)
	Create(ctx context.Context, data *Weight) error
	Update(ctx context.Context, date string, data *Weight) error
	Delete(ctx context.Context, date string) error
}

type WeightUsecase interface {
	GetAll(ctx context.Context) ([]*Weight, error)
	GetByDate(ctx context.Context, date string) (*Weight, error)
	Create(ctx context.Context, data *Weight) ([]*Weight, error)
	Update(ctx context.Context, date string, data *Weight) ([]*Weight, error)
	Delete(ctx context.Context, date string) ([]*Weight, error)
}

type WeightHandler interface {
	GetAll() http.Handler
	GetByDate() http.Handler
	Create() http.Handler
	Update() http.Handler
	Delete() http.Handler
}

type Weight struct {
	ID        string  `json:"-"`
	Date      string  `json:"date"`
	Max       float64 `json:"max"`
	Min       float64 `json:"min"`
	Different float64 `json:"different"`
}

func (w *Weight) Validate() error {
	// validate date with format DD-MM-YYYY
	timeFormat, err := timeHelper.GetGlobalFormattedDate(w.Date)
	if err != nil {
		return err
	}
	w.Date = timeFormat

	// validate if max > min
	if w.Max < w.Min {
		return errors.New("maximum weight cannot lower than minimum weight")
	}

	// get different value
	w.GetDifferent()

	return nil
}

func (w *Weight) GetDifferent() {
	w.Different = w.Max - w.Min
}

func (w *Weight) GetIndoFormattedDate() error {
	timeFormat, err := timeHelper.GetIndoFormattedDate(w.Date)
	if err != nil {
		return err
	}

	w.Date = timeFormat
	return nil
}
