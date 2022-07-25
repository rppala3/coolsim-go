package types

type Action interface {
	Name() string
	Description() string
	Do(subject Agent)
}
