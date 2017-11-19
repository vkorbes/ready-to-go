package main

import (
	"fmt"
	"sort"
)

func main() {

	si := []int{123, 987, 324, 876, 234, 987, 234, 76}

	fmt.Println(si)

	sort.Ints(si)

	fmt.Println(si)

}
