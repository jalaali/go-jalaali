# Jalaali

Golang implementation of [Jalaali JS](https://github.com/jalaali/jalaali-js) and [Jalaali Python](https://github.com/jalaali/jalaali-python) implementations of Jalaali (Jalali, Persian, Khayyami, Khorshidi, Shamsi) convertion to Gregorian calendar system and vice-versa.

This implementation is based on an [algorithm by Kazimierz M. Borkowski](http://www.astro.uni.torun.pl/~kb/Papers/EMP/PersianC-EMP.htm). Borkowski claims that this algorithm works correctly for 3000 years!

Documentation on API is available [here](https://pkg.go.dev/github.com/jalaali/go-jalaali) at Go official documentation site.

## Installation

Use `go get` on this repository:

```sh
$ go get -u github.com/jalaali/go-jalaali
```

## Usage

### Current time

`Jalaali.Now()` returns a `Jalaali` instance wrapping the current time. You can use all the methods from the Go `time` package on it.

```go
package main

import (
	"jalaali"
)

func main() {
	j := jalaali.Now()
}
```

### From a `time` instance

`Jalaali.From(t)` wraps an existing `time.Time` instance. You can then work with it the same way you work with the `time` package.

```go
package main

import (
	"jalaali"
	"time"
)

func main() {
	t := time.Date(2014, 8, 4, 0, 0, 0, 0, time.UTC)
	j := jalaali.From(t)
}
```

### Converting dates

Convert a Gregorian date to Jalaali with `ToJalaali`, and back with `ToGregorian`. Both return the date components plus an `error` that is non-nil when the input is invalid.

```go
package main

import (
	"jalaali"
	"time"
)

func main() {
	// Gregorian to Jalaali
	jy, jm, jd, err := jalaali.ToJalaali(2014, time.August, 4)
	if err != nil {
		panic(err)
	}
	// jy = 1393, jm = jalaali.Mordad, jd = 13

	// Jalaali to Gregorian
	gy, gm, gd, err := jalaali.ToGregorian(1393, jalaali.Mordad, 13)
	if err != nil {
		panic(err)
	}
	// gy = 2014, gm = time.August, gd = 4
}
```

### Formatting

Call the `JFormat` method of a `Jalaali` instance and pass it the same formatting layout used by the Go `time` package. The output is a Jalaali date using Persian digits and words.

```go
package main

import (
	"jalaali"
	"time"
)

func main() {
	iran, _ := time.LoadLocation("Asia/Tehran")
	j := jalaali.From(time.Date(2001, 2, 3, 15, 17, 1, 0, iran))

	// "۱۳۷۹ بهمن شنبه ۱۵ ۰۳ ۱۷ ۰۱ بعدازظهر"
	formatted, err := j.JFormat("2006 January Monday 15 03 17 01 PM")
	if err != nil {
		panic(err)
	}
}
```

### Utilities

- `IsValidDate(jy, jm, jd)` returns `true` if the given Jalaali date is valid.
- `IsLeapYear(jy)` returns `true` if the given Jalaali year is a leap year.
- `MonthLength(jy, jm)` returns the number of days in the given Jalaali month.
