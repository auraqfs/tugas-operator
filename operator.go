package main

import "fmt"

func main() {
	var numA, numMenu float32
	var dollar float32 = 14232.50
	var euro float32 = 16890.10

	fmt.Println("=======================")
	fmt.Println("Pilih Menu di bawah : ")
	fmt.Println("1. Penukaran Rupiah ke Dollar")
	fmt.Println("2. Penukaran Rupiah ke Euro")
	fmt.Println("3. Penukaran GBP ke knut")
	fmt.Println("=======================")
	fmt.Print("Pilih Menu angka di atas [1/2/3] : ")
	fmt.Scan(&numMenu)
	fmt.Println("=======================")

	if numMenu == 1 {
		fmt.Print("Masukan Nilai : ")
		fmt.Scan(&numA)
		fmt.Print("Hasil penukaran Rp", numA, " = $")
		fmt.Printf("%.2f", numA/dollar)
	} else if numMenu == 2 {
		fmt.Print("Masukan Nilai : ")
		fmt.Scan(&numA)
		fmt.Print("Hasil penukaran yaitu Rp", numA, " = €")
		fmt.Printf("%.2f", numA/euro)
	} else {
		fmt.Print("Masukan jumlah GB Pounds : ")
		fmt.Scan(&numA)
		fmt.Println("Jumlah Knut yang didapatkan = ", numA*100)
		fmt.Println("Hasil penukaran mendapatkan = ", numA/5, "Galleon(s)")
	}
}
