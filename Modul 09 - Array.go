package main

import (
	"fmt"
	"math"
)

// SOAL 1
type point struct {
	x, y float64
}

func jarak_1302223074(p1, p2 point) float64 {
	return math.Sqrt((p1.x-p2.x)*(p1.x-p2.x) + (p1.y-p2.y)*(p1.y-p2.y))
}

func soal1() {
	var k1, k2, k3, k4 point
	var hasil1, hasil2 float64
	fmt.Scan(&k1.x, &k1.y, &k2.x, &k2.y)
	fmt.Scan(&k3.x, &k3.y, &k4.x, &k4.y)
	hasil1 = jarak_1302223074(k1, k2)
	hasil2 = jarak_1302223074(k3, k4)
	if hasil1 < hasil2 {
		fmt.Printf("Jarak terdekat adalah titik A(%g,%g) dengan B(%g,%g) dengan jarak %g.", k1.x, k1.y, k2.x, k2.y, hasil1)
	} else {
		fmt.Printf("Jarak terdekat adalah titik C(%g,%g) dengan D(%g,%g) dengan jarak %g.", k3.x, k3.y, k4.x, k4.y, hasil2)
	}
}

// SOAL 2
type dek [52]int

func masukkan_1302223074(kartu *dek) {
	var i int
	fmt.Scan(&kartu[0])
	for i = 1; kartu[i] != 0; i++ {
		fmt.Scan(kartu[i])
	}
}

func soal2() {
	var kartu dek
	var i, n, temp, panjang int
	masukkan_1302223074(&kartu)
	n = len(kartu) / 2
	panjang = len(kartu) - 1
	for i = 0; i < n; i++ {
		temp = kartu[i]
		kartu[i] = kartu[panjang]
		kartu[panjang] = temp
		panjang--
	}
	fmt.Print(kartu)
}

// SOAL 3
func inputNilai_1302223074() []int {
	i, angka, counter := 1, 0, 0
	var output = []int{0}
	fmt.Scanln(&angka)
	counter += angka
	output[0] = angka
	for len(output) <= 52 && angka != 0 {
		fmt.Scanln(&angka)
		if angka != 0 {
			output = append(output, angka)
			counter += angka
			i++
		}
	}
	fmt.Println(i, counter, "end", output)
	return output[:]

}
func reverse_1302223074(input ...int) []int {
	j, output := 0, make([]int, len(input))
	for i := len(input) - 1; i >= 0; i-- {
		output[j] = input[i]
		j++
	}
	return output
}
func soal3() {
	output := reverse_1302223074(inputNilai_1302223074()...)
	if output[0] != 0 {
		fmt.Println(output)
	} else {
		fmt.Println()
	}
}

func main() {
	fmt.Println("SOAL 1")
	soal1()
	fmt.Println("\nSOAL 2")
	soal2()
	fmt.Println("\nSOAL 3")
	soal3()
}
