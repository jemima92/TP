package main

import (
	"b2"
	"fmt"
)

/*
func main() {
	fmt.Println(b2.Ft_coin([]int{1, 2, 5}, 11))
	fmt.Println(b2.Ft_coin([]int{2}, 3))
}
*/
/*
func main() {
	fmt.Println(b2.Ft_missing([]int{3, 1, 2}))
	fmt.Println(b2.Ft_missing([]int{0, 1}))
	fmt.Println(b2.Ft_missing([]int{9, 6, 4, 2, 3, 5, 7, 0, 1}))
}*/
func main() {

	fmt.Println(b2.Ft_non_overlap([][]int{{1, 2}, {2, 3}, {3, 4}, {1, 3}}))
	fmt.Println(b2.Ft_non_overlap([][]int{{1, 2}, {2, 3}}))
	fmt.Println(b2.Ft_non_overlap([][]int{{1, 2}, {1, 2}, {1, 2}}))
}

/*
func main() {
	println(b2.Ft_profit([]int{7, 1, 5, 3, 6, 4})) // resultat : 5
	println(b2.Ft_profit([]int{7, 6, 4, 3, 1})) // resultat : 0
}

*/
