package main

type Participants struct {
	Participants[] *Wizard
}

func (p *Participants) AddParticipant (wizard *Wizard) {
	p.Participants = append(p.Participants, wizard)
}

