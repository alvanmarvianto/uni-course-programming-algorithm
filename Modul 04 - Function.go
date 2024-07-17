package main

import (
	"fmt"
	"math"
)

// SOAL 1
func z(x, y float64) float64 {
	/* mengembalikan nilai z sesuai dengan persamaan bilangan bulat x dan y */
	return math.Sqrt(((3 * x) * (6 * y) / (4 * y)))
}

func soal1() {
	var a, b float64
	fmt.Scanln(&a, &b)
	fmt.Print(z(a, b), z(b, a))
}

// SOAL 2
func soal2() {
	var n, m int
	fmt.Scan(&n, &m)
	fmt.Println(jumlahBus(n, m), "bus")
}

func jumlahBus(penumpang, kapasitas int) int {
	/* mengembalikan jumlah Bus yang diperlukan, apabila diketahui jumlah penumpang dan kapasitas dari bus */
	var bus int
	bus = penumpang / kapasitas
	if penumpang%kapasitas > 0 {
		bus++
	}
	return bus
}

// SOAL 3
func soal3() {
	var indeks float64
	var studi int
	var status bool
	fmt.Scan(&indeks, &studi, &status)
	if cumlaude(indeks, studi, status) {
		fmt.Print("Cumlaude")
	} else if sangatMemuaskan(indeks, studi, status) {
		fmt.Print("Sangat memuaskan")
	} else if memuaskan(indeks, studi, status) {
		fmt.Print("Memuaskan")
	}

}

func cumlaude(ipk float64, masaStudi int, publikasi bool) bool {
	/* mengembalikan true apabila memenuhi kriteria "cumlaude", dan false apabila sebaliknya */
	return (ipk >= 3.51 && ipk <= 4) && masaStudi <= 8 && publikasi
}

func sangatMemuaskan(ipk float64, masaStudi int, publikasi bool) bool {
	/* mengembalikan true apabila memenuhi kriteria "Sangat memuaskan", dan false apabila sebaliknya */
	return ipk >= 2.76 && masaStudi <= 14
}

func memuaskan(ipk float64, masaStudi int, publikasi bool) bool {
	/* mengembalikan true apabila memenuhi kriteria "Memuaskan", dan false apabila sebaliknya */
	return ipk >= 2 && masaStudi <= 14
}

func main() {
	fmt.Println("SOAL 1")
	soal1()
	fmt.Println("\nSOAL 2")
	soal2()
	fmt.Println("\nSOAL 3")
	soal3()
}
