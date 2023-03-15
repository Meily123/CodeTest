package main

import "fmt"

/*
Question :
Apip meminjam uang di koperasi sebesar 10.000.000,koperasi memberikan bunga
sebesar 10%.Apip memilih cicilan selama 12 bulan,buatkan koding perhitungan
dengan golang
*/

func cicilanPinjamanPerBulan(pokokPerBulan float64, bungaPerBulan float64) float64 {
	return pokokPerBulan + bungaPerBulan
}

func pokokPinjamanPerBulan(pinjaman float64, lamaCicilan int) float64 {
	return pinjaman / float64(lamaCicilan)
}

func bungaPinjamanPerTahun(pinjaman float64, bunga float64) float64 {
	return pinjaman * bunga
}

func bungaPinjamanPerBulan(bungaPerTahun float64, lamaCicilan int) float64 {
	return bungaPerTahun / float64(lamaCicilan)
}

func main() {
	var pinjaman float64 = 10000000
	var lamaCicilan int = 12 // dalam bulan
	var bunga float64 = 0.1  // bunga per-tahun, persen dalam desimal

	var pokokPinjamanBulanan float64 = pokokPinjamanPerBulan(pinjaman, lamaCicilan)
	var bungaPinjamanTahunan float64 = bungaPinjamanPerTahun(pinjaman, bunga)
	var bungaPinjamanBulanan float64 = bungaPinjamanPerBulan(bungaPinjamanTahunan, lamaCicilan)
	var cicilanPinjamanBulanan float64 = cicilanPinjamanPerBulan(pokokPinjamanBulanan, bungaPinjamanBulanan)

	fmt.Printf("Pokok pinjaman per-bulan = Rp. %.2f \n", pokokPinjamanBulanan)
	fmt.Printf("Bunga pinjaman per-bulan = Rp. %.2f \n", bungaPinjamanBulanan)
	fmt.Printf("Cicilan pinjaman per-bulan = Rp. %.2f \n ", cicilanPinjamanBulanan)
}
