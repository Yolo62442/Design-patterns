package main

import "fmt"

type TriwizardTournamentFacade struct {
	Data *Data
	Participants *Participants
	Tours *Tours
	Tournum int
}

func (f *TriwizardTournamentFacade) FindWinner() {
	max := 0
	index := 0
	for i, w := range f.Participants.Participants{
		if w.Score > max{
			max = w.Score
			index = i
		}
		fmt.Println(w.String() + "with scores", w.Score)
	}
	fmt.Println()
	f.Data.Winner = f.Participants.Participants[index]
	f.Data.String()
}
func (f *TriwizardTournamentFacade) AddParticipant(name string, school string, age int) *Participants {
	f.Participants.AddParticipant(&Wizard{Name: name, School: school, Age: age, Score: 0})
	return f.Participants
}
func (f *TriwizardTournamentFacade) AddTour(l string, count int, wizard *Wizard)  {
	f.Tours.AddTour(l, count, f.Tournum)
	f.Tours.Tours[f.Tournum-1].AddWinner(wizard)
	f.Tournum += 1
}
func (f TriwizardTournamentFacade) Start() {
	fmt.Println("Welcome to the TriwizardTournament\n Our Participants are: ")
	for _,p := range f.Participants.Participants{
		fmt.Println(p.String())
	}
	fmt.Println()
	fmt.Println("Let's begin our tournament")
	for _,t := range f.Tours.Tours{
		t.Start()
		t.End()
		fmt.Println()
	}
	f.FindWinner()
}
func NewTriwizardFacade(s string, y int) *TriwizardTournamentFacade {
	triwizard := &TriwizardTournamentFacade{
		Data: &Data{School: s, Year: y},
		Participants: &Participants{},
		Tours: &Tours{},
		Tournum: 1,
	}
	return triwizard
}

func main()  {
	 tournament := NewTriwizardFacade("Hogwarts", 1994)
	 tournament.AddParticipant("Viktor Krum", "Durmstrang", 17)
	 tournament.AddParticipant("Fleur Delacour", "Beauxbatons", 17)
	 tournament.AddParticipant("Cedric Diggory", "Hogwarts", 17)
	 p := tournament.AddParticipant("Harry Potter", "Hogwarts", 14)
	 tournament.AddTour("Quidditch arena", 2, p.Participants[0])
	 tournament.AddTour("Black lake", 4, p.Participants[3])
	 tournament.AddTour("Maze", 5, p.Participants[1])
	 fmt.Println(p.Participants[1].Score)
	 tournament.Start()



}