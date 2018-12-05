package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Record struct {
	Ts      string
	Minutes int
	Info    string
}

type Records []Record

type Guard struct {
	Id         int
	SleepCount int
	Sleeps     [60]int
}

type Guards map[int]*Guard

func main() {
	file, err := os.Open("../input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	records := make(Records, 0)

	fs := bufio.NewScanner(file)
	for fs.Scan() {
		records = append(records, NewRecord(fs.Text()))
	}

	sort.Sort(records)

	if err := ioutil.WriteFile("../input_sorted.txt", []byte(records.String()), 0644); err != nil {
		log.Fatalf(err.Error())
	}

	guards := calculateGuardsSleepingTimes(records)

	// Find the guard that has the most minutes asleep. What minute does that guard spend asleep the most?
	maxSleepingGuard := guards.Max()
	fmt.Printf("the guard that has the most minutes asleep: %+v\n", maxSleepingGuard)
	maxMinute, sleepCount := maxSleepingGuard.Max()
	fmt.Printf("that guard spend asleep the most at minute %v with sleeping count: %v\n", maxMinute, sleepCount)

	// What is the ID of the guard you chose multiplied by the minute you chose?
	fmt.Printf("ID (#%v) of that guard multiplied by the minute: %v\n", maxSleepingGuard.Id, maxSleepingGuard.Id*maxMinute)
}

func NewRecord(line string) Record {

	line = strings.TrimSpace(line)
	if line == "" {
		log.Fatalf("empty line detected!")
	}

	// [1518-11-05 00:03] Guard #99 begins shift
	splitted := strings.Split(line, "]")

	ts := splitted[0][1:]
	tsm := ts[len(ts)-2:]
	repl := strings.NewReplacer("-", "", " ", "", ":", "")
	ts = repl.Replace(ts)

	minutes, err := strconv.Atoi(tsm)
	if err != nil {
		log.Fatalf("could not convert %v to int: %v\n", tsm, err)
	}

	return Record{
		Ts:      ts,
		Minutes: minutes,
		Info:    strings.TrimSpace(splitted[1]),
	}
}

func (r Record) String() string {
	return fmt.Sprintf("[%v-%v-%v %v:%02d] %v", r.Ts[:4], r.Ts[4:6], r.Ts[6:8], r.Ts[8:10], r.Minutes, r.Info)
}

func (rs Records) Len() int {
	return len(rs)
}

func (rs Records) Less(i, j int) bool {
	return rs[i].Ts < rs[j].Ts
}

func (rs Records) Swap(i, j int) {
	rs[i], rs[j] = rs[j], rs[i]
}

func (rs Records) String() string {
	sb := strings.Builder{}
	for i := range rs {
		sb.WriteString(rs[i].String())
		sb.WriteString("\n")
	}
	return sb.String()
}

func calculateGuardsSleepingTimes(records Records) Guards {

	guards := make(Guards)

	var id int
	var fallAsleepMinute int
	var err error

	for _, r := range records {
		if strings.HasPrefix(r.Info, "Guard") {
			splitted := strings.Split(r.Info, " ")
			id, err = strconv.Atoi(splitted[1][1:])
			if err != nil {
				log.Fatalf("could not convert %v to int: %v\n", splitted[1], err)
			}
			if _, ok := guards[id]; !ok {
				guards[id] = &Guard{Id: id, Sleeps: [60]int{}}
			}
			continue
		}

		if r.Info == "falls asleep" {
			fallAsleepMinute = r.Minutes
			continue
		}

		if r.Info == "wakes up" {
			guards[id].SleepCount += r.Minutes - fallAsleepMinute
			for i := fallAsleepMinute; i < r.Minutes; i++ {
				guards[id].Sleeps[i]++
			}
			fallAsleepMinute = 0
		}
	}

	return guards
}

func (gs Guards) Max() *Guard {
	maxGuard := &Guard{}
	for _, g := range gs {
		if g.SleepCount > maxGuard.SleepCount {
			maxGuard = g
		}
	}
	return maxGuard
}

func (g Guard) Max() (minute int, sleepCount int) {
	for i, c := range g.Sleeps {
		if c > sleepCount {
			sleepCount = c
			minute = i
		}
	}
	return minute, sleepCount
}
