package route

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/gorilla/mux"
	"github.com/hireza/sirclo-test/berat/packages"
)

func TestNewWeightRoute(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	fmgr, err := packages.NewFakeInit(ctrl)
	if err != nil {
		t.Fatalf("Error create packages instance: err %+v", err)
	}

	r := mux.NewRouter()

	type args struct {
		route *mux.Router
		mgr   packages.Packages
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "success",
			args: args{
				mgr:   fmgr,
				route: r,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			NewWeightRoute(tt.args.route, tt.args.mgr)
		})
	}
}
