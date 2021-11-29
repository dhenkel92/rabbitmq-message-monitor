package helper

import "time"

type IntervalAction func()

func StartInterval(interval time.Duration, fn IntervalAction) chan bool {
	ticker := time.NewTicker(interval)
	quit := make(chan bool)

	go func() {
		for {
			select {
			case <-ticker.C:
				fn()
			case <-quit:
				ticker.Stop()
				return
			}
		}
	}()

	return quit
}
