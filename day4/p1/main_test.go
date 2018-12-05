package main

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

const exampleRecords = `[1518-11-01 00:00] Guard #10 begins shift
[1518-11-01 00:05] falls asleep
[1518-11-01 00:25] wakes up
[1518-11-01 00:30] falls asleep
[1518-11-01 00:55] wakes up
[1518-11-01 23:58] Guard #99 begins shift
[1518-11-02 00:40] falls asleep
[1518-11-02 00:50] wakes up
[1518-11-03 00:05] Guard #10 begins shift
[1518-11-03 00:24] falls asleep
[1518-11-03 00:29] wakes up
[1518-11-04 00:02] Guard #99 begins shift
[1518-11-04 00:36] falls asleep
[1518-11-04 00:46] wakes up
[1518-11-05 00:03] Guard #99 begins shift
[1518-11-05 00:45] falls asleep
[1518-11-05 00:55] wakes up`

func Test_calculateGuardsSleepingTimes(t *testing.T) {

	lines := strings.Split(exampleRecords, "\n")
	records := make(Records, 0)
	for _, line := range lines {
		records = append(records, NewRecord(line))
	}

	tests := []struct {
		name     string
		expected Guard
	}{
		{
			"example guard 10: 50 minutes overall, 24th minute most",
			Guard{
				Id:         10,
				SleepCount: 50,
				Sleeps:     [60]int{0, 0, 0, 0, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 2, 1, 1, 1, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
			},
		},
		{
			"example guard 99: ",
			Guard{
				Id:         99,
				SleepCount: 30,
				Sleeps:     [60]int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 1, 2, 2, 2, 2, 2, 3, 2, 2, 2, 2, 1, 1, 1, 1, 1},
			},
		},
	}

	guards := calculateGuardsSleepingTimes(records)
	assert.Equal(t, 2, len(guards))

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			actual, ok := guards[tt.expected.Id]
			assert.True(t, ok)

			assert.Equal(t, tt.expected.SleepCount, actual.SleepCount)
			assert.Equal(t, tt.expected.Sleeps, actual.Sleeps)
		})
	}
}
