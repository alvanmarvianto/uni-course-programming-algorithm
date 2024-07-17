package main

import (
	"fmt"
	"math"
)

func soal1() {
	var x int
	fmt.Scan(&x)
	if (x%2 == 0 && x%3 == 0) || (x%3 == 0 && x%5 == 0) {
		fmt.Print("bilangan kelipatan 2 dan 3, atau kelipatan 3 dan 5")
	} else {
		fmt.Print("BUKAN kelipatan 2 dan 3, juga BUKAN kelipatan 3 dan 5")
	}
}

func soal2() {
	var i, n_passed, n_failed, n int
	var n1, n2, n3, avg float64
	fmt.Print("Berapa jumlah siswa yang nilainya akan diproses?")
	fmt.Scan(&n)
	n_passed = 0
	n_failed = 0
	for i = 0; i < n; i++ {
		fmt.Scan(&n1, &n2, &n3)
		avg = (n1 + n2 + n3) / 3
		if avg > 80.0 {
			fmt.Println("Memenuhi syarat administratif")
			n_passed++
		} else {
			fmt.Println("Tidak memenuhi syarat administratif")
			n_failed++
		}
	}
	fmt.Println("Jumlah siswa lolos seleksi admistrasi ", n_passed)
	fmt.Println("Jumlah siswa tidak lolos seleksi admistrasi ", n_failed)
}

func soal3() {
	var p, l, luas, keliling, diagonal float64
	fmt.Scan(&p, &l)
	luas = p * l
	keliling = (p * 2.0) + (l * 2.0)
	diagonal = math.Sqrt((p * p) + (l * l))
	fmt.Println("luas: ", luas, "\nkeliling: ", keliling, "\npanjang diagonal: ", diagonal)
}

func soal4() {
	var disc, harga, total, thn int
	fmt.Scan(&thn, &harga)
	disc = (thn / 1000) * (((thn / 10) % 10 * 10) + (thn % 10))
	total = harga - (harga * disc / 100)
	fmt.Print("besar diskon: ", disc, "%\nJumlah yang dibayar: ", total)
}

func main() {
	fmt.Println("SOAL 1")
	soal1()
	fmt.Println("\nSOAL 2")
	soal2()
	fmt.Println("\nSOAL 3")
	soal3()
	fmt.Println("\nSOAL 4")
	soal4()
}
