package persist

import "fmt"

func PropertySaver() chan interface{} {
	out := make(chan interface{})
	go func() {
		for {
			property := <- out
			fmt.Println("get property: ", property)
		}
	}()
	return out
}
