package types

type Action interface {
	Description() string
	Do(subject Agent)
}
