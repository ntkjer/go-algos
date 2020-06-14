package stopwatch

import "time"

type Stopwatch struct {
	time    time.Time
	elapsed time.Duration
}

//Starts the stopwatch timer as the current time
func (s *Stopwatch) Start() {
	s.time = time.Now()
}

//Stops the stopwatch struct and records elapsed time
func (s *Stopwatch) Stop() {
	t := time.Now()
	s.elapsed = t.Sub(s.time)
}

//Returns the elapsed time for stopwatch instance to client
func (s *Stopwatch) ElapsedTime() time.Duration {
	return s.elapsed
}

func (s *Stopwatch) ElapsedTimeSeconds() float64 {
	return s.elapsed.Seconds()
}
