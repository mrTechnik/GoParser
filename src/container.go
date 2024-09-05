package src

import (
	"errors"
)

type FreshLinksContainer struct {
	currentRamBuffer int
	MaxRamBuffer     int
	MinimalRamBuffer int
	firstElem        *RAMQueue
	lastElem         *RAMQueue
}

func (ramLinks *FreshLinksContainer) CheckCountOfElems() (int, error) {
	if ramLinks.currentRamBuffer <= 0 {
		return -2, errors.New("\ncurrentRamBuffer equal or less then zero")
	}

	if ramLinks.currentRamBuffer < ramLinks.MinimalRamBuffer {
		return -1, nil
	} else if ramLinks.currentRamBuffer > ramLinks.MaxRamBuffer {
		return 1, nil
	}
	return 0, nil
}

func (ramLinks *FreshLinksContainer) GetFirstElem() (any, error) {
	if ramLinks.currentRamBuffer <= 0 {
		return nil, errors.New("\nno elements in Container")
	}
	return ramLinks.firstElem.element, nil
}

func (ramLinks *FreshLinksContainer) GetFirstElemAndDelete() (any, error) {
	if ramLinks.firstElem == nil {
		return nil, errors.New("\nno elements in Container")
	}

	bufElem := ramLinks.firstElem.element

	if ramLinks.currentRamBuffer > 0 {
		ramLinks.firstElem = ramLinks.firstElem.nextElem
	}

	ramLinks.currentRamBuffer--

	return bufElem, nil
}

func (ramLinks *FreshLinksContainer) GetLastElem() (any, error) {
	if ramLinks.lastElem == nil {
		return nil, errors.New("\nno elements in Container")
	}
	return ramLinks.lastElem.element, nil
}

func (ramLinks *FreshLinksContainer) AddNewElem(elem any) error {
	link := &RAMQueue{elem, nil}

	if ramLinks.currentRamBuffer == 0 {
		ramLinks.firstElem = link
	} else {
		ramLinks.lastElem.nextElem = link
	}
	ramLinks.lastElem = link

	ramLinks.currentRamBuffer++
	return nil
}

func (ramLinks *FreshLinksContainer) DeleteFirstElem() error {
	if ramLinks.currentRamBuffer < 1 {
		return errors.New("\nqueue doesn't have any elements")
	}

	if ramLinks.currentRamBuffer == 1 {
		ramLinks.firstElem = nil
		ramLinks.lastElem = nil
	} else {
		bufLink := &ramLinks.firstElem.nextElem
		ramLinks.firstElem = *bufLink
	}
	return nil
}
