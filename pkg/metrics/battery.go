package metrics

import (
	"bytes"
	"os"
	"time"
)

type BatteryMetric struct {
	ETA      string
	Status   string
	Charging bool
}

var (
	batCh    = make(chan BatteryMetric)
	charging = []byte("Charging")
)

const (
	capacity string = "/sys/class/power_supply/BAT0/capacity"
	status   string = "/sys/class/power_supply/BAT0/status"
)

func CollectBatteryMetrics() error {
	buff, err := os.ReadFile(capacity)
	if err != nil {
		return err
	}

	m := BatteryMetric{
		Status: string(buff[:len(buff)-1]),
	}

	buff, err = os.ReadFile(status)
	if err != nil {
		return err
	}

	m.Charging = bytes.Equal(buff[:len(buff)-1], charging)

	batCh <- m
	return nil
}

func CollectBatteryMetricsPeriodically(d time.Duration) {
	for {
		CollectBatteryMetrics()
		time.Sleep(d)
	}
}

func GetBatteryChannel() chan BatteryMetric {
	return batCh
}
