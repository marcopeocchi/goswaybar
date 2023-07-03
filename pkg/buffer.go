package pkg

import (
	"fmt"
	"sync"

	"github.com/marcopeocchi/goswaybar/metrics"
)

type SyncBuffer struct {
	nic         *[]metrics.NIC
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

func (b *SyncBuffer) AppendNICmetrics(metric *[]metrics.NIC) {
	b.mu.Lock()
	defer b.mu.Unlock()

	b.nic = metric
}

func (b *SyncBuffer) GetFormatted() string {
	b.mu.RLock()
	defer b.mu.RUnlock()

	var up metrics.NIC

	if b.nic != nil {
		for _, n := range *b.nic {
			if n.IsInterfaceUp() {
				up = n
			}
		}
	}

	if b.currentTime != "" {
		return fmt.Sprintf(
			"%s | %s | %s %s",
			b.currentTime,
			b.battery.Status,
			up.Ifname,
			up.AddrInfo[0].Local,
		)
	}

	return "Collecting data..."
}
