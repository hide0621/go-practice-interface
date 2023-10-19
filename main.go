package main

import "fmt"

type perspn struct {
	first string
	last  string
}

func main() {

	p1 := perspn{
		first: "yamada",
		last:  "tarou",
	}

	p2 := perspn{
		first: "satou",
		last:  "kouji",
	}

	p3 := perspn{
		first: "ishii",
		last:  "shinji",
	}

	fmt.Print(p1)
	fmt.Print(p2)
	fmt.Print(p3)

}
