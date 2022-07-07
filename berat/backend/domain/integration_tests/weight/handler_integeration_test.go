package integration_tests

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
	"github.com/hireza/sirclo-test/berat/app/delivery/rest/handler"
	jsonHelper "github.com/hireza/sirclo-test/berat/helper/json"
	"github.com/stretchr/testify/assert"
)

func TestWeight_GetAll(t *testing.T) {
	mgr, err := initPackages()
	if err != nil {
		log.Fatal(err)
	}

	hdl := handler.NewWeightHandler(mgr)

	err = truncateWeights(mgr.GetPostgres())
	if err != nil {
		log.Fatal(err)
	}

	results, err := seedWeights(mgr.GetPostgres())
	if err != nil {
		log.Fatal(err)
	}

	tests := []struct {
		name       string
		statusCode int
		message    string
		date       []string
		max        []float64
		min        []float64
		different  []float64
	}{
		{
			name:       "success",
			statusCode: http.StatusOK,
			message:    "success get all weight",
			date:       []string{results[0].Date, results[1].Date, "Rata-rata"},
			max:        []float64{results[0].Max, results[1].Max, 55},
			min:        []float64{results[0].Min, results[1].Min, 17.5},
			different:  []float64{results[0].Different, results[1].Different, 37.5},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			route := mux.NewRouter()
			route.Handle("/weights", hdl.GetAll()).Methods("GET")

			req, err := http.NewRequest(http.MethodGet, "/weights", nil)
			if err != nil {
				t.Error(err)
			}
			rr := httptest.NewRecorder()
			route.ServeHTTP(rr, req)

			responseMap := jsonHelper.Res{}
			err = json.Unmarshal(rr.Body.Bytes(), &responseMap)
			if err != nil {
				t.Errorf("Cannot convert to struct: %v", err)
			}

			assert.Equal(t, responseMap.Meta.StatusCode, tt.statusCode)
			assert.Equal(t, responseMap.Meta.Message, tt.message)

			responseData := responseMap.Data.([]interface{})
			for i := range tt.date {
				responseData := responseData[i]
				value := responseData.(map[string]interface{})
				assert.Equal(t, value["date"], tt.date[i])
				assert.Equal(t, value["max"], tt.max[i])
				assert.Equal(t, value["min"], tt.min[i])
				assert.Equal(t, value["different"], tt.different[i])
			}
		})
	}
}

func TestWeight_GetByDate(t *testing.T) {
	mgr, err := initPackages()
	if err != nil {
		log.Fatal(err)
	}

	hdl := handler.NewWeightHandler(mgr)

	err = truncateWeights(mgr.GetPostgres())
	if err != nil {
		log.Fatal(err)
	}

	result, err := seedWeight(mgr.GetPostgres())
	if err != nil {
		log.Fatal(err)
	}

	tests := []struct {
		name         string
		statusCode   int
		message      string
		insertedDate string
		date         string
		max          float64
		min          float64
		different    float64
	}{
		{
			name:         "success",
			statusCode:   http.StatusOK,
			message:      "success get weight",
			insertedDate: "15-06-1997",
			date:         result.Date,
			max:          result.Max,
			min:          result.Min,
			different:    result.Different,
		},
		{
			name:         "error format date",
			statusCode:   http.StatusBadRequest,
			message:      "date format must be dd-mm-yyyy",
			insertedDate: "15-13-1997",
		},
		{
			name:         "error getByData - not found",
			statusCode:   http.StatusNotFound,
			message:      "failed get weight - not found",
			insertedDate: "15-12-1997",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			route := mux.NewRouter()
			route.Handle("/weight", hdl.GetByDate()).Methods("GET")

			req, err := http.NewRequest(http.MethodGet, "/weight?date="+tt.insertedDate, nil)
			if err != nil {
				t.Error(err)
			}
			rr := httptest.NewRecorder()
			route.ServeHTTP(rr, req)

			responseMap := jsonHelper.Res{}
			err = json.Unmarshal(rr.Body.Bytes(), &responseMap)
			if err != nil {
				t.Errorf("Cannot convert to struct: %v", err)
			}

			assert.Equal(t, responseMap.Meta.StatusCode, tt.statusCode)
			assert.Equal(t, responseMap.Meta.Message, tt.message)

			if tt.statusCode == 200 {
				value := responseMap.Data.(map[string]interface{})
				assert.Equal(t, value["date"], tt.date)
				assert.Equal(t, value["max"], tt.max)
				assert.Equal(t, value["min"], tt.min)
				assert.Equal(t, value["different"], tt.different)
			}
		})
	}
}

