package engine

import "fmt"

type controller struct {

}

func (c *controller) run() {
	fmt.Println("controller run")
}