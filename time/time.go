package time

import (
	time2 "time"
)

func ParseTime(t string) (rv time2.Duration) {

	t0 := 0
	for k, _ := range t {
		v := int(t[k])
		switch {
		case v >= '0' && v <= '9':
			t0 = t0*10 + (v - '0')
		case v == 's':
			rv += time2.Duration(t0) * time2.Second
			t0 = 0
		case v == 'm':
			if k+1 < len(t) && t[k+1] == 's' {
				rv += time2.Duration(t0) * time2.Millisecond
				t0 = 0
				k++
				continue
			}
			rv += time2.Duration(t0*60) * time2.Second
			t0 = 0
		case v == 'h':
			rv += time2.Duration(t0*60*60) * time2.Second
			t0 = 0
		case v == 'd':
			rv += time2.Duration(t0*60*60*24) * time2.Second
			t0 = 0
		case v == 'w':
			rv += time2.Duration(t0*60*60*24*7) * time2.Second
			t0 = 0
		case v == 'M':
			rv += time2.Duration(t0*60*60*24*7*31) * time2.Second
			t0 = 0
		case v == 'y':
			rv += time2.Duration(t0*60*60*24*7*31*365) * time2.Second
			t0 = 0
		}
	}

	return
}
