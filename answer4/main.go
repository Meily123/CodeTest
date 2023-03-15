package main

import (
	"fmt"
	"sort"
)

/*
Question :
Ayu, Budi dan Cinta merupakan kakak beradik. Budi 2 tahun diatas Cinta, sedangkan
Ayu 5 tahun dibawah Cinta. Umur Ayah dapat diketahui dengan menambahkan semua
umur anaknya. Sedangkan Ibu berumur 3 kali lebih besar dari Ayu, setelah itu
ditambah 2.
a. Berapa umur cinta,jika ayu berumur 10 tahun?
b. Berapa umur budi dan ibu,jika ayu berumur 10 tahun?
c. Berapa umur ayah,jika ayu berumur 10 tahun?
d. Berdasarkan perhitungan diatas,urutkan berdasarkan anak tertua hingga
termuda?

Hasil
Dari pernyataan diatas didapatkan persamaan:
budi =  cinta + 2
ayu = cinta - 5
ayah = budi + ayu + cinta
ibu = (ayu * 3) + 2

Dari persamaan tersebut didapatkan pula:
budi = ayu + 7
cinta = ayu + 5
*/

func main() {
	var ayu = 10
	var budi = ayu + 7
	var cinta = ayu + 5
	var ayah = budi + ayu + cinta
	var ibu = (ayu * 3) + 2

	var anak = map[string]int{
		"budi":  budi,
		"ayu":   ayu,
		"cinta": cinta,
	}

	kumpulanNama := make([]string, 0, len(anak))

	for nama, _ := range anak {
		kumpulanNama = append(kumpulanNama, nama)
	}

	sort.SliceStable(kumpulanNama, func(i, j int) bool {
		return anak[kumpulanNama[i]] > anak[kumpulanNama[j]]
	})

	fmt.Printf("Umur cinta,jika ayu berumur 10 tahun adalah %d \n", cinta)
	fmt.Printf("Umur budi dan ibu,jika ayu berumur 10 tahun adalah %d dan %d \n", budi, ibu)
	fmt.Printf("umur ayah,jika ayu berumur 10 tahun adalah %d \n", ayah)
	fmt.Printf("Urutan berdasarkan anak tertua hingga termuda adalah ")

	for indeks, nama := range kumpulanNama {
		if indeks == len(kumpulanNama)-1 { // berada pada indeks terakhir dari slice
			fmt.Printf("dan %s \n", nama)
		} else {
			fmt.Printf("%s, ", nama)
		}

	}

}
