package metrics

import (
	"time"
)

var (
	timeCh = make(chan string)
)

func CollectTimeMetrics() {
	timeCh <- time.Now().Format("2 Jan 2006 03:04PM")
}

func CollectTimeMetricsPeriodically() {
	for {
		CollectTimeMetrics()
		time.Sleep(time.Second)
	}
}

func GetTimeChannel() chan string {
	return timeCh
}
