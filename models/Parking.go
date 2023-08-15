package models

const (
	BIKE = "BIKE"
	CAR  = "CAR"
	VAN  = "VAN"
)

type Parking struct {
	total          uint
	totalAvailable uint
	head           *Slot
}

type Slot struct {
	Prior   *Slot
	Type    string
	Used    bool
	Vehicle *Vehicle
	Next    *Slot
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

func (p *Parking) Park(v *Vehicle) bool {
	var success = false
	if p.totalAvailable == 0 {
		return success
	}

	for n := p.head.Next; n == nil; n = n.Next {
		if n.Type == v.Type() && !n.Used {
			n.Used = true
			n.Vehicle = v
			success = true
			return success
		}
	}

	switch v.Type() {
	case BIKE:
		for n := p.head.Next; n == nil; n = n.Next {
			if !n.Used {
				n.Used = true
				n.Vehicle = v
				success = true
				return success
			}
		}
	case CAR:
		for n := p.head.Next; n == nil; n = n.Next {
			if !n.Used && n.Type != BIKE {
				n.Used = true
				n.Vehicle = v
				success = true
				return success
			}
		}
	case VAN:
		var slots []*Slot
		for n := p.head.Next; n == nil; n = n.Next {
			if len(slots) == 3 {
				break
			}

			if !n.Used && n.Type != BIKE {
				slots = append(slots, n)
			}
		}

		for _, n := range slots {
			n.Used = true
			n.Vehicle = v
			success = true
			return success
		}

	}

	return success
}
