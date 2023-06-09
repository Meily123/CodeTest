package main

import "fmt"

/*
Question :
Alan merupakan seorang dosen di STMIK IB,dia ingin memberikan nilai kepada
apip,dengan nilai Absensi =16 x hadir,nilai tugas harian = 70,nilai uts =80 dan nilai
uas = 70,dengan bobot nilai Absensi = 10%(untuk 18 x hadir),Nilai tugas harian
=20%,nilai uts = 30 % dan nilai uas =40%
*/

func nilaiBobot(nilai float64, persenBobot float64) float64 {
	return nilai * (persenBobot / 100)
}

func main() {
	var persenBobotTugasHarian float64 = 20
	var persenBobotUTS float64 = 30
	var persenBobotUAS float64 = 40
	var persenBobotAbsensi float64 = 10

	var totalAbsensi = 18
	var banyakAbsensi = 16

	var nilaiTugasHarian float64 = 70
	var nilaiUTS float64 = 80
	var nilaiUAS float64 = 70
	var nilaiAbsensi = (float64(banyakAbsensi) / float64(totalAbsensi)) * 100

	var nilaiBobotTugasHarian = nilaiBobot(nilaiTugasHarian, persenBobotTugasHarian)
	var nilaiBobotUTS = nilaiBobot(nilaiUTS, persenBobotUTS)
	var nilaiBobotUAS = nilaiBobot(nilaiUAS, persenBobotUAS)
	var nilaiBobotAbsensi = nilaiBobot(nilaiAbsensi, persenBobotAbsensi)

	var nilaiAkhir = nilaiBobotTugasHarian + nilaiBobotUTS + nilaiBobotUAS + nilaiBobotAbsensi

	fmt.Printf("Nilai hasil pembobotan tugas harian = %.2f \n", nilaiBobotTugasHarian)
	fmt.Printf("Nilai hasil pembobotan uts = %.2f \n", nilaiBobotUTS)
	fmt.Printf("Nilai hasil pembobotan uas = %.2f \n", nilaiBobotUAS)
	fmt.Printf("Nilai hasil pembobotan absensi = %.2f \n", nilaiBobotAbsensi)
	fmt.Printf("Jadi, nilai hasil akhir Apip = %.2f \n", nilaiAkhir)
}
