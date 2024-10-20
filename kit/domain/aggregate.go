package domain

import "golang-template/kit/event"

type BaseAggregate struct {
	events []event.Event `gorm:"-"`
}

func (a *BaseAggregate) PullEvents() []event.Event {
	events := a.events
	a.events = []event.Event{}
	return events
}
