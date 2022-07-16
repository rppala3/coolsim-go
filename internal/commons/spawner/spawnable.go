package spawner

import (
	"sync"

	b_ "github.com/rppala3/broadcast-channel"
)

type AckChannel chan bool

type Spawnable interface {
	// type Spawnable[T any] interface {
	// Spawn(chan T)

	Spawn(*b_.Listener, AckChannel, *sync.WaitGroup)
}
