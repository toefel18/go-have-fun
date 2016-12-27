package main

import (
	"testing"
	"time"
)

func BenchmarkFormatTimeInZone(b *testing.B) {
	now := time.Now()
	zoneShanghai, _ := time.LoadLocation("Asia/Shanghai")
	for i := 0; i < b.N; i++ {
		formatTimeInZone(now, zoneShanghai)
	}
}
