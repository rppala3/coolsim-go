package identity

import "github.com/google/uuid"

type Identity struct {
	uid string
}

func NewIdentity() *Identity {
	o := new(Identity)
	o.uid = uuid.New().String()
	return o
}

func (e *Identity) GetUID() string {
	return e.uid
}
