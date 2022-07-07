package repository

import (
	"context"
	"database/sql"
	"errors"
	"reflect"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/golang/mock/gomock"
	"github.com/hireza/sirclo-test/berat/domain"
	"github.com/hireza/sirclo-test/berat/packages"
)

var ctx = context.Background()
var date = "15-06-1997"
var sqlResult1 = &domain.Weight{
	ID:        "1",
	Date:      "1997-06-15T15:04:05Z",
	Max:       10,
	Min:       5,
	Different: 5,
}
var sqlResult2 = &domain.Weight{
	ID:        "2",
	Date:      "1997-06-16T15:04:05Z",
	Max:       10,
	Min:       5,
	Different: 5,
}
var sqlResults = []*domain.Weight{
	sqlResult1,
	sqlResult2,
}
var formattedResult1 = &domain.Weight{
	ID:        "1",
	Date:      "15-06-1997",
	Max:       10,
	Min:       5,
	Different: 5,
}
var formattedResult2 = &domain.Weight{
	ID:        "2",
	Date:      "16-06-1997",
	Max:       10,
	Min:       5,
	Different: 5,
}
var formattedResults = []*domain.Weight{
	formattedResult2,
	formattedResult1,
	{
		Date:      "Rata-rata",
		Max:       10,
		Min:       5,
		Different: 5,
	},
}

