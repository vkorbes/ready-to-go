package main

import (
	"fmt"
	"sort"
)

func main() {

	ss := []string{"abóbora", "maçã", "laranja", "beringela", "berinjela"}

	fmt.Println(ss)

	sort.Strings(ss)

	fmt.Println(ss)

}
