package main

import "fmt"

type Transaction struct {
	Product Product
	Amount  int
	Total   int
	Money   int
	Change  int
}

func (transaction *Transaction) inputProduct(products []Product) {
	var selectedProduct int
	fmt.Print("Masukkan pilihan barang \t: ")
	fmt.Scanf("%d", &selectedProduct)
	transaction.Product = products[selectedProduct-1]
}

func (transaction *Transaction) inputAmount() {
	fmt.Print("Masukkan jumlah barang \t\t: ")
	fmt.Scanf("%d", &transaction.Amount)
}

func (transaction *Transaction) printTotal() {
	transaction.Total = transaction.Product.Price * transaction.Amount
	fmt.Println("Totalnya adalah Rp.", transaction.Total)
}

func (transaction *Transaction) inputMoney() {
	transaction.Money = 0
	for transaction.Money < transaction.Total {
		fmt.Print("Masukkan uang yang dibayar \t: ")
		fmt.Scanf("%d", &transaction.Money)

		if transaction.Money < transaction.Total {
			difference := transaction.Total - transaction.Money
			fmt.Println("Maaf uang anda kurang Rp.", difference)
		}
	}
}

func (transaction *Transaction) printChange() {
	transaction.Change = transaction.Money - transaction.Total
	fmt.Println("Jumlah kembaliannya adalah Rp.", transaction.Change)
}
