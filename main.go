package main

import (
	"fmt"
	"strings"
)

func main() {
	shop := Shop{Products: []Product{
		{Name: "Pensil", Price: 2500},
		{Name: "Bulpoin", Price: 3000},
		{Name: "Buku", Price: 4000},
	}}

	isRunning := true
	for isRunning {
		shop.PrintProducts()

		transaction := Transaction{}
		transaction.inputProduct(shop.Products)
		transaction.inputAmount()
		transaction.printTotal()
		transaction.inputMoney()
		transaction.printChange()

		var answer string
		fmt.Print("Mau melakukan transaksi lagi (y/n) ? \t: ")
		fmt.Scanf("%s", &answer)
		isRunning = strings.ToLower(answer) == "y"
	}

	fmt.Println("Terimakasih dan sampai jumpa")
}
