package spidate

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

var currentReading Reading

func Now(reading Reading) (date string) {
	currentReading = reading
	t := time.Now()
	date = convert(t, reading)
	return
}

func GetMonthData(t time.Month, r Reading) DateData {
	currentReading = r
	from := time.Now().AddDate(-1, int(t), -time.Now().Day()+1)
	to := from.AddDate(0, 1, -from.Day())

	month := DateData{
		Name:     t.String(),
		From:     convertHoursToStatic(convert(from, r), Begin),
		To:       convertHoursToStatic(convert(to, r), End),
		FromTime: from,
		ToTime:   to,
	}
	return month
}

func GetToday(r Reading) DateData {
	currentReading = r
	t := time.Now()
	first := convertHoursToStatic(convert(t, r), Begin)
	ends := convertHoursToStatic(convert(t, r), End)
	_, week := t.ISOWeek()
	today := DateData{
		Name:      t.Weekday().String(),
		From:      first,
		To:        ends,
		Day:       t.Day(),
		Time:      t,
		DayOfYear: t.YearDay(),
		Week:      week,
		FromTime:  time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, time.Local),
		ToTime:    time.Date(t.Year(), t.Month(), t.Day(), 23, 59, 59, 0, time.Local),
	}
	return today
}

func AddDaysToToday(days int, r Reading) DateData {
	currentReading = r
	t := time.Now()
	newTime := t.AddDate(0, 0, days)
	first := convertHoursToStatic(convert(t, r), Begin)
	ends := convertHoursToStatic(convert(newTime, r), End)
	title := fmt.Sprintf("Range from %s to %s", first, ends)
	_, week := t.ISOWeek()
	return DateData{
		Name:      title,
		From:      first,
		To:        ends,
		Day:       t.Day(),
		Time:      t,
		DayOfYear: t.YearDay(),
		Week:      week,
		FromTime:  time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, time.Local),
		ToTime:    time.Date(newTime.Year(), newTime.Month(), newTime.Day(), 23, 59, 59, 0, time.Local),
	}
}

func AddDaysToDate(date string, days int, r Reading) DateData {
	currentReading = r
	items := strings.Split(date, "-")
	if len(items) != 3 {
		return DateData{}
	}
	year, _ := strconv.Atoi(items[0])
	month, _ := strconv.Atoi(items[1])
	day, _ := strconv.Atoi(items[2])
	t := time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.Local)
	newTime := t.AddDate(0, 0, days)
	first := convertHoursToStatic(convert(t, r), Begin)
	ends := convertHoursToStatic(convert(newTime, r), End)
	title := fmt.Sprintf("Range from %s to %s", first, ends)
	_, week := t.ISOWeek()
	return DateData{
		Name:      title,
		From:      first,
		To:        ends,
		Day:       t.Day(),
		Time:      t,
		DayOfYear: t.YearDay(),
		Week:      week,
		FromTime:  time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, time.Local),
		ToTime:    time.Date(newTime.Year(), newTime.Month(), newTime.Day(), 23, 59, 59, 0, time.Local),
	}
}

func convert(t time.Time, reading Reading) string {
	var date string
	var layout string
	switch reading {
	case Human:
		layout = "%02d/%s/%d %02d:%02d:%02d"
		date = fmt.Sprintf(layout, t.Day(), t.Month().String(), t.Year(), t.Hour(), t.Minute(), t.Second())
	case MySQL:
		layout = "%d-%02d-%02d %02d:%02d:%02d"
		date = fmt.Sprintf(layout, t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second())
	case MongoDB:
		layout = "%d-%02d-%02dT%02d:%02d:%02d"
		date = fmt.Sprintf(layout, t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second())
	}
	return date
}

func convertHoursToStatic(date string, h Hours) (result string) {
	var splitComponent string
	if strings.Contains(date, " ") {
		splitComponent = " "
	}
	if strings.Contains(date, "T") {
		splitComponent = "T"
	}
	result = strings.Split(date, splitComponent)[0]

	switch h {
	case Begin:
		result += " 00:00:00"
	case End:
		result += " 23:59:59"
	}
	if currentReading == MongoDB {
		result = strings.ReplaceAll(result, " ", "T")
		result += "Z"
	}
	return
}
