package handler

import (
	"bytes"
	"errors"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/gorilla/mux"
	"github.com/hireza/sirclo-test/berat/app/usecase"
	"github.com/hireza/sirclo-test/berat/domain"
	"github.com/hireza/sirclo-test/berat/domain/mocks"
	"github.com/hireza/sirclo-test/berat/packages"
)

var date = "15-06-1997"
var formattedDate = "1997-06-15"
var payload = []byte(`{
	"date": "15-06-1997",
	"max": 10,
	"min": 5
}`)
var dataBody = &domain.Weight{
	Date:      formattedDate,
	Max:       10,
	Min:       5,
	Different: 5,
}
var payloadErrorDecode = []byte(`{
	"date": "15-06-1997",
	"max": "abcd",
	"min": 5
}`)
var payloadErrorValidate = []byte(`{
	"date": "15-06-1997",
	"max": 1,
	"min": 5
}`)
var result = &domain.Weight{
	Date:      date,
	Max:       10,
	Min:       5,
	Different: 5,
}
var expectedResults = []*domain.Weight{
	result,
}

func TestNewWeightHandler(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	fmgr, err := packages.NewFakeInit(ctrl)
	if err != nil {
		t.Fatalf("Error create packages instance: err %+v", err)
	}

	h := new(Weight)
	h.WeightUsecase = usecase.NewWeightUsecase(fmgr)

	type args struct {
		mgr packages.Packages
	}
	tests := []struct {
		name string
		args args
		want domain.WeightHandler
	}{
		{
			name: "success",
			args: args{
				mgr: fmgr,
			},
			want: h,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewWeightHandler(tt.args.mgr); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewWeightHandler() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWeight_GetAll(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUsecase := mocks.NewMockWeightUsecase(ctrl)

	type fields struct {
		WeightUsecase domain.WeightUsecase
	}
	tests := []struct {
		name               string
		fields             fields
		requestURL         string
		expected           *gomock.Call
		expectedStatusCode int
	}{
		{
			name: "success",
			fields: fields{
				WeightUsecase: mockUsecase,
			},
			requestURL:         "/weights",
			expected:           mockUsecase.EXPECT().GetAll(gomock.Any()).Return(expectedResults, nil),
			expectedStatusCode: http.StatusOK,
		},
		{
			name: "error",
			fields: fields{
				WeightUsecase: mockUsecase,
			},
			requestURL:         "/weights",
			expected:           mockUsecase.EXPECT().GetAll(gomock.Any()).Return(nil, errors.New("error")),
			expectedStatusCode: http.StatusInternalServerError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &Weight{
				WeightUsecase: tt.fields.WeightUsecase,
			}

			r := httptest.NewRequest("GET", tt.requestURL, nil)
			rr := httptest.NewRecorder()

			router := mux.NewRouter()

			router.Handle("/weights", h.GetAll())
			router.ServeHTTP(rr, r)

			if status := rr.Code; status != tt.expectedStatusCode {
				t.Fatalf("wrong status code: got %d want %d", status, tt.expectedStatusCode)
			}
		})
	}
}

func TestWeight_GetByDate(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUsecase := mocks.NewMockWeightUsecase(ctrl)

	type fields struct {
		WeightUsecase domain.WeightUsecase
	}
	tests := []struct {
		name               string
		fields             fields
		requestURL         string
		expected           *gomock.Call
		expectedStatusCode int
	}{
		{
			name: "success",
			fields: fields{
				WeightUsecase: mockUsecase,
			},
			requestURL:         "/weight?date=15-06-1997",
			expected:           mockUsecase.EXPECT().GetByDate(gomock.Any(), formattedDate).Return(result, nil),
			expectedStatusCode: http.StatusOK,
		},
		{
			name: "error format date",
			fields: fields{
				WeightUsecase: mockUsecase,
			},
			requestURL:         "/weight?date=15-13-1997",
			expectedStatusCode: http.StatusBadRequest,
		},
		{
			name: "error getByData - not found",
			fields: fields{
				WeightUsecase: mockUsecase,
			},
			requestURL:         "/weight?date=15-06-1997",
			expected:           mockUsecase.EXPECT().GetByDate(gomock.Any(), formattedDate).Return(nil, errors.New("no rows")),
			expectedStatusCode: http.StatusNotFound,
		},
		{
			name: "error getByData - error",
			fields: fields{
				WeightUsecase: mockUsecase,
			},
			requestURL:         "/weight?date=15-06-1997",
			expected:           mockUsecase.EXPECT().GetByDate(gomock.Any(), formattedDate).Return(nil, errors.New("error")),
			expectedStatusCode: http.StatusInternalServerError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &Weight{
				WeightUsecase: tt.fields.WeightUsecase,
			}

			r := httptest.NewRequest("GET", tt.requestURL, nil)
			rr := httptest.NewRecorder()

			router := mux.NewRouter()

			router.Handle("/weight", h.GetByDate())
			router.ServeHTTP(rr, r)

			if status := rr.Code; status != tt.expectedStatusCode {
				t.Fatalf("wrong status code: got %d want %d", status, tt.expectedStatusCode)
			}
		})
	}
}

func TestWeight_Create(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUsecase := mocks.NewMockWeightUsecase(ctrl)

	type fields struct {
		WeightUsecase domain.WeightUsecase
	}
	tests := []struct {
		name               string
		fields             fields
		requestURL         string
		payload            []byte
		expected           *gomock.Call
		expectedStatusCode int
	}{
		{
			name: "success",
			fields: fields{
				WeightUsecase: mockUsecase,
			},
			requestURL:         "/weight",
			expected:           mockUsecase.EXPECT().Create(gomock.Any(), dataBody).Return(expectedResults, nil),
			expectedStatusCode: http.StatusCreated,
			payload:            payload,
		},
		{
			name: "error decode",
			fields: fields{
				WeightUsecase: mockUsecase,
			},
			requestURL:         "/weight",
			expectedStatusCode: http.StatusBadRequest,
			payload:            payloadErrorDecode,
		},
		{
			name: "error validate",
			fields: fields{
				WeightUsecase: mockUsecase,
			},
			requestURL:         "/weight",
			expectedStatusCode: http.StatusBadRequest,
			payload:            payloadErrorValidate,
		},
		{
			name: "error create",
			fields: fields{
				WeightUsecase: mockUsecase,
			},
			requestURL:         "/weight",
			expected:           mockUsecase.EXPECT().Create(gomock.Any(), dataBody).Return(nil, errors.New("error")),
			expectedStatusCode: http.StatusInternalServerError,
			payload:            payload,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &Weight{
				WeightUsecase: tt.fields.WeightUsecase,
			}

			r := httptest.NewRequest("POST", tt.requestURL, bytes.NewBuffer(tt.payload))
			rr := httptest.NewRecorder()

			router := mux.NewRouter()

			router.Handle("/weight", h.Create())
			router.ServeHTTP(rr, r)

			if status := rr.Code; status != tt.expectedStatusCode {
				t.Fatalf("wrong status code: got %d want %d", status, tt.expectedStatusCode)
			}
		})
	}
}

