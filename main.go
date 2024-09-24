package main

import (
	exo1 "exo/exo1"
	exo2 "exo/exo2"
	exo3 "exo/exo3"
	exo4 "exo/exo4"
	exo5 "exo/exo5"
	exo6 "exo/exo6"
	"fmt"
)

func main() {
	fmt.Println("exercice 1")
	fmt.Println(exo1.Ft_coin([]int{1, 2, 5}, 11)) // resultat :  3
	fmt.Println(exo1.Ft_coin([]int{2}, 3))        // resultat : -1
	fmt.Println(exo1.Ft_coin([]int{1}, 0))        // resultat :  0

	fmt.Println("exercice 2")
	fmt.Println(exo2.Ft_missing([]int{3, 1, 2}))                   // resultat : 0
	fmt.Println(exo2.Ft_missing([]int{0, 1}))                      // resultat : 2
	fmt.Println(exo2.Ft_missing([]int{9, 6, 4, 2, 3, 5, 7, 0, 1})) // resultat : 8

	fmt.Println("exercice 3")
	fmt.Println(exo3.Ft_non_overlap([][]int{{1, 2}, {2, 3}, {3, 4}, {1, 3}})) // resultat : 1
	fmt.Println(exo3.Ft_non_overlap([][]int{{1, 2}, {2, 3}}))                 // resultat : 0
	fmt.Println(exo3.Ft_non_overlap([][]int{{1, 2}, {1, 2}, {1, 2}}))         // resultat : 2

	fmt.Println("exercice 4")
	fmt.Println(exo4.Ft_profit([]int{7, 1, 5, 3, 6, 4})) // resultat : 5
	fmt.Println(exo4.Ft_profit([]int{7, 6, 4, 3, 1}))    // resultat : 0

	fmt.Println("exercice 5")
	fmt.Println(exo5.Ft_max_substring("abcabcbb")) // resultat : 3
	fmt.Println(exo5.Ft_max_substring("bbbbb"))    // resultat : 1

	fmt.Println("exercice 6")
	fmt.Println(exo6.Ft_min_window("ADOBECODEBANC", "ABC")) // resultat : "BANC"
	fmt.Println(exo6.Ft_min_window("a", "aa"))              // resultat : ""

}
