package main

import (
	"time"
	"testing"
)

const (
	ExpectedTime = "2016-12-24T18:14:54.22CST"
)

func TestFormatTimeInZone(t *testing.T) {
	christmasUtc := time.Date(2016, time.December, 24, 10, 11, 12, 222222222222, time.UTC)
	zoneShanghai, _ := time.LoadLocation("Asia/Shanghai")

	shanghaiTime := formatTimeInZone(christmasUtc, zoneShanghai)

	if shanghaiTime == "" {
		t.Error("formatted time was empty")
		// t.Error() does not stop the test! it just records it as one of the reasons of failure
	}
	if shanghaiTime != ExpectedTime {
		t.Errorf("time was %v, expected %v", shanghaiTime, ExpectedTime)
	}
}


func BenchmarkFormatTimeInZone(b *testing.B) {
	now := time.Now()
	zoneShanghai, _ := time.LoadLocation("Asia/Shanghai")
	for i := 0; i < b.N; i++ {
		formatTimeInZone(now, zoneShanghai)
	}
}