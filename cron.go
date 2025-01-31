package cron

import "time"

const (
	NbSecondsInDay = 86400
)

type Scheduler struct {
	Next  time.Time
	Every time.Duration
	Slip  time.Duration
	Loc   *time.Location
}

func NewCronSchedulerUTC(every, slip time.Duration) *Scheduler {
	return NewCronScheduler(every, slip, time.UTC)
}

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

func (thiz *Scheduler) IsPassed() bool {
	if thiz.Compare(time.Now()) {
		thiz.ComputeNext()
		return true
	}
	return false
}

func (thiz *Scheduler) Compare(t time.Time) bool {
	return t.In(thiz.Loc).UnixNano() >= thiz.Next.UnixNano()
}