func TestNewWeightRepository(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	fmgr, _ := packages.NewFakeInit(ctrl)

	type args struct {
		mgr packages.Packages
	}
	tests := []struct {
		name string
		args args
		want domain.WeightRepository
	}{
		{
			name: "success",
			args: args{mgr: fmgr},
			want: NewWeightRepository(fmgr),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewWeightRepository(tt.args.mgr); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewWeightRepository() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWeight_CheckByDate(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatal("Error connection to DB")
	}
	defer db.Close()

	query := "SELECT"

	result := true

	rows := sqlmock.NewRows([]string{"exists"}).
		AddRow(result)

	type fields struct {
		DB *sql.DB
	}
	type args struct {
		ctx  context.Context
		date string
	}
	tests := []struct {
		name          string
		fields        fields
		args          args
		expectedQuery *sqlmock.ExpectedQuery
		want          bool
		wantErr       bool
	}{
		{
			name: "success",
			fields: fields{
				DB: db,
			},
			args: args{
				ctx:  ctx,
				date: date,
			},
			expectedQuery: mock.ExpectQuery(query).WillReturnRows(rows),
			want:          true,
		},
		{
			name: "error",
			fields: fields{
				DB: db,
			},
			args: args{
				ctx:  ctx,
				date: date,
			},
			expectedQuery: mock.ExpectQuery(query).WillReturnError(errors.New("error")),
			wantErr:       true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Weight{
				DB: tt.fields.DB,
			}
			got, err := r.CheckByDate(tt.args.ctx, tt.args.date)
			if (err != nil) != tt.wantErr {
				t.Errorf("Weight.CheckByDate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Weight.CheckByDate() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWeight_GetAll(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatal("Error connection to DB")
	}
	defer db.Close()

	query := "SELECT"

	rows := sqlmock.NewRows([]string{"id", "date", "max", "min"}).
		AddRow(sqlResults[0].ID, sqlResults[0].Date, sqlResults[0].Max, sqlResults[0].Min).
		AddRow(sqlResults[1].ID, sqlResults[1].Date, sqlResults[1].Max, sqlResults[1].Min)

	rowsFalse := sqlmock.NewRows([]string{"id", "date", "max", "min"}).
		AddRow(sqlResults[0].ID, sqlResults[0].Date, "abcd", "efgh")

	rowsErrFormat := sqlmock.NewRows([]string{"id", "date", "max", "min"}).
		AddRow(sqlResults[0].ID, "15-06-1997", 2, 1)

	type fields struct {
		DB *sql.DB
	}
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name          string
		fields        fields
		args          args
		expectedQuery *sqlmock.ExpectedQuery
		want          []*domain.Weight
		wantErr       bool
	}{
		{
			name: "success",
			fields: fields{
				DB: db,
			},
			args: args{
				ctx: ctx,
			},
			expectedQuery: mock.ExpectQuery(query).WillReturnRows(rows),
			want:          formattedResults,
		},
		{
			name: "error select",
			fields: fields{
				DB: db,
			},
			args: args{
				ctx: ctx,
			},
			expectedQuery: mock.ExpectQuery(query).WillReturnError(errors.New("error")),
			wantErr:       true,
		},
		{
			name: "error scan",
			fields: fields{
				DB: db,
			},
			args: args{
				ctx: ctx,
			},
			expectedQuery: mock.ExpectQuery(query).WillReturnRows(rowsFalse),
			wantErr:       true,
		},
		{
			name: "error format",
			fields: fields{
				DB: db,
			},
			args: args{
				ctx: ctx,
			},
			expectedQuery: mock.ExpectQuery(query).WillReturnRows(rowsErrFormat),
			wantErr:       true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Weight{
				DB: tt.fields.DB,
			}
			got, err := r.GetAll(tt.args.ctx)
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
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatal("Error connection to DB")
	}
	defer db.Close()

	query := "SELECT"

	rows := sqlmock.NewRows([]string{"id", "date", "max", "min"}).
		AddRow(sqlResult1.ID, sqlResult1.Date, sqlResult1.Max, sqlResult1.Min)

	rowsFalse := sqlmock.NewRows([]string{"id", "date", "max", "min"}).
		AddRow(sqlResult1.ID, sqlResult1.Date, "abc", sqlResult1.Min)

	rowsErrFormat := sqlmock.NewRows([]string{"id", "date", "max", "min"}).
		AddRow(sqlResult1.ID, "15-06-1997", 2, 1)

	type fields struct {
		DB *sql.DB
	}
	type args struct {
		ctx  context.Context
		date string
	}
	tests := []struct {
		name          string
		fields        fields
		args          args
		expectedQuery *sqlmock.ExpectedQuery
		want          *domain.Weight
		wantErr       bool
	}{
		{
			name: "success",
			fields: fields{
				DB: db,
			},
			args: args{
				ctx:  ctx,
				date: date,
			},
			expectedQuery: mock.ExpectQuery(query).WillReturnRows(rows),
			want:          formattedResult1,
		},
		{
			name: "error scan",
			fields: fields{
				DB: db,
			},
			args: args{
				ctx:  ctx,
				date: date,
			},
			expectedQuery: mock.ExpectQuery(query).WillReturnRows(rowsFalse),
			wantErr:       true,
		},
		{
			name: "error format",
			fields: fields{
				DB: db,
			},
			args: args{
				ctx:  ctx,
				date: date,
			},
			expectedQuery: mock.ExpectQuery(query).WillReturnRows(rowsErrFormat),
			wantErr:       true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Weight{
				DB: tt.fields.DB,
			}
			got, err := r.GetByDate(tt.args.ctx, tt.args.date)
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
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatal("Error connection to DB")
	}
	defer db.Close()

	query := "INSERT"

	type fields struct {
		DB *sql.DB
	}
	type args struct {
		ctx  context.Context
		data *domain.Weight
	}
	tests := []struct {
		name       string
		fields     fields
		args       args
		expectExec *sqlmock.ExpectedExec
		wantErr    bool
	}{
		{
			name: "success",
			fields: fields{
				DB: db,
			},
			args: args{
				ctx:  ctx,
				data: formattedResult1,
			},
			expectExec: mock.ExpectExec(query).WithArgs(
				formattedResult1.Date,
				formattedResult1.Max,
				formattedResult1.Min,
			).WillReturnResult(sqlmock.NewResult(1, 1)),
		},
		{
			name: "error",
			fields: fields{
				DB: db,
			},
			args: args{
				ctx:  ctx,
				data: formattedResult1,
			},
			expectExec: mock.ExpectExec(query).WithArgs(
				formattedResult1.Date,
				formattedResult1.Max,
				formattedResult1.Min,
			).WillReturnError(errors.New("error")),
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Weight{
				DB: tt.fields.DB,
			}
			if err := r.Create(tt.args.ctx, tt.args.data); (err != nil) != tt.wantErr {
				t.Errorf("Weight.Create() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestWeight_Update(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatal("Error connection to DB")
	}
	defer db.Close()

	query := "UPDATE"

	type fields struct {
		DB *sql.DB
	}
	type args struct {
		ctx  context.Context
		date string
		data *domain.Weight
	}
	tests := []struct {
		name       string
		fields     fields
		args       args
		expectExec *sqlmock.ExpectedExec
		wantErr    bool
	}{
		{
			name: "success",
			fields: fields{
				DB: db,
			},
			args: args{
				ctx:  ctx,
				date: date,
				data: sqlResult1,
			},
			expectExec: mock.ExpectExec(query).WithArgs(
				sqlResult1.Date,
				sqlResult1.Max,
				sqlResult1.Min,
			).WillReturnResult(sqlmock.NewResult(1, 1)),
		},
		{
			name: "error",
			fields: fields{
				DB: db,
			},
			args: args{
				ctx:  ctx,
				date: date,
				data: sqlResult1,
			},
			expectExec: mock.ExpectExec(query).WithArgs(
				sqlResult1.Date,
				sqlResult1.Max,
				sqlResult1.Min,
			).WillReturnError(errors.New("error")),
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Weight{
				DB: tt.fields.DB,
			}
			if err := r.Update(tt.args.ctx, tt.args.date, tt.args.data); (err != nil) != tt.wantErr {
				t.Errorf("Weight.Update() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestWeight_Delete(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatal("Error connection to DB")
	}
	defer db.Close()

	query := "UPDATE"

	type fields struct {
		DB *sql.DB
	}
	type args struct {
		ctx  context.Context
		date string
	}
	tests := []struct {
		name       string
		fields     fields
		args       args
		expectExec *sqlmock.ExpectedExec
		wantErr    bool
	}{
		{
			name: "success",
			fields: fields{
				DB: db,
			},
			args: args{
				ctx:  ctx,
				date: date,
			},
			expectExec: mock.ExpectExec(query).WithArgs(
				date,
			).WillReturnResult(sqlmock.NewResult(1, 1)),
		},
		{
			name: "error",
			fields: fields{
				DB: db,
			},
			args: args{
				ctx:  ctx,
				date: date,
			},
			expectExec: mock.ExpectExec(query).WithArgs(
				date,
			).WillReturnError(errors.New("error")),
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Weight{
				DB: tt.fields.DB,
			}
			if err := r.Delete(tt.args.ctx, tt.args.date); (err != nil) != tt.wantErr {
				t.Errorf("Weight.Delete() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