func TestWeight_Update(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUsecase := mocks.NewMockWeightUsecase(ctrl)

	type fields struct {
		WeightUsecase domain.WeightUsecase
	}
	tests := []struct {
		name               string
		fields             fields
		requestURL         string
		payload            []byte
		expected           *gomock.Call
		expectedStatusCode int
	}{
		{
			name: "success",
			fields: fields{
				WeightUsecase: mockUsecase,
			},
			requestURL:         "/weight?date=15-06-1997",
			expected:           mockUsecase.EXPECT().Update(gomock.Any(), formattedDate, dataBody).Return(expectedResults, nil),
			expectedStatusCode: http.StatusCreated,
			payload:            payload,
		},
		{
			name: "error format date",
			fields: fields{
				WeightUsecase: mockUsecase,
			},
			requestURL:         "/weight?date=15-13-1997",
			expectedStatusCode: http.StatusBadRequest,
			payload:            payload,
		},
		{
			name: "error decode",
			fields: fields{
				WeightUsecase: mockUsecase,
			},
			requestURL:         "/weight?date=15-06-1997",
			expectedStatusCode: http.StatusBadRequest,
			payload:            payloadErrorDecode,
		},
		{
			name: "error validate",
			fields: fields{
				WeightUsecase: mockUsecase,
			},
			requestURL:         "/weight?date=15-06-1997",
			expectedStatusCode: http.StatusBadRequest,
			payload:            payloadErrorValidate,
		},
		{
			name: "error update",
			fields: fields{
				WeightUsecase: mockUsecase,
			},
			requestURL:         "/weight?date=15-06-1997",
			expected:           mockUsecase.EXPECT().Update(gomock.Any(), formattedDate, dataBody).Return(nil, errors.New("error")),
			expectedStatusCode: http.StatusInternalServerError,
			payload:            payload,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &Weight{
				WeightUsecase: tt.fields.WeightUsecase,
			}

			r := httptest.NewRequest("PUT", tt.requestURL, bytes.NewBuffer(tt.payload))
			rr := httptest.NewRecorder()

			router := mux.NewRouter()

			router.Handle("/weight", h.Update())
			router.ServeHTTP(rr, r)

			if status := rr.Code; status != tt.expectedStatusCode {
				t.Fatalf("wrong status code: got %d want %d", status, tt.expectedStatusCode)
			}
		})
	}
}

func TestWeight_Delete(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUsecase := mocks.NewMockWeightUsecase(ctrl)

	type fields struct {
		WeightUsecase domain.WeightUsecase
	}
	tests := []struct {
		name               string
		fields             fields
		requestURL         string
		expected           *gomock.Call
		expectedStatusCode int
	}{
		{
			name: "success",
			fields: fields{
				WeightUsecase: mockUsecase,
			},
			requestURL:         "/weight?date=15-06-1997",
			expected:           mockUsecase.EXPECT().Delete(gomock.Any(), formattedDate).Return(expectedResults, nil),
			expectedStatusCode: http.StatusAccepted,
		},
		{
			name: "error format date",
			fields: fields{
				WeightUsecase: mockUsecase,
			},
			requestURL:         "/weight?date=15-13-1997",
			expectedStatusCode: http.StatusBadRequest,
		},
		{
			name: "error delete",
			fields: fields{
				WeightUsecase: mockUsecase,
			},
			requestURL:         "/weight?date=15-06-1997",
			expected:           mockUsecase.EXPECT().Delete(gomock.Any(), formattedDate).Return(nil, errors.New("error")),
			expectedStatusCode: http.StatusInternalServerError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &Weight{
				WeightUsecase: tt.fields.WeightUsecase,
			}

			r := httptest.NewRequest("DELETE", tt.requestURL, nil)
			rr := httptest.NewRecorder()

			router := mux.NewRouter()

			router.Handle("/weight", h.Delete())
			router.ServeHTTP(rr, r)

			if status := rr.Code; status != tt.expectedStatusCode {
				t.Fatalf("wrong status code: got %d want %d", status, tt.expectedStatusCode)
			}
		})
	}
}
