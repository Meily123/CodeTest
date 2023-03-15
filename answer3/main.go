package main

import "fmt"

/*
Question :
Yayat akan pulang kampung ke daerah serang,dia akan menempuh jarak 63 km,dia
menggunakan sepeda motor untuk sampai ke serang,motor yang ia gunakan
menghabiskan bensin 17km / 1 liter,berapakah biaya yang dikeluarkan oleh yayat jika
harga bensin 8.350 / liter?
*/

func main() {
	var jarakTempuh float64 = 63
	var jarakPerliter float64 = 17
	var hargaBensinPerliter float64 = 8350

	var banyakBensinTerpakai = jarakTempuh / jarakPerliter
	var biayaBensin = hargaBensinPerliter * banyakBensinTerpakai

	fmt.Printf("Biaya yang dikeluarkan oleh Yayat adalah RP. %.2f \n", biayaBensin)
}
