package main

import (
	"fmt"

	"repo.1/pkg1"
	"repo.2/pkg2"
	"repo.3/pkg3"
	"repo.4/pkg4"
)

func main() {
	fmt.Println(pkg1.Location())
	fmt.Println(pkg2.Location())
	fmt.Println(pkg3.Location())
	fmt.Println(pkg4.Location())
}
