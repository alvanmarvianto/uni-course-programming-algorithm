package main

import "fmt"

// SOAL 1
func inputBilangan_1302223074(bil *int) {
	fmt.Scan(bil)
	for *bil < 0 {
		fmt.Scan(bil)
	}
}

func stop_1302223074(bil int) bool {
	return bil == 0
}

func hitung_1302223074(mean *float64, min, max, n *int) {
	var bil int
	inputBilangan_1302223074(&bil)
	*min = bil
	for !stop_1302223074(bil) {
		*mean += float64(bil)
		*n++
		if bil > *max {
			*max = bil
		}
		if bil < *min {
			*min = bil
		}
		inputBilangan_1302223074(&bil)
	}
}

func soal1() {
	var rata float64
	var kecil, besar, angka int
	hitung_1302223074(&rata, &kecil, &besar, &angka)
	if besar == 0 {
		fmt.Print("- - - -")
	} else {
		fmt.Print(rata/float64(angka), besar, kecil, angka)
	}
}

// SOAL 2
func average_1302223074(bil, i int, hasil *float64) {
	var angka int
	if bil != 0 {
		angka += bil % 10
		*hasil += float64(angka)
		average_1302223074(bil/10, i+1, hasil)
	} else {
		fmt.Print(*hasil / float64(i))
	}
}

func soal2() {
	var n, i int
	var total float64
	fmt.Scan(&n)
	average_1302223074(n, i, &total)
}

func main() {
	fmt.Println("SOAL 1")
	soal1()
	fmt.Println("\nSOAL 2")
	soal2()
}
