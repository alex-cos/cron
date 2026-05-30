// Package cron provides a simple time-based scheduler for computing
// the next execution time of recurring tasks.
package cron

import "time"

const (
	// NbSecondsInDay is the number of seconds in a full day (24 hours).
	NbSecondsInDay = 86400
)

// Scheduler computes the next execution time for a recurring task
// based on a fixed interval and an optional delay.
type Scheduler struct {
	// Next is the next scheduled execution time.
	Next time.Time

	// Every is the interval between consecutive executions.
	Every time.Duration

	// Slip is an optional delay added to each scheduled time.
	Slip time.Duration

	// Loc is the timezone used for scheduling.
	Loc *time.Location
}

// NewCronSchedulerUTC creates a new Scheduler using UTC timezone.
// The every parameter defines the interval between executions.
// The slip parameter defines an optional delay added to each scheduled time.
func NewCronSchedulerUTC(every, slip time.Duration) *Scheduler {
	return NewCronScheduler(every, slip, time.UTC)
}

// NewCronScheduler creates a new Scheduler using the given timezone.
// The every parameter defines the interval between executions.
// The slip parameter defines an optional delay added to each scheduled time.
// The loc parameter specifies the timezone to use for scheduling.
func NewCronScheduler(every, slip time.Duration, loc *time.Location) *Scheduler {
	thiz := &Scheduler{
		Next:  time.Now().In(loc),
		Every: every,
		Slip:  slip,
		Loc:   loc,
	}
	thiz.ComputeNext()
	return thiz
}

// ComputeNext recalculates the next execution time based on the current time,
// the Every interval, and the Slip delay.
// When Every >= 24 hours, the next time is set to midnight of the following day.
// Otherwise, the next time is truncated to the Every interval boundary and advanced.
func (thiz *Scheduler) ComputeNext() {
	now := time.Now().In(thiz.Loc)
	thiz.Next = now
	if thiz.Every.Seconds() >= NbSecondsInDay {
		thiz.Next = time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, thiz.Loc)
	} else {
		thiz.Next = now.Truncate(thiz.Every)
	}
	thiz.Next = thiz.Next.Add(thiz.Every)
	thiz.Next = thiz.Next.Add(thiz.Slip)
}

// IsPassed reports whether the scheduled time has been reached.
// If so, it advances the scheduler to the next execution time and returns true.
// If the scheduled time has not been reached, it returns false without modification.
func (thiz *Scheduler) IsPassed() bool {
	if thiz.Compare(time.Now()) {
		thiz.ComputeNext()
		return true
	}
	return false
}

// Compare reports whether the given time t is at or after the scheduled Next time.
// The comparison is performed in the scheduler's timezone.
func (thiz *Scheduler) Compare(t time.Time) bool {
	return t.In(thiz.Loc).UnixNano() >= thiz.Next.UnixNano()
}
