package service

import (
	"fmt"
	log "goapp/log_config"
	"sync"
	"time"
)

func doWork(wg *sync.WaitGroup, id int) {
	defer wg.Done()
	log.Info(fmt.Sprintf("Goroutine %d is working...", id))
	time.Sleep(1 * time.Second)
	log.Info(fmt.Sprintf("Goroutine %d is done!", id))
}

func PocMultiThread(threadAmount int) {
	var wg sync.WaitGroup
	defer wg.Wait()

	// Setting running thread number to 20
	log.Info("Number of goroutines: " + fmt.Sprint(threadAmount))
	wg.Add(threadAmount)
	for i := 1; i <= threadAmount; i++ {
		go doWork(&wg, i)
	}
}
