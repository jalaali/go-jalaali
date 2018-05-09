package jalaali

import (
	"fmt"
	"time"
)

// to_jalaali            (gy, gm, gd) {d2j, g2d}
// to_gregorian          (jy, jm, jd) {j2d, d2g}

// is_valid_jalaali_date (jy, jm, jd) {jalaali_month_length}
// jalaali_month_length  (jy, jm)     {is_leap_jalaali_year}
// is_leap_jalaali_year  (jy)         {jal_cal}
// j2d                   (jy, jm, jd) {jal_cal, g2d}
// d2j                   (jdn)        {d2g, jal_cal, g2d}
// jal_cal               (jy)
// g2d                   (gy, gm, gd)
// d2g                   (jdn)

var (
	breaks = [...]int64{-61, 9, 38, 199, 426, 686, 756, 818, 1111, 1181, 1210,
		1635, 2060, 2097, 2192, 2262, 2324, 2394, 2456, 3178}
)

type Jalaali struct {
}

func New(t *time.Time) {

}

func (j *Jalaali) toJalaali(gy, gm, gd int64) (int64, int64, int64, error) {
	jy, jm, jd, err := j.d2j(j.g2d(gy, gm, gd))
	return jy, jm, jd, err
}

func (j *Jalaali) toGregorian(jy, jm, jd int64) (int64, int64, int64, error) {
	jdn, err := j.j2d(jy, jm, jd)
	if err != nil {
		return 0, 0, 0, err
	}

	gy, gm, gd := j.d2g(jdn)
	return gy, gm, gd, nil
}

func (j *Jalaali) isValidJalaaliDate(jy, jm, jd int64) bool {
	d, err := j.jalaaliMonthLength(jy, jm)
	return err != nil &&
		jy >= -61 && jy <= 3177 &&
		jm >= 1 && jm <= 12 &&
		jd >= 1 && jd <= d
}

func (j *Jalaali) jalaaliMonthLength(jy, jm int64) (int64, error) {
	if jm <= 6 {
		return 31, nil
	} else if jm <= 11 {
		return 30, nil
	}

	leap, err := j.isLeapJalaaliYear(jy)
	if err != nil {
		return 0, err
	} else if leap {
		return 30, nil
	}
	return 29, nil
}

func (j *Jalaali) isLeapJalaaliYear(jy int64) (bool, error) {
	leap, _, _, err := j.jalCal(jy)
	return leap == 0, err
}

func (j *Jalaali) j2d(jy, jm, jd int64) (jdn int64, err error) {
	_, gy, march, err := j.jalCal(jy)
	if err != nil {
		return
	}
	return j.g2d(gy, 3, march) + (jm-1)*31 - (jm/7)*(jm-7) + jd - 1, nil
}

func (j *Jalaali) d2j(jdn int64) (jy, jm, jd int64, err error) {
	gy, _, _ := j.d2g(jdn) // Calculate Gregorian year (gy).
	jy = gy - 621
	leap, _, march, err := j.jalCal(jy)
	if err != nil {
		return
	}
	jdn1f := j.g2d(gy, 3, march)

	// Find number of days that passed since 1 Farvardin.
	k := jdn - jdn1f
	if k >= 0 {
		if k <= 185 {
			// The first 6 months.
			jm = 1 + k/31
			jd = k%31 + 1
			return
		} else {
			// The remaining months.
			k -= 186
		}
	} else {
		// Previous Jalaali year.
		jy--
		k += 179
		if leap == 1 {
			k--
		}
	}
	jm = 7 + k/30
	jd = k%30 + 1
	return
}

func (j *Jalaali) jalCal(jy int64) (leap, gy, march int64, err error) {
	b, bl := &breaks, int64(len(breaks))
	if jy < b[0] || b[bl-1] <= jy {
		return 0, 0, 0, fmt.Errorf("Wrong Jalaali year: %v", jy)
	}

	// Find the limiting years for the Jalaali year jy.
	var leapJ, jump, n, i int64
	leapJ = -14
	for i = 1; i < bl; i++ {
		jump = b[i] - b[i-1]
		if jy < b[i] {
			break
		}
		leapJ += (jump/33)*8 + (jump%33)/4
	}
	n = jy - b[i-1]

	// Find the number of leap years from AD 621 to the beginning
	// of the current Jalaali year in the Persian calendar.
	leapJ += (n/33)*8 + ((n%33)+3)/4
	if (jump%33) == 4 && (jump-n) == 4 {
		leapJ++
	}

	// And the same in the Gregorian calendar (until the year gy).
	leapG := (gy / 4) - ((gy/100+1)*3)/4 - 150

	// Determine the Gregorian date of Farvardin the 1st.
	march = 20 + leapJ - leapG

	// Find how many years have passed since the last leap year.
	if jump-n < 6 {
		n -= jump + ((jump+4)/33)*33
	}
	leap = (((n + 1) % 33) - 1) % 4
	if leap == -1 {
		leap = 4
	}

	return
}

func (j *Jalaali) g2d(gy, gm, gd int64) (jdn int64) {
	a := (gm - 8) / 6
	gy = gy + a + 100100
	gd = gd - 34839656
	jdn = (gy*1461)/4 + (153*((gm+9)%12)+2)/5 - ((gy/100)*3)/4 + gd

	return
}

func (j *Jalaali) d2g(jdn int64) (gy, gm, gd int64) {
	jdn = 4 * jdn
	i := jdn + 139361631 + (((jdn+183187720)/146097)*3/4)*4 - 3908
	jdn = ((i%1461)/4)*5 + 308

	gd = (jdn%153)/5 + 1
	gm = (jdn/153)%12 + 1
	gy = (i / 1461) - 100100 + (8-gm)/6

	return
}
