package metrics

import (
	"encoding/json"
	"os/exec"
	"time"
)

type NIC struct {
	Ifindex   int64      `json:"ifindex"`
	Ifname    string     `json:"ifname"`
	Flags     []string   `json:"flags"`
	MTU       int64      `json:"mtu"`
	Qdisc     string     `json:"qdisc"`
	Operstate string     `json:"operstate"`
	Group     string     `json:"group"`
	Txqlen    int64      `json:"txqlen"`
	LinkType  string     `json:"link_type"`
	Address   string     `json:"address"`
	Broadcast string     `json:"broadcast"`
	AddrInfo  []AddrInfo `json:"addr_info"`
}

type AddrInfo struct {
	Family            string `json:"family"`
	Local             string `json:"local"`
	Prefixlen         int64  `json:"prefixlen"`
	Scope             string `json:"scope"`
	ValidLifeTime     int64  `json:"valid_life_time"`
	PreferredLifeTime int64  `json:"preferred_life_time"`
}

func (n *NIC) IsInterfaceUp() bool {
	return n.Operstate == "UP"
}

var (
	nicCh = make(chan NIC)
)

func CollectNICMetrics() error {
	cmd := exec.Command("ip", "-j", "addr")

	out, err := cmd.Output()
	if err != nil {
		return err
	}

	n := new([]NIC)
	err = json.Unmarshal(out, &n)
	if err != nil {
		return err
	}

	for _, e := range *n {
		if e.IsInterfaceUp() {
			nicCh <- e
		}
	}
	return err
}

func CollectNICMetricsPeriodically(d time.Duration) {
	for {
		CollectNICMetrics()
		time.Sleep(d)
	}
}

func GetNICChannel() chan NIC {
	return nicCh
}
