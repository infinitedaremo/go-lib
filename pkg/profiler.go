package profiler

import (
	"time"
)

// Profiler : Interface for a profiler
type Profiler interface {
	Start()
	End() time.Duration
}

// TimeProfiler : Allows profiling between two points in code
type TimeProfiler struct {
	StartTime time.Time
}

// NewTimeProfiler : Returns a new time profiler
func NewTimeProfiler() *TimeProfiler {
	return &TimeProfiler{}
}

// Start : Caches the start time of the profiler
func (p *TimeProfiler) Start() {
	p.StartTime = time.Now()
}

// End : Calculates the difference between the start cache and now
func (p *TimeProfiler) End() time.Duration {
	return time.Since(p.StartTime)
}
