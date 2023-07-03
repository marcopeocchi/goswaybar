package pkg

import (
	"fmt"
	"sync"

	"github.com/marcopeocchi/goswaybar/metrics"
)

type SyncBuffer struct {
	nic         metrics.NIC
	battery     metrics.BatteryMetric
	currentTime string

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

func (b *SyncBuffer) AppendNICmetrics(metric metrics.NIC) {
	b.mu.Lock()
	defer b.mu.Unlock()

	b.nic = metric
}

func (b *SyncBuffer) GetFormatted() string {
	b.mu.RLock()
	defer b.mu.RUnlock()

	if b.currentTime == "" {
		return "Collecting data..."
	}

	return fmt.Sprintf(
		"%s | %s | %s %s",
		b.currentTime,
		b.battery.Status,
		b.nic.Ifname,
		b.nic.AddrInfo[0].Local,
	)
}