func TestWeight_Create(t *testing.T) {
	mgr, err := initPackages()
	if err != nil {
		log.Fatal(err)
	}

	hdl := handler.NewWeightHandler(mgr)

	err = truncateWeights(mgr.GetPostgres())
	if err != nil {
		log.Fatal(err)
	}

	tests := []struct {
		name       string
		inputJSON  string
		statusCode int
		message    string
		date       []string
		max        []float64
		min        []float64
		different  []float64
	}{
		{
			name:       "success",
			inputJSON:  `{"date":"15-06-1997", "max": 10, "min": 4}`,
			statusCode: http.StatusCreated,
			message:    "success create weight",
			date:       []string{"15-06-1997", "Rata-rata"},
			max:        []float64{10, 10},
			min:        []float64{4, 4},
			different:  []float64{6, 6},
		},
		{
			name:       "error decode",
			inputJSON:  `{"date": "15-06-1997", "max": "abcd", "min": 5}`,
			statusCode: http.StatusBadRequest,
			message:    "json: cannot unmarshal string into Go struct field Weight.max of type float64",
		},
		{
			name:       "error validate",
			inputJSON:  `{"date": "15-06-1997", "max": 1, "min": 5}`,
			statusCode: http.StatusBadRequest,
			message:    "maximum weight cannot lower than minimum weight",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			route := mux.NewRouter()
			route.Handle("/weight", hdl.Create()).Methods("POST")

			req, err := http.NewRequest(http.MethodPost, "/weight", bytes.NewBufferString(tt.inputJSON))
			if err != nil {
				t.Error(err)
			}
			rr := httptest.NewRecorder()
			route.ServeHTTP(rr, req)

			responseMap := jsonHelper.Res{}
			err = json.Unmarshal(rr.Body.Bytes(), &responseMap)
			if err != nil {
				t.Errorf("Cannot convert to struct: %v", err)
			}

			assert.Equal(t, responseMap.Meta.StatusCode, tt.statusCode)
			assert.Equal(t, responseMap.Meta.Message, tt.message)

			if tt.statusCode == 201 {
				responseData := responseMap.Data.([]interface{})
				for i := range tt.date {
					responseData := responseData[i]
					value := responseData.(map[string]interface{})
					assert.Equal(t, value["date"], tt.date[i])
					assert.Equal(t, value["max"], tt.max[i])
					assert.Equal(t, value["min"], tt.min[i])
					assert.Equal(t, value["different"], tt.different[i])
				}
			}
		})
	}
}

