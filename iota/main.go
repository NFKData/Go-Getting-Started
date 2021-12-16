package main

import "fmt"

const PI = 3.14

const (
	FIRST = iota
	SECOND
)

const (
	THIRD = iota
	FOURTH
)

const (
	FIRST_EXPR = iota + 3
	SECOND_EXPR
	THIRD_EXPR = iota * 2
	FOURTH_EXPR
	FIFTH_EXPR
	SIXTH_EXPR = iota / 2
	SEVENTH_EXPR
	EIGHT_EXPR
	NINTH_EXPR
	TENTH_EXPR = iota * 0
	ELEVENTH_EXPR
)

func main() {
	fmt.Println(PI)
	fmt.Println(FIRST, SECOND, THIRD, FOURTH)
	fmt.Println(FIRST_EXPR, SECOND_EXPR, THIRD_EXPR, FOURTH_EXPR, FIFTH_EXPR, SIXTH_EXPR, SEVENTH_EXPR, EIGHT_EXPR, NINTH_EXPR, TENTH_EXPR, ELEVENTH_EXPR)
}
