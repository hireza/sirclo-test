package time

import (
	"testing"
)

func TestGetIndoFormattedDate(t *testing.T) {
	type args struct {
		date string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "success",
			args: args{
				date: "1997-06-15T15:04:05Z",
			},
			want: "15-06-1997",
		},
		{
			name: "error",
			args: args{
				date: "1997-13-15T15:04:05Z",
			},
			want:    "",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetIndoFormattedDate(tt.args.date)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetIndoFormattedDate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("GetIndoFormattedDate() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetGlobalFormattedDate(t *testing.T) {
	type args struct {
		date string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "success",
			args: args{
				date: "15-06-1997",
			},
			want: "1997-06-15",
		},
		{
			name: "error",
			args: args{
				date: "1997-13-15",
			},
			want:    "",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetGlobalFormattedDate(tt.args.date)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetGlobalFormattedDate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("GetGlobalFormattedDate() = %v, want %v", got, tt.want)
			}
		})
	}
}
