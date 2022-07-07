package main

import (
	"testing"
)

func TestCart_tambahProduk(t *testing.T) {
	availableCart := map[string]int{"barang 1": 1}

	type fields struct {
		items map[string]int
	}
	type args struct {
		kodeProduk string
		kuantitas  int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name: "available in cart",
			fields: fields{
				items: availableCart,
			},
			args: args{
				kodeProduk: "barang 1",
				kuantitas:  2,
			},
		},
		{
			name: "not available in cart",
			fields: fields{
				items: map[string]int{},
			},
			args: args{
				kodeProduk: "barang 2",
				kuantitas:  2,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Cart{
				items: tt.fields.items,
			}
			c.tambahProduk(tt.args.kodeProduk, tt.args.kuantitas)
		})
	}
}

func TestCart_hapusProduk(t *testing.T) {
	availableCart := map[string]int{"barang 1": 1}

	type fields struct {
		items map[string]int
	}
	type args struct {
		kodeProduk string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name: "delete cart",
			fields: fields{
				items: availableCart,
			},
			args: args{
				kodeProduk: "barang 1",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Cart{
				items: tt.fields.items,
			}
			c.hapusProduk(tt.args.kodeProduk)
		})
	}
}

func TestCart_tampilkanCart(t *testing.T) {
	availableCart := map[string]int{"barang 1": 1}

	type fields struct {
		items map[string]int
	}
	tests := []struct {
		name   string
		fields fields
	}{
		{
			name: "show cart",
			fields: fields{
				items: availableCart,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Cart{
				items: tt.fields.items,
			}
			c.tampilkanCart()
		})
	}
}

func Test_main(t *testing.T) {
	tests := []struct {
		name string
	}{
		{
			name: "do main",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			main()
		})
	}
}
