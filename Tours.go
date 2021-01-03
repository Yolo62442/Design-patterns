package main

type Tours struct {
	Tours[] *Tour
}

func (t *Tours) AddTour(Location string, count int, tournum int)  {
	t.Tours = append(t.Tours, &Tour{Location: Location, Count: count, TourNumber: tournum})
}
