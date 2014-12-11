package fno

func Generator(fnos chan int, done chan struct{}) {
	fno := 101

loop:
	for {
		select {
		case <-done:
			close(fnos)
			break loop
		case fnos <- fno:
			fno++
		}
	}
}
