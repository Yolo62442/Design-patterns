package main

import "fmt"

type Wizard struct {
	Name string
	Age int
	School string
	Score int
}

func (w *Wizard) AddScore (score int)  	{
	w.Score += score
}
func (w *Wizard) String() string  {
	return fmt.Sprintf(w.Name + ", %d years old from " + w.School , w.Age )
}
