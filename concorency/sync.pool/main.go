package main

import (
	"fmt"
	"sync"
)

type Pooling struct {
	Host string
	Port int
	Ip   string
}

var (
	Pool = sync.Pool{
		New: func() any {
			return &Pooling{Host: "localhost", Port: 22, Ip: "192.168.1.1"}
		},
	}
)

func main() {
	// Get a new instance from the pool
	res := Pool.Get()
	fmt.Println(res)

	// Put the instance back to the pool
	Pool.Put(res)

	// Get another instance from the pool
	newRes := Pool.Get()
	fmt.Println(newRes)

	// Put the instance back to the pool
	Pool.Put(newRes)
}
