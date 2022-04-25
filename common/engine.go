package engine

import "fmt"

type Engine interface {
	Run() error
}

type engineDefaultInstance struct {
	host string
}

func (e *engineDefaultInstance) Run() error  {
	fmt.Println("TODO")
	return nil
}

