package utils

import "time"

// RepeatAction calls the given action in intervals of delay. The function finishes when some info is sent to the channel.
func RepeatAction(action func(), delay time.Duration, kill <-chan any) {
	ticker := time.NewTicker(delay)
	for {
		select {
		case <-ticker.C:
			action()
		case <-kill:
			break
		}
	}
}
