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

func GetTimeChannel() chan string {
	return timeCh
}
