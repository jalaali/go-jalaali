package jalaali

import (
	"strconv"
	"time"
)

// From initialize new instance of Jalaali from a time instance.
func From(t time.Time) Jalaali {
	return Jalaali{t}
}

// Now with return Jalaali instance of current time.
func Now() Jalaali {
	return From(time.Now())
}

// Jalaali is a simple wrapper around Golang default time package. You have all
// the functionality of default time package and functionalities needed for
// Jalaali calender.
type Jalaali struct {
	time.Time
}

func (j Jalaali) Weekday() Weekday {
	return Weekday((int(j.Time.Weekday()) + 1) % 7)
}

func (j Jalaali) Date() (year int, month Month, day int, err error) {
	return ToJalaali(j.Time.Date())
}

// A Month specifies a month of the year (Farvardin = 1, ...).
type Month int

const (
	Farvardin Month = 1 + iota
	Ordibehesht
	Khordad
	Tir
	Mordad
	Shahrivar
	Mehr
	Aban
	Azar
	Dey
	Bahman
	Esfand
)

var months = []string{
	"فروردین", "اردیبهشت", "خرداد",
	"تیر", "مرداد", "شهریور",
	"مهر", "آبان", "آذر",
	"دی", "بهمن", "اسفند",
}

func (m Month) String() string {
	if Farvardin <= m && m <= Esfand {
		return months[m-1]
	}
	return "%!Month(" + strconv.Itoa(int(m)) + ")"
}

// A Weekday specifies a day of the week (Shanbe = 0, ...).
type Weekday int

const (
	Shanbe Weekday = iota
	IekShanbe
	DoShanbe
	SeShanbe
	ChaharShanbe
	PanjShanbe
	Jome

	YekShanbe = IekShanbe
)

var days = []string{
	"شنبه", "یک‌شنبه", "دوشنبه", "سه‌شنبه", "چهارشنبه", "پنج‌شنبه", "جمعه",
}

func (d Weekday) String() string {
	if Shanbe <= d && d <= Jome {
		return days[d]
	}
	return "%!Weekday(" + strconv.Itoa(int(d)) + ")"
}
