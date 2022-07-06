package utils

import "time"

// MakeInf returns (in, out) channels for write and read, simulating an infinite buffer
func MakeInf[T any]() (chan<- T, <-chan T) {
	in := make(chan T, 1)
	out := make(chan T, 1)

	go func() {
		var inQueue []T
		outCh := func() chan T {
			if len(inQueue) > 0 {
				return out
			}
			return nil // nil channel not allow to write on it
		}

		curVal := func() T {
			if len(inQueue) > 0 {
				return inQueue[0]
			}
			return getZero[T]()
		}

		for len(inQueue) > 0 || in != nil {
			select {
			case val, ok := <-in:
				if !ok {
					in = nil
				} else {
					inQueue = append(inQueue, val)
				}
			case outCh() <- curVal():
				inQueue = inQueue[1:]
			}
		}
		close(out)
	}()
	return in, out
}

func getZero[T any]() T {
	var result T
	return result
}

// Consume all the elements of the channel and returns after 1 second
func Consume[T any](ch <-chan T) {
	for {
		select {
		case <-ch:
		case <-time.After(time.Second):
			return
		}
	}

}
