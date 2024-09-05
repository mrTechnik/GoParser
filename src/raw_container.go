package src

import (
	"errors"
)

type NewLinksContainer struct {
	currentRamBuffer int
	MaxRamBuffer     int
	MinimalRamBuffer int
	firstElem        *RAMQueue
	lastElem         *RAMQueue
}

func (newLinks *NewLinksContainer) getFirstElem() (any, error) {
	if newLinks.firstElem == nil {
		return nil, errors.New("\nNo queue is empty nothing to return")
	}

	return newLinks.firstElem.element, nil
}

func (newLinks *NewLinksContainer) AddNewElem(elem any) error {
	bufElem := &RAMQueue{elem, nil}

	if newLinks.firstElem == nil {
		newLinks.firstElem = bufElem
	} else {
		newLinks.lastElem.nextElem = bufElem
	}
	newLinks.lastElem = bufElem

	newLinks.currentRamBuffer++
	return nil
}

func (newLinks *NewLinksContainer) DeleteFirstElem() error {
	if newLinks.lastElem == nil {
		return errors.New("\nno last element")
	}
	if newLinks.currentRamBuffer == 0 {
		errors.New("\ncounter doesn't contain elements")
	}
	buf := newLinks.firstElem.nextElem
	newLinks.firstElem = buf

	newLinks.currentRamBuffer--
	return nil
}

func (newLinks *NewLinksContainer) GetAndDeleteFirstElem() (any, error) {
	elem, err := newLinks.getFirstElem()
	err_ := newLinks.DeleteFirstElem()

	if err != nil {
		return nil, err
	}
	if err_ != nil {
		return nil, err_
	}

	return elem, nil
}

func (newLinks *NewLinksContainer) CheckCountOfLinks() (bool, error) {
	if newLinks.currentRamBuffer < 0 {
		return false, errors.New("\nbuffer lower than zero")
	}
	if newLinks.currentRamBuffer >= newLinks.MaxRamBuffer {
		return true, nil
	}

	return false, nil
}

func (NewLinks *NewLinksContainer) SaveLinksInFile() (int, error) {
	var countOfLines int = NewLinks.currentRamBuffer - NewLinks.MinimalRamBuffer
	if countOfLines < 0 {
		return 0, errors.New("\nminimal count of links for saving more than container contain now")
	}

	i := countOfLines
	for i <= 0 {
		bufElem, err := NewLinks.GetAndDeleteFirstElem()
		if err != nil {
			return 0, err
		}
		NewLinks.SaveLinkInFile(bufElem)
	}

	return countOfLines, nil
}

func (newNewLinks *NewLinksContainer) SaveLinkInFile(link any) error {
	return nil
}
