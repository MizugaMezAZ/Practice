package main

import "fmt"

func twoGoFuncPrint() {
	ch := make(chan bool)
	done := make(chan struct{})

	// wg := sync.WaitGroup{}

	// wg.Add(1)

	go func() {
		n := 'a'
	BR:
		for {
			<-ch
			fmt.Print(string(n))
			fmt.Print(string(n + 1))
			n += 2

			if n > 'z' {
				ch <- true
				break BR
			}

			ch <- false
		}

	}()

	ch <- false

	go func() {
		n := 1
	BR:
		for {
			f := <-ch
			fmt.Print(n)
			fmt.Print(n + 1)
			n += 2

			if f {
				done <- struct{}{}
				break BR
			}

			ch <- false

		}
	}()

	<-done
}
