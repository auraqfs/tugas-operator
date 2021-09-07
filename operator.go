package main

import (
	"fmt"

	"github.com/yudapc/go-rupiah"
)

func main() {
	var numA, numMenu float32
	var numB int

	fmt.Println("=======================")
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
		numAFloat := float64(numA)
		formatRupiah := rupiah.FormatRupiah(numAFloat)
		formatDollar := numAFloat / 14232.50

		fmt.Print("Hasil penukaran ", formatRupiah, ",00", " = $")
		fmt.Printf("%.2f", formatDollar)

	} else if numMenu == 2 {
		fmt.Print("Masukan Nilai : ")
		fmt.Scan(&numA)
		numAFloat := float64(numA)
		formatRupiah := rupiah.FormatRupiah(numAFloat)
		formatEuro := numAFloat / 16890.10

		fmt.Print("Hasil penukaran yaitu ", formatRupiah, ",00", " = €")
		fmt.Printf("%.2f", formatEuro)
	} else {
		fmt.Print("Masukan jumlah GB Pounds : ")
		fmt.Scan(&numB)
		knut := numB * 100
		fmt.Println("Jumlah Knut yang didapatkan = ", knut)
		galleon := numB / 5
		fmt.Println("Hasil penukaran mendapatkan = ", galleon, "Galleon(s)")
		sickle := galleon * 17
		fmt.Println("Sisa ditukar menjadi = ", sickle%galleon, "Sickle(s)")
		fmt.Println("Keping knut yang tersisa = ", knut-(sickle*29), "Knut(s)")

	}
}
