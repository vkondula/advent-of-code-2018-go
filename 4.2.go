package main

import (
	"./lib"
	"time"
	"regexp"
	"sort"
	"strings"
	"strconv"
	"fmt"
)

const HOUR = 60

type Record struct {
	start time.Time
	text string
}


type Office struct {
	sleepingSchedule map[int][HOUR]int
	lastRecord time.Time
	currentOfficer int
}

func (office *Office) FallAsleep(start time.Time) { office.lastRecord = start }
func (office *Office) Assign(officer int)         { office.currentOfficer = officer }
func (office *Office) WakeUp(end time.Time) {
	value, ok := office.sleepingSchedule[office.currentOfficer]
	if !ok {
		value = [HOUR]int{}
	}
	startMinute := office.lastRecord.Minute()
	endMinute := end.Minute()
	for i := startMinute ; i < endMinute ; i++ {
		value[i]++
	}
	office.sleepingSchedule[office.currentOfficer] = value
}


func maxIndexMap(records map[int][HOUR]int) (maxKey int, maxIndex int) {
	max := 0
	for key, record := range records {
		for index, value := range record {
			if value > max {
				max = value
				maxKey = key
				maxIndex = index
			}
		}
	}
	return
}


func main() {
	lines := lib.GetItems("inputs/4.txt")
	const template = "2006-01-02 15:04"
	re, _ := regexp.Compile(`\[(.*)] (.*)`)
	var records []Record
	for _, line := range lines {
		matched := re.FindStringSubmatch(line)
		start, _ := time.Parse(template, matched[1])
		records = append(records, Record{start, matched[2]})
	}

	sort.Slice(records, func (i, j int) bool { return records[i].start.Before(records[j].start) })

	office := Office{make(map[int][HOUR]int), time.Now(), 0}
	reOfficer, _ := regexp.Compile(`#(\d+)`)
	for _, record := range records {
		if strings.Contains(record.text, "wakes up") {
			office.WakeUp(record.start)
		} else if strings.Contains(record.text, "falls asleep") {
			office.FallAsleep(record.start)
		} else {
			matched := reOfficer.FindStringSubmatch(record.text)
			value, _ := strconv.Atoi(matched[1])
			office.Assign(value)
		}
	}
	officer, sleeping := maxIndexMap(office.sleepingSchedule)
	fmt.Printf("%d\n", officer * sleeping)
}