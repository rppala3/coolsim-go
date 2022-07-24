package hospitality

type Hospitable interface {
	Come(guest Guest)
	Leave(guest Guest)
}
