package buffer

import (
	"fmt"
	"sync"

	"github.com/marcopeocchi/goswaybar/pkg/metrics"
)

type SyncBuffer struct {
	nic         metrics.NIC
	battery     metrics.BatteryMetric
	currentTime string
	volume      string

	mu sync.RWMutex
}

func NewSyncBuffer() *SyncBuffer {
	return &SyncBuffer{}
}

func (b *SyncBuffer) AppendBatteryLevel(metric metrics.BatteryMetric) {
	b.mu.Lock()
	defer b.mu.Unlock()

	b.battery = metric
}

func (b *SyncBuffer) AppendCurrentTime(metric string) {
	b.mu.Lock()
	defer b.mu.Unlock()

	b.currentTime = metric
}

func (b *SyncBuffer) AppendNICMetrics(metric metrics.NIC) {
	b.mu.Lock()
	defer b.mu.Unlock()

	b.nic = metric
}

func (b *SyncBuffer) AppendVolumeMetrics(metric string) {
	b.mu.Lock()
	defer b.mu.Unlock()

	b.volume = metric
}

func (b *SyncBuffer) GetFormatted() string {
	b.mu.RLock()
	defer b.mu.RUnlock()

	if b.currentTime == "" {
		return "Collecting data..."
	}

	var batPercent string

	if b.battery.Charging {
		batPercent = "<span foreground=\"green\">" + b.battery.Status + "%</span>"
	} else {
		batPercent = "<span foreground=\"red\">" + b.battery.Status + "%</span>"
	}

	return fmt.Sprintf(
		"%s | BAT: %s | NET: %s %s",
		b.currentTime,
		batPercent,
		b.nic.Ifname,
		b.nic.AddrInfo[0].Local,
	)
}
