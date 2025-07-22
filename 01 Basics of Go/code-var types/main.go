package main

import "fmt"

func main() {

	// var age int
	// var name string
	// age = 21
	// name = "kunal"
	var age = 21
	var name = "kunal"

	fmt.Println(name, age)
	fmt.Println()
/* Zero values/default:
	The zero values ensure your program behaves predictably, 
	even if variables are not explicitly initialized.
*/
	var i int         // 0
	var f float32     // 0.0
	var b bool		  // false
	var s string      // ""
	var arr [3]int    // [0, 0, 0]
	var slc []int     // nil (uninitialized slice)
	var m map[string]int // nil (uninitialized map)
	var p *int  // nil (pointer to int)

	fmt.Println(i, f, b, s, arr, slc, m, p)
	fmt.Println()

	// shorthand declaration
	// type inference: use for local var
	count := 50		// type inferred as int
	msg := "Type inference in Go"	// type inferred as string
	isActive := true	// type inferred as bool

	// Explicit declaration: use for global var, need to specify type
	var version string = "1.24.5"

	fmt.Println(count, msg, isActive, version)
	fmt.Println()
	// arrays and slices
	arr2 := [3]int{1, 2, 3}	// array with 3 elements
	slc2 := []int{1, 2, 3, 4} // slice of integers
	fmt.Println(arr2, slc2)
}
