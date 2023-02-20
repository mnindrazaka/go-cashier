package main

import (
	"fmt"
	"strings"
)

func main() {
	productNames := []string{"Pensil", "Bulpoin", "Buku"}
	productPrices := []int{2500, 3000, 4000}
	isRunning := true

	for isRunning {
		fmt.Println("Selamat Datang di Toko Saya")
		fmt.Println("Daftar Barang : ")
		for i, productName := range productNames {
			fmt.Println(i+1, ".", productName, "\t: Rp.", productPrices[i])
		}

		var selectedProduct int
		fmt.Print("Masukkan pilihan barang \t: ")
		fmt.Scanf("%d", &selectedProduct)

		var amount int
		fmt.Print("Masukkan jumlah barang \t\t: ")
		fmt.Scanf("%d", &amount)

		price := productPrices[selectedProduct-1]
		total := price * amount

		fmt.Println("Totalnya adalah Rp.", total)

		money := 0
		for money < total {
			fmt.Print("Masukkan uang yang dibayar \t: ")
			fmt.Scanf("%d", &money)

			if money < total {
				difference := total - money
				fmt.Println("Maaf uang anda kurang Rp.", difference)
			}
		}

		change := money - total
		fmt.Println("Jumlah kembaliannya adalah Rp.", change)

		var answer string
		fmt.Print("Mau melakukan transaksi lagi (y/n) ? \t: ")
		fmt.Scanf("%s", &answer)

		isRunning = strings.ToLower(answer) == "y"
	}

	fmt.Println("Terimakasih dan sampai jumpa")
}
