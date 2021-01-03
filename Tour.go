package main

import "fmt"

type Tour struct {
	Location string
	Winner *Wizard
	Count int
	TourNumber int
}

func (t *Tour) Start() {
	fmt.Printf("The %d tour started and it is located " + t.Location, t.TourNumber)
	fmt.Println()
}
func (t *Tour) End() {
	fmt.Printf("The %d tour finished and the winner is " + t.Winner.String(), t.TourNumber)
	fmt.Println()
}

func (t *Tour) AddWinner (wizard *Wizard) {
	wizard.AddScore(t.Count)
	t.Winner = wizard
}
func (t *Tour) String() string {
	return fmt.Sprintf("The Tour winner is ", t.Winner.String)
}

