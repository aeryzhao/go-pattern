package main

type mediator interface {
	canArrive(train) bool
	notifyAboutDeparture()
}

type train interface {
	arrive()
	depart()
	permitArrival()
}

type passengerTrain struct {
	mediator mediator
}

func (p *passengerTrain) arrive() {
	if !p.mediator.canArrive(p) {
		println("PassengerTrain: Arrival blocked, waiting")
		return
	}
	println("PassengerTrain: Arrived")
}

func (p *passengerTrain) depart() {
	println("PassengerTrain: Leaving")
	p.mediator.notifyAboutDeparture()
}

func (p *passengerTrain) permitArrival() {
	println("PassengerTrain: Arrival permitted")
	p.arrive()
}

type freightTrain struct {
	mediator mediator
}

func (f *freightTrain) arrive() {
	if !f.mediator.canArrive(f) {
		println("FreightTrain: Arrival blocked, waiting")
		return
	}
	println("FreightTrain: Arrived")
}

func (f *freightTrain) depart() {
	println("FreightTrain: Leaving")
	f.mediator.notifyAboutDeparture()
}

func (f *freightTrain) permitArrival() {
	println("FreightTrain: Arrival permitted")
	f.arrive()
}

type stationManager struct {
	isPlatformFree bool
	trainQueue     []train
}

func newStationManger() *stationManager {
	return &stationManager{
		isPlatformFree: true,
	}
}

func (s *stationManager) canArrive(t train) bool {
	if s.isPlatformFree {
		s.isPlatformFree = false
		return true
	}
	s.trainQueue = append(s.trainQueue, t)
	return false
}

func (s *stationManager) notifyAboutDeparture() {
	if !s.isPlatformFree {
		s.isPlatformFree = true
	}
	if len(s.trainQueue) > 0 {
		firstTrainInQueue := s.trainQueue[0]
		s.trainQueue = s.trainQueue[1:]
		firstTrainInQueue.permitArrival()
	}
}

func main() {
	stationManager := newStationManger()
	passengerTrain := &passengerTrain{
		mediator: stationManager,
	}
	freightTrain := &freightTrain{
		mediator: stationManager,
	}
	passengerTrain.arrive()
	freightTrain.arrive()
	passengerTrain.depart()
}
