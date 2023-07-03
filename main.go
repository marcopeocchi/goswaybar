package main

import (
	"fmt"
	"time"

	"github.com/marcopeocchi/goswaybar/metrics"
	"github.com/marcopeocchi/goswaybar/pkg"
)

func main() {
	go func() {
		for {
			metrics.CollectBatteryMetrics()
			time.Sleep(time.Second * 10)
		}
	}()
	go func() {
		for {
			metrics.CollectTimeMetrics()
			time.Sleep(time.Second)
		}
	}()
	go func() {
		for {
			metrics.CollectNICMetrics()
			time.Sleep(time.Second * 5)
		}
	}()

	b := pkg.NewSyncBuffer()

	go func() {
		for {
			select {
			case m := <-metrics.GetBatteryChannel():
				b.AppendBatteryLevel(m)
			case m := <-metrics.GetTimeChannel():
				b.AppendCurrentTime(m)
			case m := <-metrics.GetNICChannel():
				b.AppendNICmetrics(m)
			}
		}
	}()

	for {
		fmt.Println(b.GetFormatted())
		time.Sleep(time.Second)
	}
}
