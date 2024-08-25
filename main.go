package main

import "fmt"

type item struct {
	id    int
	name  string
	price int
}

type orderItem struct {
	item   item
	jumlah int
	total  int
}

var menu = []item{
	{id: 1, name: "Ramen", price: 15000},
	{id: 2, name: "Nasi Goreng", price: 20000},
	{id: 3, name: "Ayam Bakar", price: 25000},
	{id: 4, name: "Mie Goreng", price: 20000},
	{id: 5, name: "Ayam Goreng", price: 25000},
	{id: 6, name: "Es Teh", price: 5000},
	{id: 7, name: "Es Jeruk", price: 7000},
	{id: 8, name: "Es Cincau", price: 8000},
	{id: 9, name: "Air", price: 1000},
}

var orders []orderItem

func main() {
	var pilih int
	// var orders [][]item
	var totalBill int
	var jml int

	for {
		daftarMenu()
		fmt.Print("Pilih Menu (0 untuk keluar) : ")
		fmt.Scan(&pilih)

		//untuk keluar dari program
		if pilih == 0 {
			break
		}

		if pilih > 0 && pilih <= len(menu) {

			item := menu[pilih-1]
			fmt.Print("Masukkan Jumlah : ")
			fmt.Scan(&jml)

			newOrder := orderItem{
				item:   item,
				jumlah: jml,
				total:  item.price * jml,
			}
			orders = append(orders, newOrder)
			totalBill += newOrder.total
			fmt.Printf("%s ditambahkan ke pesanan.\n", item.name)
		} else {
			fmt.Println("Menu tidak valid. Silakan pilih lagi.")
		}

		// Tampilkan pesanan dan total tagihan
		fmt.Println("\nPesanan Anda:")
		for _, order := range orders {
			fmt.Printf("%dx - %s - Rp%d\n", order.jumlah, order.item.name, order.total)
		}
		fmt.Printf("----------------------------------\n")
		fmt.Printf("Total tagihan: Rp%d\n\n", totalBill)

	}

	paymentBill(orders, totalBill)

}

func daftarMenu() {
	fmt.Println("\n Daftar Menu :")
	for k, item := range menu {
		fmt.Printf("%d. %s - Rp%d\n", k+1, item.name, item.price)
	}
}

func paymentBill(orders []orderItem, totalBill int) {
	fmt.Printf("----------------------------------\n")
	fmt.Print("\nDaftar Pembelian Anda : \n\n")
	for _, order := range orders {
		fmt.Printf("%dx - %s - Rp%d\n", order.jumlah, order.item.name, order.total)
	}
	fmt.Printf("----------------------------------\n")
	fmt.Printf("Total tagihan: Rp%d\n", totalBill)
	fmt.Printf("----------------------------------\n\n")


	payment := payment(totalBill)
	change := payment - totalBill

	fmt.Printf("\nPembayaran: Rp.%d\n", payment)
	

	if change == 0 {
		fmt.Println("Uang Anda sudah sesuai. Terima kasih.")
	} else {
		fmt.Printf("Kembalian Anda: Rp.%d\n\n", change)
	}

	

	fmt.Println("Terima kasih sudah berbelanja di warung kami.")
}

func payment(totalBill int) int{
	for {
		var payment int
		fmt.Print("Masukkan nominal pembayaran : Rp.")
		fmt.Scanln(&payment)
		
		if payment < totalBill {
			fmt.Println("Uang Anda kurang. Silakan ulangi lagi.")
		} else {
			return payment
		}
		
	}
}
