package domain

import (
	"testing"
)

func TestWeight_Validate(t *testing.T) {
	type fields struct {
		ID        string
		Date      string
		Max       float64
		Min       float64
		Different float64
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name: "success",
			fields: fields{
				Date: "15-06-1997",
				Max:  2,
				Min:  1,
			},
		},
		{
			name: "error date",
			fields: fields{
				Date: "15-13-1997",
				Max:  2,
				Min:  1,
			},
			wantErr: true,
		},
		{
			name: "error max min",
			fields: fields{
				Date: "15-06-1997",
				Max:  0,
				Min:  1,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := &Weight{
				ID:        tt.fields.ID,
				Date:      tt.fields.Date,
				Max:       tt.fields.Max,
				Min:       tt.fields.Min,
				Different: tt.fields.Different,
			}
			if err := w.Validate(); (err != nil) != tt.wantErr {
				t.Errorf("Weight.Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestWeight_GetDifferent(t *testing.T) {
	type fields struct {
		ID        string
		Date      string
		Max       float64
		Min       float64
		Different float64
	}
	tests := []struct {
		name   string
		fields fields
	}{
		{
			name: "success",
			fields: fields{
				Max: 2,
				Min: 1,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := &Weight{
				ID:        tt.fields.ID,
				Date:      tt.fields.Date,
				Max:       tt.fields.Max,
				Min:       tt.fields.Min,
				Different: tt.fields.Different,
			}
			w.GetDifferent()
		})
	}
}

func TestWeight_GetIndoFormattedDate(t *testing.T) {
	type fields struct {
		ID        string
		Date      string
		Max       float64
		Min       float64
		Different float64
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name: "success",
			fields: fields{
				Date: "1997-06-15T15:04:05Z",
			},
		},
		{
			name: "error",
			fields: fields{
				Date: "15-06-1997",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := &Weight{
				ID:        tt.fields.ID,
				Date:      tt.fields.Date,
				Max:       tt.fields.Max,
				Min:       tt.fields.Min,
				Different: tt.fields.Different,
			}
			if err := w.GetIndoFormattedDate(); (err != nil) != tt.wantErr {
				t.Errorf("Weight.GetIndoFormattedDate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
