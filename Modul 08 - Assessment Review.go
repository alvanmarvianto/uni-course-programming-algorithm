package main

import (
	"fmt"
	"math"
)

// SOAL 1
const phi float64 = 3.14

func luasAlas_1302223074(r int) float64 {
	return phi * float64(r*r)
}

func garisPelukis_1302223074(r, t int) float64 {
	return math.Sqrt(float64(r*r) + float64(t*t))
}

func hitungLuasSelimut_1302223074(r, t int, luas *float64) {
	*luas = phi * float64(r) * garisPelukis_1302223074(r, t)
}

func soal1() {
	var R, r, t, t1 int
	var luasSBesar, luasSKecil, luasPermukaan float64
	fmt.Scan(&R, &r, &t, &t1)
	luasAlas_1302223074(R)
	hitungLuasSelimut_1302223074(R, t, &luasSBesar)
	hitungLuasSelimut_1302223074(r, t1, &luasSKecil)
	luasPermukaan = luasSBesar - luasSKecil + luasAlas_1302223074(R) + luasAlas_1302223074(r)
	fmt.Println(luasSBesar)
	fmt.Println(luasSKecil)
	fmt.Println(luasPermukaan)
}

// SOAL 2
func durasi_1302223074(jam, menit int) int {
	if jam == 0 && menit >= 1 {
		return 1
	} else if jam >= 1 {
		if menit >= 10 {
			return jam + 1
		} else {
			return jam
		}
	} else {
		return 0
	}
}

func potongan_1302223074(durasi, tarif int) float64 {
	if durasi > 3 {
		return float64(durasi*tarif) / 0.1
	} else {
		return 0
	}
}

func tarif_1302223074(member bool) int {
	if member {
		return 3500
	} else {
		return 5000
	}
}

func hitungSewa_1302223074(jam, menit int, member bool, biaya *float64) {
	var lamaDurasi, tarifSewa int
	var diskon float64
	lamaDurasi = durasi_1302223074(jam, menit)
	tarifSewa = tarif_1302223074(member)
	diskon = potongan_1302223074(lamaDurasi, tarifSewa)
	*biaya = float64(lamaDurasi*tarifSewa) - diskon
}

func soal2() {
	var jam, menit int
	var member bool
	var total float64
	fmt.Scan(&jam, &menit, &member)
	hitungSewa_1302223074(jam, menit, member, &total)
	fmt.Print(total)
}

// SOAL 3
func isFaktor_1302223074(num, x int) bool {
	return num%x == 0
}

func jumlahFaktor_1302223074(num int, total *int) {
	var i int
	*total = 0
	for i = 1; i < num; i++ {
		if isFaktor_1302223074(num, i) {
			*total += i
		}
	}
}

func perfect_1302223074(num int) bool {
	var total int
	jumlahFaktor_1302223074(num, &total)
	return num == total
}

func display_1302223074(x, y int) {
	var i int
	fmt.Print("Barisan perfect number: ")
	for i = x; i <= y; i++ {
		if perfect_1302223074(i) {
			fmt.Print(i, " ")
		}
	}
}

func soal3() {
	var a, b int
	fmt.Scan(&a, &b)
	display_1302223074(a, b)
}

// SOAL 4
func len_1302223084(num int) int {
	var total int
	for num > 0 {
		num = num / 10
		total++
	}
	return total
}

func pangkat_1302223084(n int) int {
	if n == 0 {
		return 1
	} else {
		return 10 * pangkat_1302223084(n-1)
	}
}

func split_1302223084(num int, num1, num2 *int) {
	var mid int = len_1302223084(num) / 2
	*num1 = num / pangkat_1302223084((mid))
	*num2 = num % pangkat_1302223084((mid))
}

func soal4() {
	var number, number1, number2 int
	fmt.Scan(&number)
	split_1302223084(number, &number1, &number2)
	fmt.Println(number1, number2)
	fmt.Println(number1 + number2)
}

// SOAL 5
func tarif2_1302223074(menu int) int {
	if menu <= 3 {
		return 10000
	} else if menu > 3 && menu <= 50 {
		return 10000 + (menu-3)*2500
	} else {
		return 100000
	}
}

func hitungTarif2_1302223074(menu int, bersisa bool, n int, total *int) {
	if bersisa {
		*total = tarif2_1302223074(menu) * n
	} else {
		*total = tarif2_1302223074(menu)
	}
}

func soal5() {
	var M, i, menu, orang, total int
	var sisa bool
	fmt.Scan(&M)
	for i = 0; i < M; i++ {
		fmt.Scan(&menu, &orang, &sisa)
		hitungTarif2_1302223074(menu, sisa, orang, &total)
		fmt.Println(total)
	}
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
	fmt.Println("\nSOAL 5")
	soal5()
}
