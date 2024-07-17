package main

import "fmt"

type student struct {
	name                            string
	math, indo, eng, sains, average float64
}

type arrStudent [2048]student

func main() {
	var T arrStudent
	var n int
	entryStudent(&T, &n)
	calculateAverage(&T, n)
	rangking(&T, n)
	printTop3(T, n)
	printMax(T, n)
}

func entryStudent(T *arrStudent, n *int) {
	var name string
	fmt.Scan(&name)
	for name != "STOP" {
		T[*n].name = name
		fmt.Scan(&T[*n].math, &T[*n].indo, &T[*n].eng, &T[*n].sains)
		*n++
		fmt.Scan(&name)
	}
}

func calculateAverage(T *arrStudent, n int) {
	var total float64
	for i := 0; i < n; i++ {
		total = T[i].math + T[i].indo + T[i].eng + T[i].sains
		T[i].average = total / 4
	}
}

func max(T arrStudent, n int, flag string) int {
	var max, temp float64
	var max_idx int
	for i := 0; i < n; i++ {
		if flag == "math" {
			temp = T[i].math
		} else if flag == "indo" {
			temp = T[i].indo
		} else if flag == "eng" {
			temp = T[i].eng
		} else {
			temp = T[i].sains
		}
		if temp > max {
			max = temp
			max_idx = i
		}
	}
	return max_idx
}

func rangking(T *arrStudent, n int) {
	var i, idx_max, j int
	for i = 0; i < n-1; i++ {
		idx_max = i
		for j = i + 1; j < n; j++ {
			if T[idx_max].average < T[j].average {
				idx_max = j
			}
		}
		T[idx_max], T[i] = T[i], T[idx_max]
	}
}

func printTop3(T arrStudent, n int) {
	for i := 0; i < 3; i++ {
		fmt.Printf("Rangking %d: %s rata-rata %.1f\n", i+1, T[i].name, T[i].average)
	}
}

func printMax(T arrStudent, n int) {
	fmt.Printf("Nilai Matematika tertinggi diraih oleh %s\n", T[max(T, n, "math")].name)
	fmt.Printf("Nilai Bahasa Indonesia tertinggi diraih oleh %s\n", T[max(T, n, "indo")].name)
	fmt.Printf("Nilai Bahasa Inggris tertinggi diraih oleh %s\n", T[max(T, n, "eng")].name)
	fmt.Printf("Nilai pelajaran IPA/IPS tertinggi diraih oleh %s\n", T[max(T, n, "")].name)
}
