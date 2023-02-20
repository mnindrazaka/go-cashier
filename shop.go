package main

import "fmt"

type Shop struct {
	Products []Product
}

func (shop *Shop) PrintProducts() {
	fmt.Println("Selamat Datang di Toko Saya")
	fmt.Println("Daftar Barang : ")
	for i, product := range shop.Products {
		fmt.Println(i+1, ".", product.Name, "\t: Rp.", product.Price)
	}
}
