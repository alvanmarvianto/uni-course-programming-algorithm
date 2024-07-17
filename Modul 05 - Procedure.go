package main

import "fmt"

// SOAL 1
func soal1() {
	var tgl1, bln1, thn1, tgl2, bln2, thn2 int
	inputTglPinjam(&tgl1, &bln1, &thn1)
	for valid(tgl1, bln1, thn1) {
		hitungTanggalKembali(tgl1, bln1, thn1, &tgl2, &bln2, &thn2)
		fmt.Println(tgl2, bln2, thn2)
		inputTglPinjam(&tgl1, &bln1, &thn1)
	}
}

func inputTglPinjam(tanggal, bulan, tahun *int) {
	fmt.Scan(tanggal, bulan, tahun)
}

func valid(tanggal, bulan, tahun int) bool {
	var bulanTgl int
	getJumlahHari(bulan, tahun, &bulanTgl)
	return tahun > 0 && bulan >= 1 && bulan <= 12 && (0 < tanggal && bulanTgl >= tanggal)
}

func kabisat(tahun int) bool {
	return tahun%4 == 0 && (tahun%100 != 0 || tahun%400 == 0)
}

func getJumlahHari(bulan, tahun int, jmlHari *int) {
	if bulan <= 7 {
		if bulan%2 == 1 {
			*jmlHari = 31
		} else if kabisat(tahun) && bulan == 2 {
			*jmlHari = 29
		} else if bulan == 2 {
			*jmlHari = 28
		} else {
			*jmlHari = 30
		}
	} else {
		if bulan%2 == 1 {
			*jmlHari = 30
		} else {
			*jmlHari = 31
		}
	}
}

func hitungTanggalKembali(tanggal1, bulan1, tahun1 int, tanggal2, bulan2, tahun2 *int) {
	var i int
	*tanggal2 = tanggal1
	*bulan2 = bulan1
	*tahun2 = tahun1
	for i = 1; i <= 3; i++ {
		*tanggal2 += 1
		if !valid(*tanggal2, *bulan2, *tahun2) {
			*tanggal2 = 1
			*bulan2 += 1
			if !valid(*tanggal2, *bulan2, *tahun2) {
				*tanggal2 = 1
				*bulan2 = 1
				*tahun2 += 1
			}
		}
	}
}

// SOAL 2
func soal2() {
	var modal, saldo, pengeluaran float64
	fmt.Scan(&modal)
	membeliKain_1302223074(modal, &saldo, &pengeluaran)
}

func membeliKain_1302223074(uangAwal float64, tUang, tPengeluaran *float64) {
	*tPengeluaran = uangAwal / 4
	*tUang = uangAwal - uangAwal/4
	membeliBenangJahit_1302223074(uangAwal, tUang, tPengeluaran)
}

func membeliBenangJahit_1302223074(uangAwal float64, tUang, tPengeluaran *float64) {
	*tPengeluaran += uangAwal / 20
	*tUang -= uangAwal / 20
	membuatPolaBaju_1302223074(uangAwal, tUang, tPengeluaran)
}

func membuatPolaBaju_1302223074(uangAwal float64, tUang, tPengeluaran *float64) {
	*tPengeluaran += (uangAwal / 200) * 2
	*tUang -= (uangAwal / 200) * 2
	menjahitBaju_1302223074(uangAwal, tUang, tPengeluaran)
}

func menjahitBaju_1302223074(uangAwal float64, tUang, tPengeluaran *float64) {
	*tPengeluaran += (uangAwal / 200) * 4
	*tUang -= (uangAwal / 200) * 4
	mengemasBaju_1302223074(uangAwal, tUang, tPengeluaran)
}

func mengemasBaju_1302223074(uangAwal float64, tUang, tPengeluaran *float64) {
	var tPemasukan float64
	*tPengeluaran += (3 * uangAwal / 200) * 2
	*tUang -= (3 * uangAwal / 200) * 2
	mendistribusikan_1302223074(uangAwal, tUang, &tPemasukan, tPengeluaran)
}

func mendistribusikan_1302223074(uangAwal float64, tUang, tPemasukan, tPengeluaran *float64) {
	*tPengeluaran += 3 * uangAwal / 200
	*tPemasukan += uangAwal / 2
	*tUang -= 3*uangAwal/200 + uangAwal/2
	fmt.Printf("%.0f %.0f %.0f", *tPengeluaran, *tPemasukan, *tUang+uangAwal)
}

// SOAL 3
func soal3() {
	var modalAwal, saldo, pengeluaran, pemasukan float64
	fmt.Scan(&modalAwal)
	membeliKain(modalAwal, &saldo, &pengeluaran)
	membeliBenangJahit(modalAwal, &saldo, &pengeluaran)
	membuatPolaBaju(modalAwal, &saldo, &pengeluaran)
	menjahitBaju(modalAwal, &saldo, &pengeluaran)
	mengemasBaju(modalAwal, &saldo, &pengeluaran)
	mendistribusikan(modalAwal, &saldo, &pemasukan, &pengeluaran)
	fmt.Printf("%.0f %.0f %.0f", pengeluaran, pemasukan, pemasukan+modalAwal-pengeluaran)
}

func membeliKain(uangAwal float64, tUang, tPengeluaran *float64) {
	*tPengeluaran = uangAwal / 4
	*tUang = uangAwal - uangAwal/4
}

func membeliBenangJahit(uangAwal float64, tUang, tPengeluaran *float64) {
	*tPengeluaran += uangAwal / 20
	*tUang -= uangAwal / 20
}

func membuatPolaBaju(uangAwal float64, tUang, tPengeluaran *float64) {
	*tPengeluaran += (uangAwal / 200) * 2
	*tUang -= (uangAwal / 200) * 2
}

func menjahitBaju(uangAwal float64, tUang, tPengeluaran *float64) {
	*tPengeluaran += (uangAwal / 200) * 4
	*tUang -= (uangAwal / 200) * 4
}

func mengemasBaju(uangAwal float64, tUang, tPengeluaran *float64) {
	*tPengeluaran += (3 * uangAwal / 200) * 2
	*tUang -= (3 * uangAwal / 200) * 2
}

func mendistribusikan(uangAwal float64, tUang, tPemasukan, tPengeluaran *float64) {
	*tPengeluaran += 3 * uangAwal / 200
	*tPemasukan += uangAwal / 2
	*tUang -= 3*uangAwal/200 + uangAwal/2
}

func main() {
	fmt.Println("SOAL 1")
	soal1()
	fmt.Println("\nSOAL 2")
	soal2()
	fmt.Println("\nSOAL 3")
	soal3()
}
