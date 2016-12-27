package main

import (
	"testing"
	"time"
)

const ExpectedTime = "2016-12-24T18:14:54.22CST"

func TestFormatTimeInZone(t *testing.T) {
	christmasUtc := time.Date(2016, time.December, 24, 10, 11, 12, 222222222222, time.UTC)
	zoneShanghai, _ := time.LoadLocation("Asia/Shanghai")

	shanghaiTime := formatTimeInZone(christmasUtc, zoneShanghai)

	if shanghaiTime != ExpectedTime {
		t.Errorf("time was %v, expected %v", shanghaiTime, ExpectedTime)
	}
}
