package types

type Perception interface {
	Description() string
	GetActions() []Action
}
