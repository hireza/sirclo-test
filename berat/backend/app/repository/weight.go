package repository

import (
	"context"
	"database/sql"
	"sort"
	"time"

	"github.com/hireza/sirclo-test/berat/domain"
	"github.com/hireza/sirclo-test/berat/packages"
	"github.com/rs/zerolog/log"
)

type Weight struct {
	DB *sql.DB
}

func NewWeightRepository(mgr packages.Packages) domain.WeightRepository {
	repository := new(Weight)
	repository.DB = mgr.GetPostgres()

	return repository
}

func (r *Weight) CheckByDate(ctx context.Context, date string) (bool, error) {
	var exist bool

	query := `
		select exists(SELECT 1 FROM weights WHERE deleted_at IS NULL and date = $1 );
	`

	err := r.DB.QueryRowContext(ctx, query, date).Scan(&exist)
	if err != nil {
		log.Error().Err(err).Msg("[REPO] error when check weight")
		return false, err
	}

	return exist, nil
}

func GetAllFormatted(datas []*domain.Weight) []*domain.Weight {
	layout := "02-01-2006"

	if len(datas) > 0 {
		sort.Slice(datas, func(i, j int) bool {
			time1, _ := time.Parse(layout, datas[i].Date)
			time2, _ := time.Parse(layout, datas[j].Date)

			return time1.After(time2)
		})

		maxMean := 0.0
		minMean := 0.0

		for _, v := range datas {
			maxMean += float64(v.Max)
			minMean += float64(v.Min)
		}

		max := maxMean / float64(len(datas))
		min := minMean / float64(len(datas))

		weight := &domain.Weight{
			Date:      "Rata-rata",
			Max:       max,
			Min:       min,
			Different: max - min,
		}

		datas = append(datas, weight)
	}

	return datas
}

func (r *Weight) GetAll(ctx context.Context) ([]*domain.Weight, error) {
	weights := []*domain.Weight{}

	query := `
		SELECT id, date, max, min FROM weights WHERE deleted_at IS NULL;
	`

	rows, err := r.DB.QueryContext(ctx, query)
	if err != nil {
		log.Error().Err(err).Msg("[REPO] error when get all weight")
		return nil, err
	}

	for rows.Next() {
		weight := &domain.Weight{}

		err := rows.Scan(&weight.ID, &weight.Date, &weight.Max, &weight.Min)
		if err != nil {
			log.Error().Err(err).Msg("[REPO] error when scan all weight")
			return nil, err
		}

		weight.GetDifferent()
		err = weight.GetIndoFormattedDate()
		if err != nil {
			log.Error().Err(err).Msg("[REPO] error when format date all weight")
			return nil, err
		}

		weights = append(weights, weight)
	}

	weights = GetAllFormatted(weights)
	return weights, nil
}

func (r *Weight) GetByDate(ctx context.Context, date string) (*domain.Weight, error) {
	weight := &domain.Weight{}

	query := `
		SELECT id, date, max, min FROM weights WHERE deleted_at IS NULL AND date = $1;
	`

	err := r.DB.QueryRowContext(ctx, query, date).Scan(&weight.ID, &weight.Date, &weight.Max, &weight.Min)
	if err != nil {
		log.Error().Err(err).Msg("[REPO] error when get & scan weight")
		return nil, err
	}

	weight.GetDifferent()
	err = weight.GetIndoFormattedDate()
	if err != nil {
		log.Error().Err(err).Msg("[REPO] error when format date weight")
		return nil, err
	}

	return weight, nil
}

func (r *Weight) Create(ctx context.Context, data *domain.Weight) error {
	query := `
		INSERT INTO weights (date, max, min) VALUES ($1, $2, $3);
	`

	_, err := r.DB.ExecContext(ctx, query, data.Date, data.Max, data.Min)
	if err != nil {
		log.Error().Err(err).Msg("[REPO] error when create weight")
		return err
	}

	return nil
}

func (r *Weight) Update(ctx context.Context, date string, data *domain.Weight) error {
	query := `
		UPDATE weights SET max = $2, min = $3, updated_at = NOW()
		WHERE deleted_at IS NULL AND date = $1;
	`

	_, err := r.DB.ExecContext(ctx, query, data.Date, data.Max, data.Min)
	if err != nil {
		log.Error().Err(err).Msg("[REPO] error when update weight")
		return err
	}

	return nil
}

func (r *Weight) Delete(ctx context.Context, date string) error {
	query := `
		UPDATE weights SET deleted_at = NOW()
		WHERE deleted_at IS NULL AND date = $1;
	`

	_, err := r.DB.ExecContext(ctx, query, date)
	if err != nil {
		log.Error().Err(err).Msg("[REPO] error when delete weight")
		return err
	}

	return nil
}
