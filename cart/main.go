package main

import (
	"fmt"
)

type Cart struct {
	items map[string]int
}

func (c *Cart) tambahProduk(kodeProduk string, kuantitas int) {
	if len(c.items) == 0 {
		c.items = make(map[string]int)
	}

	// if product available in cart (only add quantity)
	if v, ok := c.items[kodeProduk]; ok {
		c.items[kodeProduk] = v + kuantitas
		return
	}

	// if product not available in cart (add new item & quantity)
	c.items[kodeProduk] = kuantitas
}

func (c *Cart) hapusProduk(kodeProduk string) {
	// delete key hashmap
	delete(c.items, kodeProduk)
}

func (c *Cart) tampilkanCart() {
	for k, v := range c.items {
		fmt.Printf("%s (%v)\n", k, v)
	}
}

func main() {
	keranjang := Cart{}

	keranjang.tambahProduk("Pisang Hijau", 2)

	keranjang.tambahProduk("Semangka Kuning", 3)

	keranjang.tambahProduk("Apel Merah", 1)
	keranjang.tambahProduk("Apel Merah", 4)
	keranjang.tambahProduk("Apel Merah", 2)

	keranjang.hapusProduk("Semangka Kuning")

	keranjang.hapusProduk("Semangka Merah")

	keranjang.tampilkanCart()
}
