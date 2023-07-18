package main

import (
	"fmt"
	"sync"
	"time"
)

func or(channels ...<-chan interface{}) <-chan interface{} {
	wg := &sync.WaitGroup{}
	wg.Add(len(channels))
	
	// create single channel
	single := make(chan interface{})							
	

	
	for _, channel := range channels {
		
		// create goroutine for each channel 
		go func(channel <-chan interface{}) {
			defer wg.Done()
			<-channel
		}(channel)
	}

	// close single channel when all channels are closed
	go func(){
		wg.Wait()
		close(single)
	}()

	return single
}

func main() {
	sig := func(after time.Duration) <- chan interface{} {
		c := make(chan interface{})
		go func() {
			defer close(c)
			time.Sleep(after)
		}()
		return c
		}
	
	start := time.Now()
	<-or(
		sig(2*time.Second),
		sig(5*time.Second),
		sig(1*time.Second),
		sig(15*time.Second),
		sig(10*time.Second),
	)
	
	fmt.Printf("done after %v", time.Since(start))
	
}