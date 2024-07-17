package main

import "fmt"

const NMAX = 100

type ASEAN struct {
	GDP  [3]int
	nama string
}

type tab_Asean [10]ASEAN

type negara [10]string

func input_negara(A *tab_Asean, list negara) {
	for i := 0; i < 10; i++ {
		A[i].nama = list[i]
	}

}
func input_1302220068(A *tab_Asean) {
	var negara string
	var tahun, idxN, gdp, idxT int
	list := [10]string{"Brunei", "Cambodia", "Indonesia", "Laos", "Malaysia",
		"Myanmar", "Philippines", "Singapore", "Thailand", "Vietnam"}
	input_negara(A, list)
	fmt.Scan(&negara, &tahun)
	idxN = is_asean(*A, negara)
	idxT = find_tahun(tahun)
	for idxN != -1 {
		fmt.Scan(&gdp)
		A[idxN].GDP[idxT] = gdp
		fmt.Scan(&negara, &tahun)
		idxN = is_asean(*A, negara)
		idxT = find_tahun(tahun)
	}
	sort(A, idxT)
}

func print(A tab_Asean) {
	for i := 0; i < 10; i++ {
		fmt.Println(A[i].nama, A[i].GDP[0], A[i].GDP[1], A[i].GDP[2])
	}
}

func is_asean(A tab_Asean, negara string) int {
	for i := 0; i < 10; i++ {
		if A[i].nama == negara {
			return i
		}
	}
	return -1
}

func find_tahun(tahun int) int {
	if tahun == 2021 {
		return 0
	} else if tahun == 2022 {
		return 1
	} else {
		return 2
	}
}

func sort(A *tab_Asean, n int) {
	var idx, i, j int
	for i = 1; i <= 10-1; i++ {
		idx = i - 1
		for j = i; j < 10; j++ {
			if A[idx].GDP[n] <= A[j].GDP[n] {
				idx = j
			}
		}
		A[idx], A[i-1] = A[i-1], A[idx]
	}
}

func main() {
	var A tab_Asean
	input_1302220068(&A)
	print(A)
}
