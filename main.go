package main

import "fmt"

type perspn struct {
	first string
	last  string
}

func main() {

	p := []perspn{
		{
			first: "yamada",
			last:  "tarou",
		},
		{
			first: "satou",
			last:  "kouji",
		},
		{
			first: "ishii",
			last:  "shinji",
		},
	}

	for _, v := range p {
		fmt.Println(v)
	}
}
