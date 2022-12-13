package spidate

import (
	"time"
)

type Reading int
type Options int
type Hours int

const (
	Default Reading = iota
	Human
	MySQL
	MongoDB
)
const (
	Date Options = iota
	DateTime
)

const (
	Begin Hours = iota
	End
)

type DateData struct {
	Name      string    `json:"name,omitempty" bson:"name,omitempty"`
	From      string    `json:"fromDate,omitempty" bson:"fromDate,omitempty"`
	To        string    `json:"toDate,omitempty" bson:"toDate,omitempty"`
	Day       int       `json:"day,omitempty" bson:"day,omitempty"`
	Time      time.Time `json:"time,omitempty" bson:"time,omitempty"`
	DayOfYear int       `json:"dayOfYear,omitempty" bson:"dayOfYear,omitempty"`
	Week      int       `json:"week,omitempty" bson:"week,omitempty"`
	FromTime  time.Time `json:"fromTime,omitempty" bson:"fromTime,omitempty"`
	ToTime    time.Time `json:"toTime,omitempty" bson:"toTime,omitempty"`
}
