package metrics

import (
	"bufio"
	"os/exec"
	"strings"
	"time"
)

type BatteryMetric struct {
	ETA      string
	Status   string
	Charging bool
}

var (
	batCh = make(chan BatteryMetric)
)

func CollectBatteryMetrics() error {
	batCmd := exec.Command("acpi", "-b")

	r, err := batCmd.StdoutPipe()
	if err != nil {
		return err
	}

	scan := bufio.NewScanner(r)

	go func() {
		for scan.Scan() {
			data := strings.Split(scan.Text(), ": ")[1]
			stats := strings.Split(data, ",")

			m := BatteryMetric{
				Charging: strings.TrimSpace(stats[0]) == "Charging",
				Status:   strings.TrimSpace(stats[1]),
			}

			if len(stats) > 2 {
				m.ETA = strings.TrimSpace(stats[2])
			}

			if !strings.HasPrefix(m.Status, "0") {
				batCh <- m
			}
		}
	}()

	return batCmd.Run()
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
