# cron

> **⚠️ This repository is no longer maintained. Please use [github.com/alex-cos/scheduler](https://github.com/alex-cos/scheduler) instead.**

[![Go Version](https://img.shields.io/badge/Go-1.23%2B-blue)](https://go.dev/)
[![Test Status](https://github.com/alex-cos/cron/actions/workflows/test.yml/badge.svg)](https://github.com/alex-cos/cron/actions/workflows/test.yml)
[![Codecov](https://codecov.io/gh/alex-cos/cron/branch/main/graph/badge.svg)](https://codecov.io/gh/alex-cos/cron)
[![Lint Status](https://github.com/alex-cos/cron/actions/workflows/lint.yml/badge.svg)](https://github.com/alex-cos/cron/actions/workflows/lint.yml)
[![License](https://img.shields.io/badge/License-MIT-green)](LICENSE)
[![Go Report Card](https://goreportcard.com/badge/github.com/alex-cos/cron)](https://goreportcard.com/report/github.com/alex-cos/cron)

A simple Go library for computing the next execution time of recurring tasks based on a fixed interval.

## Installation

```bash
go get github.com/alex-cos/cron
```

## Usage

```go
package main

import (
  "fmt"
  "time"

  "github.com/alex-cos/cron"
)

func main() {
  // Run every 6 hours in UTC
  scheduler := cron.NewCronSchedulerUTC(6*time.Hour, 0)

  fmt.Println("Next execution:", scheduler.Next)

  // Check if the scheduled time has passed
  if scheduler.IsPassed() {
    fmt.Println("Task is due!")
  }

  // Run every day at midnight UTC with a 5-minute delay
  daily := cron.NewCronSchedulerUTC(86400*time.Second, 5*time.Minute)
  fmt.Println("Next daily execution:", daily.Next)

  // Run every 30 minutes in a specific timezone
  loc, _ := time.LoadLocation("Europe/Paris")
  halfHour := cron.NewCronScheduler(30*time.Minute, 0, loc)
  fmt.Println("Next execution (Paris):", halfHour.Next)
}
```

## API

### Types

#### `Scheduler`

```go
type Scheduler struct {
    Next  time.Time       // Next scheduled execution time
    Every time.Duration   // Interval between executions
    Slip  time.Duration   // Optional delay added to each scheduled time
    Loc   *time.Location  // Timezone used for scheduling
}
```

### Functions

- `NewCronSchedulerUTC(every, slip time.Duration) *Scheduler` — Creates a scheduler in UTC.
- `NewCronScheduler(every, slip time.Duration, loc *time.Location) *Scheduler` — Creates a scheduler with a custom timezone.

### Methods

- `ComputeNext()` — Recalculates the next execution time.
- `IsPassed() bool` — Returns true if the scheduled time has been reached, then advances to the next one.
- `Compare(t time.Time) bool` — Returns true if `t` is at or after the scheduled time.
