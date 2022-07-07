package json

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestWriteJSON(t *testing.T) {
	w := httptest.NewRecorder()
	code := 200
	res := Res{
		Meta: Meta{
			StatusCode: code,
			Message:    "success",
		},
		Data: true,
	}

	type args struct {
		w    http.ResponseWriter
		code int
		v    interface{}
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "success",
			args: args{
				w:    w,
				code: code,
				v:    res,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := WriteJSON(tt.args.w, tt.args.code, tt.args.v); (err != nil) != tt.wantErr {
				t.Errorf("WriteJSON() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestResponse(t *testing.T) {
	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "/", nil)

	type args struct {
		w          http.ResponseWriter
		r          *http.Request
		statusCode int
		message    interface{}
		data       interface{}
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "success",
			args: args{
				w:          w,
				r:          r,
				statusCode: 200,
				message:    "success",
				data:       true,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Response(tt.args.w, tt.args.r, tt.args.statusCode, tt.args.message, tt.args.data)
		})
	}
}
