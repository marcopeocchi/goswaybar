package metrics

import (
	"strings"
	"time"
)

var (
	timeCh = make(chan string)
)

func CollectTimeMetrics() {
	timeCh <- strings.ReplaceAll(time.Now().Format(time.ANSIC), "  ", " ")
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
