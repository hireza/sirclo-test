package integration_tests

import (
	"database/sql"

	"github.com/hireza/sirclo-test/berat/packages"
	"github.com/rs/zerolog/log"

	"github.com/hireza/sirclo-test/berat/app/repository"
	"github.com/hireza/sirclo-test/berat/domain"
)

const (
	queryTruncateWeights = "TRUNCATE TABLE weights;"
	queryCheckByDate     = `select exists(SELECT 1 FROM weights WHERE deleted_at IS NULL and date = $1 );`
	queryGetAll          = `SELECT id, date, max, min FROM weights WHERE deleted_at IS NULL;`
	queryGetByDate       = `SELECT id, date, max, min FROM weights WHERE deleted_at IS NULL AND date = $1;`
	queryCreate          = `INSERT INTO weights (date, max, min) VALUES ($1, $2, $3);`
	queryUpdate          = `UPDATE weights SET max = $2, min = $3, updated_at = NOW() WHERE deleted_at IS NULL AND date = $1;`
	queryDelete          = `UPDATE weights SET deleted_at = NOW() WHERE deleted_at IS NULL AND date = $1;`
)

func initPackages() (packages.Packages, error) {
	mgr, err := packages.NewInit("../../../packages/config", "127.0.0.1:5432")
	if err != nil {
		log.Error().Err(err).Msg("")
		return nil, err
	}

	return mgr, nil
}

func truncateWeights(db *sql.DB) error {
	_, err := db.Exec(queryTruncateWeights)
	if err != nil {
		log.Error().Err(err).Msg("")
	}

	return nil
}

func seedWeight(db *sql.DB) (*domain.Weight, error) {
	data := &domain.Weight{
		Date: "1997-06-15",
		Max:  13,
		Min:  3,
	}

	weight := &domain.Weight{}

	_, err := db.Exec(queryCreate, data.Date, data.Max, data.Min)
	if err != nil {
		log.Error().Err(err).Msg("")
		return nil, err
	}

	err = db.QueryRow(queryGetByDate, data.Date).Scan(&weight.ID, &weight.Date, &weight.Max, &weight.Min)
	if err != nil {
		log.Error().Err(err).Msg("")
		return nil, err
	}

	weight.GetDifferent()
	err = weight.GetIndoFormattedDate()
	if err != nil {
		log.Error().Err(err).Msg("")
		return nil, err
	}

	return weight, nil
}

func seedWeights(db *sql.DB) ([]*domain.Weight, error) {
	datas := []*domain.Weight{
		{
			Date: "1997-06-15",
			Max:  10,
			Min:  5,
		},
		{
			Date: "1997-06-16",
			Max:  100,
			Min:  30,
		},
	}

	weights := []*domain.Weight{}

	for _, v := range datas {
		_, err := db.Exec(queryCreate, v.Date, v.Max, v.Min)
		if err != nil {
			log.Error().Err(err).Msg("")
			return nil, err
		}
	}

	rows, err := db.Query(queryGetAll)
	if err != nil {
		log.Error().Err(err).Msg("")
		return nil, err
	}

	for rows.Next() {
		weight := &domain.Weight{}

		err := rows.Scan(&weight.ID, &weight.Date, &weight.Max, &weight.Min)
		if err != nil {
			log.Error().Err(err).Msg("")
			return nil, err
		}

		weight.GetDifferent()
		err = weight.GetIndoFormattedDate()
		if err != nil {
			log.Error().Err(err).Msg("")
			return nil, err
		}

		weights = append(weights, weight)
	}

	weights = repository.GetAllFormatted(weights)
	return weights, nil
}
