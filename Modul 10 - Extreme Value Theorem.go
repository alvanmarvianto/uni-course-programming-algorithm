package main

import "fmt"

const NMAX int = 12345

type prodi struct {
	nama, akreditasi      string
	tahun, aktif, lulusan int
}

type fakultas struct {
	nama     string
	arrProdi [NMAX - 1]prodi
	N        int
}

func main() {
	var f fakultas
	var banyak prodi
	var hasilBanyak int
	buatFakultas_1302223074(&f)
	daftarProdi_1302223074(&f)
	banyak = terbanyak_1302223074(f)
	hasilBanyak = banyak.aktif + banyak.lulusan
	fmt.Printf("Prodi %s memiliki mahasiswa lulusan terbanyak yaitu %d\n", banyak.nama, hasilBanyak)
	fmt.Println("Prodi termuda adalah", f.arrProdi[termuda_1302223074(f)].nama)
	fmt.Printf("Akreditasi Prodi terbanyak di Fakultas %s adalah %s\n", f.nama, prestasi_1302223074(f))
	cetakProdi_1302223074(f, prestasi_1302223074(f))
}

func buatFakultas_1302223074(f *fakultas) {
	fmt.Scan(&f.nama)
	f.N = 0
}

func daftarProdi_1302223074(f *fakultas) {
	var tahun, aktif, lulus, j, n int
	var nama, akre string
	for i := 0; i < 10; i++ {
		fmt.Scan(&nama, &akre, &tahun, &aktif, &lulus)
		j = cekProdi_1302223074(nama, *f)
		if j != -1 {
			f.arrProdi[j].aktif += aktif
			f.arrProdi[j].lulusan += lulus
		} else {
			f.arrProdi[n].nama = nama
			f.arrProdi[n].akreditasi = akre
			f.arrProdi[n].tahun = tahun
			f.arrProdi[n].aktif = aktif
			f.arrProdi[n].lulusan = lulus
			n++
		}
	}
	f.N = n
}

func cekProdi_1302223074(s string, f fakultas) int {
	var j int
	for j < f.N && f.arrProdi[j].nama != s {
		j++
	}
	if j == f.N {
		return -1
	} else {
		return j
	}

}

func terbanyak_1302223074(f fakultas) prodi {
	var i, jumlah, max, n int
	for i = 0; i < f.N; i++ {
		jumlah = f.arrProdi[i].aktif + f.arrProdi[i].lulusan
		if jumlah > max {
			max = jumlah
			n = i
		}
	}
	return f.arrProdi[n]
}

func termuda_1302223074(f fakultas) int {
	var i, max, prodi int
	for i = 0; i < f.N; i++ {
		if f.arrProdi[i].tahun >= max {
			max = f.arrProdi[i].tahun
			prodi = i
		}
	}
	return prodi
}

func prestasi_1302223074(f fakultas) string {
	var unggul, baik, cukup, i int
	var akre string
	for i = 0; i < f.N; i++ {
		akre = f.arrProdi[i].akreditasi
		if akre == "unggul" {
			unggul++
		} else if akre == "baik" {
			baik++
		} else if akre == "cukup" {
			cukup++
		}
	}
	if unggul > baik && unggul > cukup {
		return "unggul"
	} else if baik > unggul && baik > cukup {
		return "baik"
	} else {
		return "cukup"
	}
}

func cetakProdi_1302223074(f fakultas, x string) {
	var i int
	for i = 0; i < f.N; i++ {
		if x == f.arrProdi[i].akreditasi {
			fmt.Print(f.arrProdi[i].nama, " ")
		}
	}
}
