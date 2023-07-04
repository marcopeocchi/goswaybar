package main

import (
	"fmt"
	"time"

	"github.com/marcopeocchi/goswaybar/pkg/buffer"
	"github.com/marcopeocchi/goswaybar/pkg/metrics"
)

func main() {
	// Collectors
	go metrics.CollectTimeMetricsPeriodically()
	go metrics.CollectNICMetricsPeriodically(time.Second * 5)
	go metrics.CollectBatteryMetricsPeriodically(time.Second * 10)

	// Common buffer
	b := buffer.NewSyncBuffer()

	// Retrieve from collectors channels
	go func() {
		for {
			select {
			case m := <-metrics.GetBatteryChannel():
				b.AppendBatteryLevel(m)
			case m := <-metrics.GetTimeChannel():
				b.AppendCurrentTime(m)
			case m := <-metrics.GetNICChannel():
				b.AppendNICMetrics(m)
			}
		}
	}()

	// Main Goroutine prints formatted text
	for {
		fmt.Println(b.GetFormatted())
		time.Sleep(time.Second)
	}
}
