package usecase

import (
	"context"
	"errors"
	"reflect"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/hireza/sirclo-test/berat/domain"
	"github.com/hireza/sirclo-test/berat/packages"

	"github.com/hireza/sirclo-test/berat/app/repository"
	"github.com/hireza/sirclo-test/berat/domain/mocks"
)

var ctx = context.Background()
var date = "15-06-1997"
var dataBody = &domain.Weight{
	ID:        "ABCDEFGHIJ",
	Date:      "15-06-1997",
	Max:       10,
	Min:       5,
	Different: 5,
}
var expectedResults = []*domain.Weight{
	dataBody,
}

func TestNewWeightUsecase(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	fmgr, err := packages.NewFakeInit(ctrl)
	if err != nil {
		t.Fatalf("Error create packages instance: err %+v", err)
	}

	h := new(Weight)
	h.WeightRepository = repository.NewWeightRepository(fmgr)

	type args struct {
		mgr packages.Packages
	}
	tests := []struct {
		name string
		args args
		want domain.WeightUsecase
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
			if got := NewWeightUsecase(tt.args.mgr); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewWeightUsecase() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWeight_GetAll(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepository := mocks.NewMockWeightRepository(ctrl)

	type fields struct {
		WeightRepository domain.WeightRepository
	}
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		call    *gomock.Call
		want    []*domain.Weight
		wantErr bool
	}{
		{
			name: "success",
			fields: fields{
				WeightRepository: mockRepository,
			},
			args: args{
				ctx: ctx,
			},
			call: mockRepository.EXPECT().GetAll(ctx).Return(expectedResults, nil),
			want: expectedResults,
		},
		{
			name: "error",
			fields: fields{
				WeightRepository: mockRepository,
			},
			args: args{
				ctx: ctx,
			},
			call:    mockRepository.EXPECT().GetAll(ctx).Return(nil, errors.New("error")),
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &Weight{
				WeightRepository: tt.fields.WeightRepository,
			}
			got, err := u.GetAll(tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("Weight.GetAll() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Weight.GetAll() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWeight_GetByDate(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepository := mocks.NewMockWeightRepository(ctrl)

	type fields struct {
		WeightRepository domain.WeightRepository
	}
	type args struct {
		ctx  context.Context
		date string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		call    *gomock.Call
		want    *domain.Weight
		wantErr bool
	}{
		{
			name: "success",
			fields: fields{
				WeightRepository: mockRepository,
			},
			args: args{
				ctx:  ctx,
				date: date,
			},
			call: mockRepository.EXPECT().GetByDate(ctx, date).Return(dataBody, nil),
			want: dataBody,
		},
		{
			name: "error",
			fields: fields{
				WeightRepository: mockRepository,
			},
			args: args{
				ctx:  ctx,
				date: date,
			},
			call:    mockRepository.EXPECT().GetByDate(ctx, date).Return(nil, errors.New("error")),
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &Weight{
				WeightRepository: tt.fields.WeightRepository,
			}
			got, err := u.GetByDate(tt.args.ctx, tt.args.date)
			if (err != nil) != tt.wantErr {
				t.Errorf("Weight.GetByDate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Weight.GetByDate() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWeight_Create(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepository := mocks.NewMockWeightRepository(ctrl)

	type fields struct {
		WeightRepository domain.WeightRepository
	}
	type args struct {
		ctx  context.Context
		data *domain.Weight
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		call1   *gomock.Call
		call2   *gomock.Call
		call3   *gomock.Call
		want    []*domain.Weight
		wantErr bool
	}{
		{
			name: "success create",
			fields: fields{
				WeightRepository: mockRepository,
			},
			args: args{
				ctx:  ctx,
				data: dataBody,
			},
			call1: mockRepository.EXPECT().CheckByDate(ctx, dataBody.Date).Return(false, nil),
			call2: mockRepository.EXPECT().Create(ctx, dataBody).Return(nil),
			call3: mockRepository.EXPECT().GetAll(ctx).Return(expectedResults, nil),
			want:  expectedResults,
		},
		{
			name: "success update",
			fields: fields{
				WeightRepository: mockRepository,
			},
			args: args{
				ctx:  ctx,
				data: dataBody,
			},
			call1: mockRepository.EXPECT().CheckByDate(ctx, dataBody.Date).Return(true, nil),
			call2: mockRepository.EXPECT().Update(ctx, dataBody.Date, dataBody).Return(nil),
			call3: mockRepository.EXPECT().GetAll(ctx).Return(expectedResults, nil),
			want:  expectedResults,
		},
		{
			name: "error checkByDate",
			fields: fields{
				WeightRepository: mockRepository,
			},
			args: args{
				ctx:  ctx,
				data: dataBody,
			},
			call1:   mockRepository.EXPECT().CheckByDate(ctx, dataBody.Date).Return(false, errors.New("error")),
			wantErr: true,
		},
		{
			name: "error create",
			fields: fields{
				WeightRepository: mockRepository,
			},
			args: args{
				ctx:  ctx,
				data: dataBody,
			},
			call1:   mockRepository.EXPECT().CheckByDate(ctx, dataBody.Date).Return(false, nil),
			call2:   mockRepository.EXPECT().Create(ctx, dataBody).Return(errors.New("error")),
			wantErr: true,
		},
		{
			name: "error update",
			fields: fields{
				WeightRepository: mockRepository,
			},
			args: args{
				ctx:  ctx,
				data: dataBody,
			},
			call1:   mockRepository.EXPECT().CheckByDate(ctx, dataBody.Date).Return(true, nil),
			call2:   mockRepository.EXPECT().Update(ctx, dataBody.Date, dataBody).Return(errors.New("error")),
			wantErr: true,
		},
		{
			name: "error getAll",
			fields: fields{
				WeightRepository: mockRepository,
			},
			args: args{
				ctx:  ctx,
				data: dataBody,
			},
			call1:   mockRepository.EXPECT().CheckByDate(ctx, dataBody.Date).Return(false, nil),
			call2:   mockRepository.EXPECT().Create(ctx, dataBody).Return(nil),
			call3:   mockRepository.EXPECT().GetAll(ctx).Return(nil, errors.New("error")),
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &Weight{
				WeightRepository: tt.fields.WeightRepository,
			}
			got, err := u.Create(tt.args.ctx, tt.args.data)
			if (err != nil) != tt.wantErr {
				t.Errorf("Weight.Create() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Weight.Create() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWeight_Update(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepository := mocks.NewMockWeightRepository(ctrl)

	type fields struct {
		WeightRepository domain.WeightRepository
	}
	type args struct {
		ctx  context.Context
		date string
		data *domain.Weight
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		call1   *gomock.Call
		call2   *gomock.Call
		want    []*domain.Weight
		wantErr bool
	}{
		{
			name: "success",
			fields: fields{
				WeightRepository: mockRepository,
			},
			args: args{
				ctx:  ctx,
				date: date,
				data: dataBody,
			},
			call1: mockRepository.EXPECT().Update(ctx, date, dataBody).Return(nil),
			call2: mockRepository.EXPECT().GetAll(ctx).Return(expectedResults, nil),
			want:  expectedResults,
		},
		{
			name: "error update",
			fields: fields{
				WeightRepository: mockRepository,
			},
			args: args{
				ctx:  ctx,
				date: date,
				data: dataBody,
			},
			call1:   mockRepository.EXPECT().Update(ctx, date, dataBody).Return(errors.New("error")),
			wantErr: true,
		},
		{
			name: "error getAll",
			fields: fields{
				WeightRepository: mockRepository,
			},
			args: args{
				ctx:  ctx,
				date: date,
				data: dataBody,
			},
			call1:   mockRepository.EXPECT().Update(ctx, date, dataBody).Return(nil),
			call2:   mockRepository.EXPECT().GetAll(ctx).Return(nil, errors.New("error")),
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &Weight{
				WeightRepository: tt.fields.WeightRepository,
			}
			got, err := u.Update(tt.args.ctx, tt.args.date, tt.args.data)
			if (err != nil) != tt.wantErr {
				t.Errorf("Weight.Update() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Weight.Update() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWeight_Delete(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepository := mocks.NewMockWeightRepository(ctrl)

	type fields struct {
		WeightRepository domain.WeightRepository
	}
	type args struct {
		ctx  context.Context
		date string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		call1   *gomock.Call
		call2   *gomock.Call
		want    []*domain.Weight
		wantErr bool
	}{
		{
			name: "success",
			fields: fields{
				WeightRepository: mockRepository,
			},
			args: args{
				ctx:  ctx,
				date: date,
			},
			call1: mockRepository.EXPECT().Delete(ctx, date).Return(nil),
			call2: mockRepository.EXPECT().GetAll(ctx).Return(expectedResults, nil),
			want:  expectedResults,
		},
		{
			name: "error delete",
			fields: fields{
				WeightRepository: mockRepository,
			},
			args: args{
				ctx:  ctx,
				date: date,
			},
			call1:   mockRepository.EXPECT().Delete(ctx, date).Return(errors.New("error")),
			wantErr: true,
		},
		{
			name: "error delete",
			fields: fields{
				WeightRepository: mockRepository,
			},
			args: args{
				ctx:  ctx,
				date: date,
			},
			call1:   mockRepository.EXPECT().Delete(ctx, date).Return(nil),
			call2:   mockRepository.EXPECT().GetAll(ctx).Return(nil, errors.New("error")),
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &Weight{
				WeightRepository: tt.fields.WeightRepository,
			}
			got, err := u.Delete(tt.args.ctx, tt.args.date)
			if (err != nil) != tt.wantErr {
				t.Errorf("Weight.Delete() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Weight.Delete() = %v, want %v", got, tt.want)
			}
		})
	}
}