func TestWeight_Update(t *testing.T) {
	mgr, err := initPackages()
	if err != nil {
		log.Fatal(err)
	}

	hdl := handler.NewWeightHandler(mgr)

	err = truncateWeights(mgr.GetPostgres())
	if err != nil {
		log.Fatal(err)
	}

	_, err = seedWeight(mgr.GetPostgres())
	if err != nil {
		log.Fatal(err)
	}

	tests := []struct {
		name         string
		inputJSON    string
		statusCode   int
		message      string
		insertedDate string
		date         []string
		max          []float64
		min          []float64
		different    []float64
	}{
		{
			name:         "success",
			inputJSON:    `{"max": 99, "min": 22}`,
			statusCode:   http.StatusCreated,
			message:      "success update weight",
			insertedDate: "15-06-1997",
			date:         []string{"15-06-1997", "Rata-rata"},
			max:          []float64{99, 99},
			min:          []float64{22, 22},
			different:    []float64{77, 77},
		},
		{
			name:         "error format date",
			inputJSON:    `{"max": "abcd", "min": 5}`,
			statusCode:   http.StatusBadRequest,
			message:      "date format must be dd-mm-yyyy",
			insertedDate: "15-13-1997",
		},
		{
			name:         "error decode",
			inputJSON:    `{"max": "abcd", "min": 5}`,
			statusCode:   http.StatusBadRequest,
			message:      "json: cannot unmarshal string into Go struct field Weight.max of type float64",
			insertedDate: "15-06-1997",
		},
		{
			name:         "error validate",
			inputJSON:    `{"date": "15-06-1997", "max": 1, "min": 5}`,
			statusCode:   http.StatusBadRequest,
			message:      "maximum weight cannot lower than minimum weight",
			insertedDate: "15-06-1997",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			route := mux.NewRouter()
			route.Handle("/weight", hdl.Update()).Methods("PUT")

			req, err := http.NewRequest(http.MethodPut, "/weight?date="+tt.insertedDate, bytes.NewBufferString(tt.inputJSON))
			if err != nil {
				t.Error(err)
			}
			rr := httptest.NewRecorder()
			route.ServeHTTP(rr, req)

			responseMap := jsonHelper.Res{}
			err = json.Unmarshal(rr.Body.Bytes(), &responseMap)
			if err != nil {
				t.Errorf("Cannot convert to struct: %v", err)
			}

			assert.Equal(t, responseMap.Meta.StatusCode, tt.statusCode)
			assert.Equal(t, responseMap.Meta.Message, tt.message)

			if tt.statusCode == 201 {
				responseData := responseMap.Data.([]interface{})
				for i := range tt.date {
					responseData := responseData[i]
					value := responseData.(map[string]interface{})
					assert.Equal(t, value["date"], tt.date[i])
					assert.Equal(t, value["max"], tt.max[i])
					assert.Equal(t, value["min"], tt.min[i])
					assert.Equal(t, value["different"], tt.different[i])
				}
			}
		})
	}
}

func TestWeight_Delete(t *testing.T) {
	mgr, err := initPackages()
	if err != nil {
		log.Fatal(err)
	}

	hdl := handler.NewWeightHandler(mgr)

	err = truncateWeights(mgr.GetPostgres())
	if err != nil {
		log.Fatal(err)
	}

	results, err := seedWeights(mgr.GetPostgres())
	if err != nil {
		log.Fatal(err)
	}

	tests := []struct {
		name         string
		inputJSON    string
		statusCode   int
		message      string
		insertedDate string
		date         []string
		max          []float64
		min          []float64
		different    []float64
	}{
		{
			name:         "success",
			inputJSON:    `{"max": 99, "min": 22}`,
			statusCode:   http.StatusAccepted,
			message:      "success delete weight",
			insertedDate: "15-06-1997",
			date:         []string{results[0].Date, "Rata-rata"},
			max:          []float64{results[0].Max, 100},
			min:          []float64{results[0].Min, 30},
			different:    []float64{results[0].Different, 70},
		},
		{
			name:         "error format date",
			inputJSON:    `{"max": "abcd", "min": 5}`,
			statusCode:   http.StatusBadRequest,
			message:      "date format must be dd-mm-yyyy",
			insertedDate: "15-13-1997",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			route := mux.NewRouter()
			route.Handle("/weight", hdl.Delete()).Methods("DELETE")

			req, err := http.NewRequest(http.MethodDelete, "/weight?date="+tt.insertedDate, nil)
			if err != nil {
				t.Error(err)
			}
			rr := httptest.NewRecorder()
			route.ServeHTTP(rr, req)

			responseMap := jsonHelper.Res{}
			err = json.Unmarshal(rr.Body.Bytes(), &responseMap)
			if err != nil {
				t.Errorf("Cannot convert to struct: %v", err)
			}

			assert.Equal(t, responseMap.Meta.StatusCode, tt.statusCode)
			assert.Equal(t, responseMap.Meta.Message, tt.message)

			if tt.statusCode == 202 {
				responseData := responseMap.Data.([]interface{})
				for i := range tt.date {
					responseData := responseData[i]
					value := responseData.(map[string]interface{})
					assert.Equal(t, value["date"], tt.date[i])
					assert.Equal(t, value["max"], tt.max[i])
					assert.Equal(t, value["min"], tt.min[i])
					assert.Equal(t, value["different"], tt.different[i])
				}
			}
		})
	}
}
