package dsarecursion

import "fmt"

func Triangle(row, coloumn int) {
	if row == 0 {
		return
	}
	if coloumn < row {
		Triangle(row, coloumn+1)
		fmt.Print("*")
	} else {
		Triangle(row-1, 0)
		fmt.Println()
	}
}

//For Reverse Triangle

func Triangle2(row, coloumn int) {
	if row == 0 {
		return
	}
	if coloumn < row {
		fmt.Print("*")
		Triangle2(row, coloumn+1)
	} else {
		fmt.Println()
		Triangle2(row-1, 0)
	}
}
