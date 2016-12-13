package util

import (
	"time"
)

var (
	pow10tab [19]time.Duration
)

func init() {
	pow10tab[0] = 1
	for i := 1; i < len(pow10tab); i++ {
		pow10tab[i] = 10 * pow10tab[i-1]
	}
}

func round(d, r time.Duration) time.Duration {
	if r <= 0 {
		return d
	}
	neg := d < 0
	if neg {
		d = -d
	}
	if m := d % r; m+m < r {
		d = d - m
	} else {
		d = d + r - m
	}
	if neg {
		return -d
	}
	return d
}

func roundN(d time.Duration, n int) time.Duration {
	if n < 1 {
		return d
	}
	if d >= time.Hour {
		k := digits(d / time.Hour)
		if k >= n {
			return round(d, pow10(time.Hour, k-n))
		}
		n -= k
		k = digits(d % time.Hour / time.Minute)
		if k >= n {
			return round(d, pow10(time.Minute, k-n))
		}
		return round(d, pow10(100*time.Second, k-n))
	}
	if d >= time.Minute {
		k := digits(d / time.Minute)
		if k >= n {
			return round(d, pow10(time.Minute, k-n))
		}
		return round(d, pow10(100*time.Second, k-n))
	}
	if k := digits(d); k > n {
		return round(d, pow10(1, k-n))
	}
	return d
}

func digits(d time.Duration) int {
	if d < 0 {
		d = -d
	}
	i := 1
	for d > 9 {
		d /= 10
		i++
	}
	return i
}

func pow10(d time.Duration, i int) time.Duration {
	if i < 0 {
		return d / pow10tab[-i]
	}
	return d * pow10tab[i]
}
