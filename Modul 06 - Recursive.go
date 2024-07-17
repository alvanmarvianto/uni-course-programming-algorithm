package main

import "fmt"

// SOAL 1
func abc(x int) string {
	var a int
	var s string
	if x != 0 {
		a = x % 2
		if a == 0 {
			return s + "0" + abc(x/2)
		} else {
			return s + "1" + abc(x/2)
		}
	}
	return ""
}

func soal1() {
	var n int
	fmt.Scan(&n)
	fmt.Println(abc(n))
}

// SOAL 2
func beliBuku(N, M int) {
	var B int
	B = N - M
	fmt.Println("Sisa rak kosong", B, "buku")
	for B != 1 {
		B--
		fmt.Println("beli 1 buku, sisa", B)
	}
	fmt.Println("beli 1 buku, rak buku penuh")
}

func soal2() {
	var n, m int
	fmt.Scan(&n, &m)
	beliBuku(n, m)
}

// SOAL 3

func pola(n int, s *string) {
	if n != 0 {
		*s = *s + "*"
		fmt.Println(*s)
		pola(n-1, s)
	}
}

func soal3() {
	var n int
	var s string
	fmt.Scan(&n)
	pola(n, &s)
}

func main() {
	fmt.Println("SOAL 1")
	soal1()
	fmt.Println("\nSOAL 2")
	soal2()
	fmt.Println("\nSOAL 3")
	soal3()
}
