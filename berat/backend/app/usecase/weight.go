package usecase

import (
	"context"

	"github.com/hireza/sirclo-test/berat/app/repository"
	"github.com/hireza/sirclo-test/berat/domain"
	"github.com/hireza/sirclo-test/berat/packages"
	"github.com/rs/zerolog/log"
)

type Weight struct {
	WeightRepository domain.WeightRepository
}

func NewWeightUsecase(mgr packages.Packages) domain.WeightUsecase {
	usecase := new(Weight)
	usecase.WeightRepository = repository.NewWeightRepository(mgr)

	return usecase
}

func (u *Weight) GetAll(ctx context.Context) ([]*domain.Weight, error) {
	weights, err := u.WeightRepository.GetAll(ctx)
	if err != nil {
		log.Error().Err(err).Msg("[UC] error when get all weight")
		return nil, err
	}

	return weights, nil
}

func (u *Weight) GetByDate(ctx context.Context, date string) (*domain.Weight, error) {
	weight, err := u.WeightRepository.GetByDate(ctx, date)
	if err != nil {
		log.Error().Err(err).Msg("[UC] error when get weight")
		return nil, err
	}

	return weight, nil
}

func (u *Weight) Create(ctx context.Context, data *domain.Weight) ([]*domain.Weight, error) {
	exist, err := u.WeightRepository.CheckByDate(ctx, data.Date)
	if err != nil {
		log.Error().Err(err).Msg("[UC] error when check weight")
		return nil, err
	}

	if exist {
		err := u.WeightRepository.Update(ctx, data.Date, data)
		if err != nil {
			log.Error().Err(err).Msg("[UC] error when update weight after check")
			return nil, err
		}
	} else {
		err := u.WeightRepository.Create(ctx, data)
		if err != nil {
			log.Error().Err(err).Msg("[UC] error when create weight after check")
			return nil, err
		}
	}

	weights, err := u.WeightRepository.GetAll(ctx)
	if err != nil {
		log.Error().Err(err).Msg("[UC] error when get all weight after create/update")
		return nil, err
	}

	return weights, nil
}

func (u *Weight) Update(ctx context.Context, date string, data *domain.Weight) ([]*domain.Weight, error) {
	err := u.WeightRepository.Update(ctx, date, data)
	if err != nil {
		log.Error().Err(err).Msg("[UC] error when update weight")
		return nil, err
	}

	weights, err := u.WeightRepository.GetAll(ctx)
	if err != nil {
		log.Error().Err(err).Msg("[UC] error when get all weight after update")
		return nil, err
	}

	return weights, nil
}

func (u *Weight) Delete(ctx context.Context, date string) ([]*domain.Weight, error) {
	err := u.WeightRepository.Delete(ctx, date)
	if err != nil {
		log.Error().Err(err).Msg("[UC] error when delete weight")
		return nil, err
	}

	weights, err := u.WeightRepository.GetAll(ctx)
	if err != nil {
		log.Error().Err(err).Msg("[UC] error when get all weight after delete")
		return nil, err
	}

	return weights, nil
}
