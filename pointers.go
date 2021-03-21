package main

import "fmt"

func zeroval(ival int) {
	fmt.Println("&ival:", &ival)
	ival = 0
}

func zeroptr(iptr *int) {
	*iptr = 0
}

func pointPlusSeven(i int) *int {
	newval := i + 7
	fmt.Println("infunc:", &newval)
	return &newval
}

func main() {
	i := 1
	fmt.Println("pointer:", &i)
	fmt.Println("initial:", i)
	zeroval(i)
	fmt.Println("zeroval:", i)
	zeroptr(&i)
	fmt.Println("seroptr:", i)
	fmt.Println("pointer:", &i)

	startval := 10
	pp7 := pointPlusSeven(startval)
	fmt.Println("outfunc:", pp7)
	fmt.Println("pp7:", *pp7)
}
