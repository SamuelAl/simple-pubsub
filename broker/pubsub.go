package pubsub

import (
	"net"
	"sync"
)

type Pubsub struct {
	mu   sync.Mutex
	subs map[string][]chan string
}

type Subscription struct {
	channel  chan string
	address  *net.UDPAddr
	clientID string
}

func NewPubsub() *Pubsub {
	return &Pubsub{
		mu:   sync.Mutex{},
		subs: make(map[string][]chan string),
	}
}
