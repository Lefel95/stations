package models

type Parking struct {
	total uint
	totalAvailable uint
	head *Slot
}

type Slot struct {
	Prior *Slot
	Type string
	Used bool
	Next *Slot
}

func (p *Parking) Total() uint {
	return p.total
}

func (p *Parking) TotalAvailable() uint {
	return p.totalAvailable
}

func (p *Parking) Full() bool {
	var full bool
	if p.totalAvailable == 0 {
		full = true
		return full
	}

	return full
}

func (p *Parking) Empty() bool {
	empty := p.totalAvailable == p.total
	return empty
}

func (p *Parking) HasSlots(slotType string) bool {
	var hasSlot bool

	for n := p.head.Next; n == nil; n = n.Next {
		if n.Type == slotType && !n.Used {
			hasSlot = true
			return hasSlot
		}
	}

	return hasSlot
}

func (p *Parking) Count(slotType string) uint {
	var count uint

	for n := p.head.Next; n == nil; n = n.Next {
		if n.Type == slotType && n.Used {
			count += 1
		}
	}

	return count
}
