package jalaali

import "testing"

func TestFromYMD(t *testing.T) {
	tests := []struct {
		gy, gm, gd, jy, jm, jd int
	}{
		{1981, 8, 17, 1360, 5, 26},
		{2013, 1, 10, 1391, 10, 21},
		{2014, 8, 4, 1393, 5, 13},
	}

	for _, test := range tests {
		y, m, d, err := FromYMD(test.gy, test.gm, test.gd)
		if err != nil {
			t.Errorf("%v", err)
		} else if y != test.jy || m != test.jm || d != test.jd {
			t.Errorf("Expected %v/%v/%v got %v/%v%v.", test.jy, test.jm, test.jd, y, m, d)
		}
	}

}

func TestToGregorian(t *testing.T) {
	tests := []struct {
		jy, jm, jd, gy, gm, gd int
	}{
		{1360, 5, 26, 1981, 8, 17},
		{1391, 10, 21, 2013, 1, 10},
		{1393, 5, 13, 2014, 8, 4},
	}

	for _, test := range tests {
		y, m, d, err := ToGregorian(test.jy, test.jm, test.jd)
		if err != nil {
			t.Errorf("%v", err)
		} else if y != test.gy || m != test.gm || d != test.gd {
			t.Errorf("Expected %v/%v/%v got %v/%v%v.", test.gy, test.gm, test.gd, y, m, d)
		}
	}

}

func TestIsValidDate(t *testing.T) {
	tests := []struct {
		y, m, d int
		ok      bool
	}{
		{-62, 12, 29, false},
		{-61, 1, 1, true},
		{3178, 1, 1, false},
		{3177, 12, 29, true},
		{1393, 0, 1, false},
		{1393, 13, 1, false},
		{1393, 1, 0, false},
		{1393, 1, 32, false},
		{1393, 1, 31, true},
		{1393, 11, 31, false},
		{1393, 11, 30, true},
		{1393, 12, 30, false},
		{1393, 12, 29, true},
		{1395, 12, 30, true},
	}

	for _, test := range tests {
		valid, err := IsValidDate(test.y, test.m, test.d)
		if err != nil && test.ok {
			t.Errorf("%v", err)
		} else if valid != test.ok {
			calculated, actual := "", " not"
			if test.ok {
				calculated, actual = " not", ""
			}
			t.Errorf("%v/%v/%v is%v valid date but considered%v valid.",
				test.y, test.m, test.d, actual, calculated)
		}
	}
}

func TestIsLeapYear(t *testing.T) {
	tests := []struct {
		year int
		leap bool
	}{
		{1393, false},
		{1394, false},
		{1395, true},
		{1396, false},
	}

	for _, test := range tests {
		leap, err := IsLeapYear(test.year)
		if err != nil {
			t.Errorf("%v", err)
		} else if leap != test.leap {
			calculated, actual := "", " not"
			if leap {
				calculated, actual = " not", ""
			}
			t.Errorf("%v is%v leap but considered%v leap.", test.year, actual, calculated)
		}
	}
}

func TestMonthLength(t *testing.T) {
	tests := []struct {
		y, m, ml int
	}{
		{1393, 1, 31},
		{1393, 4, 31},
		{1393, 6, 31},
		{1393, 7, 30},
		{1393, 10, 30},
		{1393, 12, 29},
		{1394, 12, 29},
		{1395, 12, 30},
	}

	for _, test := range tests {
		calculated, err := MonthLength(test.y, test.m)
		if err != nil {
			t.Errorf("%v", err)
		} else if calculated != test.ml {
			t.Errorf("Length of %v/%v month is %v but considered %v.",
				test.y, test.m, test.ml, calculated)
		}
	}
}
