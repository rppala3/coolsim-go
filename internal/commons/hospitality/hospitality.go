package hospitality

type Hospitality struct {
	guests map[string]Guest
}

func NewHospitality() *Hospitality {
	return &Hospitality{
		make(map[string]Guest),
	}
}

func (host *Hospitality) Come(guest Guest) {
	host.guests[guest.GetUID()] = guest
}

func (host *Hospitality) Leave(guest Guest) {
	delete(host.guests, guest.GetUID())
}
