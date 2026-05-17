package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"log"
)

// Shape is the interface we want to include in our struct
type Shape interface {
	Area() float32
}

// Circle is a concrete implementation of Shape
type Circle struct {
	Radius float32
}

func (c Circle) Area() float32 {
	return 3.14 * c.Radius * c.Radius
}

type Config struct {
	Value []Shape // Interface field

}

// Container holds the interface field
type Container struct {
	ID     int
	Config Config
}

func main() {
	// CRITICAL: Register the concrete type that implements the interface
	gob.Register(Circle{})

	var network bytes.Buffer // Stand-in for a network connection

	// 1. Encoding
	enc := gob.NewEncoder(&network)
	c := Container{
		ID: 1,
		Config: Config{
			Value: []Shape{
				Circle{Radius: 10},
			},
		},
	}

	err := enc.Encode(c)
	if err != nil {
		log.Fatal("encode error:", err)
	}

	// 2. Decoding
	dec := gob.NewDecoder(&network)
	var result Container
	err = dec.Decode(&result)
	if err != nil {
		log.Fatal("decode error:", err)
	}

	fmt.Printf("ID: %d, Area: %v\n", result.ID, result.Config.Value[0].Area())
}
