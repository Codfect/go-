package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {

	contador := 0
	totaldegoroutines := 100

	var wg sync.WaitGroup
	wg.Add(totaldegoroutines)
	
	//MUTEX
	var mtx sync.Mutex

	for i := 0; i < totaldegoroutines; i++ {
		go func() {
			mtx.Lock()
			
			v := contador
			runtime.Gosched()
			v++
			contador = v
			
			fmt.Println(contador)
			wg.Done()
			
			mtx.Unlock()
		}()
	}
	wg.Wait()

	fmt.Println("Total", contador)
}
