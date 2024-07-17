package main

import "fmt"

func soal1() {
	var total_belanja int
	var status_membership string
	fmt.Scan(&total_belanja, &status_membership)
	if status_membership == "Yes" {
		if total_belanja > 100000 {
			fmt.Println("Anda mendapat discount 5%, tambahan poin 10, dan free gift")
		} else if total_belanja > 75000 {
			fmt.Println("Anda mendapat discount 5% dan tambahan poin 5")
		} else {
			fmt.Println("Anda mendapat tambahan poin 5")
		}
	}
}

func soal2() {
	var total1, total2, total3 int
	var predicate string
	total1 = 0
	total2 = 0
	total3 = 0
	for predicate != "STOP" {
		fmt.Scan(&predicate)
		if predicate == "Sufficient" {
			total1++
		}
		if predicate == "Good" {
			total2++
		}
		if predicate == "Excellent" {
			total3++
		}
	}
	fmt.Println("Total siswa dengan predikat Sufficient =", total1)
	fmt.Println("Total siswa dengan predikat Good =", total2)
	fmt.Println("Total siswa dengan predikat Excellent =", total3)
}

func soal3() {
	var g string
	var jr, jl, gaji, total_gaji int
	for g != "Z" {
		fmt.Scan(&g, &jr, &jl)
		if g == "A" {
			gaji += (jr * 75000) + (jl * 90000)
			fmt.Println(gaji)
		} else if g == "B" {
			gaji += (jr * 125000) + (jl * 70000)
			fmt.Println(gaji)
		} else {
			fmt.Println("Golongan tidak dikenali")
		}
		total_gaji += gaji
		gaji = 0
	}
	fmt.Println("Biaya yang dikeluarkan perusahaan untuk gaji karyawan", total_gaji)
}

func soal4() {
	var harian, total, rata, tinggi, rendah, n int
	fmt.Scan(&harian)
	rendah = harian
	for harian > 0 {
		total += harian
		if harian > tinggi {
			tinggi = harian
		}
		if harian < rendah {
			rendah = harian
		}
		fmt.Scan(&harian)
		n++
	}
	rata = total / n
	fmt.Println("Jumlah =", total)
	fmt.Println("Rata - rata =", rata)
	fmt.Println("Uang Tertinggi =", tinggi)
	fmt.Println("Uang Terendah =", rendah)
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
