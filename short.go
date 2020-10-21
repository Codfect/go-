package main

import "fmt"
import "sort"

func main() {
	wutc := []string {"RZA", "GZA", "Method Man", "Ghostface Killah", "Inspectah Deck", "Ol' Dirty Bastard", "Raekwon", "Masta Killa", "U-God"}

	fmt.Println(wutc)
	
	sort.Strings(wutc)
	
	fmt.Println(wutc)
	
	
	si := []int {456, 789, 123, 777, 999}
	
	fmt.Println(si)
	
	sort.Ints(si)
	
	fmt.Println(si)
}
