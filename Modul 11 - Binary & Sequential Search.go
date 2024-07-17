package main

import "fmt"

const NMAX int = 1000

type himpunan struct {
	info    [NMAX]string
	nElemen int
}

func main() {
	var A, B, C, D himpunan
	createSet_1302223074(&A)
	createSet_1302223074(&B)
	intersection_1302223074(A, B, &C)
	union_1302223074(A, B, &D)
	printSet_1302223074(C)
	printSet_1302223074(D)
}

func createSet_1302223074(set *himpunan) {
	var s string
	fmt.Scan(&s)
	for !isMember_1302223074(*set, s) {
		set.info[set.nElemen] = s
		set.nElemen++
		fmt.Scan(&s)
	}
}

func printSet_1302223074(set himpunan) {
	for i := 0; i <= set.nElemen; i++ {
		fmt.Print(set.info[i], " ")
	}
	fmt.Println()
}

func isMember_1302223074(set himpunan, s string) bool {
	for i := 0; i <= set.nElemen; i++ {
		if set.info[i] == s {
			return true
		}
	}
	return false
}

func intersection_1302223074(set1, set2 himpunan, set3 *himpunan) {
	for i := 0; i <= set1.nElemen; i++ {
		if isMember_1302223074(set2, set1.info[i]) {
			set3.info[set3.nElemen] = set1.info[i]
			set3.nElemen++
		}
	}
}

func union_1302223074(set1, set2 himpunan, set3 *himpunan) {
	*set3 = set1
	n := set3.nElemen
	for i := 0; i <= set2.nElemen; i++ {
		if !isMember_1302223074(*set3, set2.info[i]) {
			set3.info[n] = set2.info[i]
			set3.nElemen = n
			n++
		}
	}
}
