package handler

import (
	"errors"
	"net/http"
	"strings"

	"github.com/hireza/sirclo-test/berat/app/usecase"
	"github.com/hireza/sirclo-test/berat/domain"
	"github.com/rs/zerolog/log"

	jsonHelper "github.com/hireza/sirclo-test/berat/helper/json"
	timeHelper "github.com/hireza/sirclo-test/berat/helper/universal"
	"github.com/hireza/sirclo-test/berat/packages"

	"encoding/json"
)

type Weight struct {
	WeightUsecase domain.WeightUsecase
}

func NewWeightHandler(mgr packages.Packages) domain.WeightHandler {
	handler := new(Weight)
	handler.WeightUsecase = usecase.NewWeightUsecase(mgr)

	return handler
}

func (h *Weight) GetAll() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		weights, err := h.WeightUsecase.GetAll(ctx)
		if err != nil {
			log.Error().Err(err).Msg("[HDL] error when get all weight")
			jsonHelper.Response(w, r, http.StatusInternalServerError, err.Error(), nil)
			return
		}

		jsonHelper.Response(w, r, http.StatusOK, "success get all weight", weights)
	})
}

func (h *Weight) GetByDate() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		date := r.FormValue("date")
		timeFormat, err := timeHelper.GetGlobalFormattedDate(date)
		if err != nil {
			log.Error().Err(err).Msg("[HDL] error when formatting to global date")
			jsonHelper.Response(w, r, http.StatusBadRequest, err.Error(), nil)
			return
		}
		newDate := timeFormat

		weight, err := h.WeightUsecase.GetByDate(ctx, newDate)
		if err != nil {
			log.Error().Err(err).Msg("[HDL] error when get weight")
			if strings.Contains(err.Error(), "no rows") {
				err = errors.New("failed get weight - not found")
				jsonHelper.Response(w, r, http.StatusNotFound, err.Error(), weight)
				return
			}

			jsonHelper.Response(w, r, http.StatusInternalServerError, err.Error(), nil)
			return
		}

		jsonHelper.Response(w, r, http.StatusOK, "success get weight", weight)
	})
}

func (h *Weight) Create() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		var data *domain.Weight
		if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
			log.Error().Err(err).Msg("[HDL] error when decode body before create")
			jsonHelper.Response(w, r, http.StatusBadRequest, err.Error(), nil)
			return
		}

		err := data.Validate()
		if err != nil {
			log.Error().Err(err).Msg("[HDL] error validate weight")
			jsonHelper.Response(w, r, http.StatusBadRequest, err.Error(), nil)
			return
		}

		weights, err := h.WeightUsecase.Create(ctx, data)
		if err != nil {
			log.Error().Err(err).Msg("[HDL] error when create weight")
			jsonHelper.Response(w, r, http.StatusInternalServerError, err.Error(), nil)
			return
		}

		jsonHelper.Response(w, r, http.StatusCreated, "success create weight", weights)
	})
}

func (h *Weight) Update() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		date := r.FormValue("date")
		timeFormat, err := timeHelper.GetGlobalFormattedDate(date)
		if err != nil {
			log.Error().Err(err).Msg("[HDL] error when formatting to global date")
			jsonHelper.Response(w, r, http.StatusBadRequest, err.Error(), nil)
			return
		}
		newDate := timeFormat

		var data *domain.Weight
		if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
			log.Error().Err(err).Msg("[HDL] error when decode body before update")
			jsonHelper.Response(w, r, http.StatusBadRequest, err.Error(), nil)
			return
		}
		data.Date = date

		err = data.Validate()
		if err != nil {
			log.Error().Err(err).Msg("[HDL] error validate weight")
			jsonHelper.Response(w, r, http.StatusBadRequest, err.Error(), nil)
			return
		}

		weights, err := h.WeightUsecase.Update(ctx, newDate, data)
		if err != nil {
			log.Error().Err(err).Msg("[HDL] error when update weight")
			jsonHelper.Response(w, r, http.StatusInternalServerError, err.Error(), nil)
			return
		}

		jsonHelper.Response(w, r, http.StatusCreated, "success update weight", weights)
	})
}

func (h *Weight) Delete() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		date := r.FormValue("date")
		timeFormat, err := timeHelper.GetGlobalFormattedDate(date)
		if err != nil {
			log.Error().Err(err).Msg("[HDL] error when formatting to global date")
			jsonHelper.Response(w, r, http.StatusBadRequest, err.Error(), nil)
			return
		}
		newDate := timeFormat

		weights, err := h.WeightUsecase.Delete(ctx, newDate)
		if err != nil {
			log.Error().Err(err).Msg("[HDL] error when delete weight")
			jsonHelper.Response(w, r, http.StatusInternalServerError, err.Error(), nil)
			return
		}

		jsonHelper.Response(w, r, http.StatusAccepted, "success delete weight", weights)
	})
}
