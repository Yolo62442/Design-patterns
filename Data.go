package main

import "fmt"

type Data struct {
	School string
	Year int
	Winner *Wizard
}

func (d *Data) String()  {
	fmt.Printf("In %d year the tournament was held at " + d.School + " and the winner is ", d.Year)
	fmt.Print(d.Winner.String())
}