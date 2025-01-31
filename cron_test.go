package cron_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/alex-cos/cron"
	"github.com/stretchr/testify/assert"
)

func Test1(t *testing.T) {
	t.Parallel()

	every := 86400 * time.Second
	slip := time.Duration(0)
	cron := cron.NewCronSchedulerUTC(every, slip)

	v := time.Now().UTC().AddDate(0, 0, 1)
	expected := time.Date(v.Year(), v.Month(), v.Day(), 0, 0, 0, 0, time.UTC)
	fmt.Printf("expected=%v, got=%v\n", expected, cron.Next)
	assert.Equal(t, expected, cron.Next)
	assert.False(t, cron.IsPassed(), "Passed comparison failed")
}

func Test2(t *testing.T) {
	t.Parallel()

	every := 86400 * time.Second
	slip := 5 * time.Minute
	cron := cron.NewCronSchedulerUTC(every, slip)

	v := time.Now().UTC().AddDate(0, 0, 1)
	expected := time.Date(v.Year(), v.Month(), v.Day(), 0, 5, 0, 0, time.UTC)
	fmt.Printf("expected=%v, got=%v\n", expected, cron.Next)
	assert.Equal(t, expected, cron.Next)
	assert.False(t, cron.IsPassed(), "Passed comparison failed")
}

func Test3(t *testing.T) {
	t.Parallel()

	every := 6 * time.Hour
	slip := time.Duration(0)
	cron := cron.NewCronSchedulerUTC(every, slip)

	v := time.Now().UTC()
	expected := v.Truncate(6 * time.Hour).Add(6 * time.Hour)
	fmt.Printf("expected=%v, got=%v\n", expected, cron.Next)
	assert.Equal(t, expected, cron.Next)
	assert.False(t, cron.IsPassed(), "Passed comparison failed")
}
