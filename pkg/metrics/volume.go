package metrics

import (
	"os/exec"
	"regexp"
	"strings"
	"time"
)

var (
	volumeCh = make(chan string)
	volumeRe = regexp.MustCompile(`(?m)\[(.*?)\]`)
	replacer = strings.NewReplacer("[", "", "]", "")
)

func CollectVolumeMetrics() error {
	cmd := exec.Command("amixer", "sget", "Master")

	buff, err := cmd.Output()
	if err != nil {
		return err
	}

	volumeCh <- replacer.Replace(string(volumeRe.Find(buff)))
	return err
}

func CollectVolumePeriodically(d time.Duration) {
	for {
		CollectVolumeMetrics()
		time.Sleep(d)
	}
}

func GetVolumeChannel() chan string {
	return volumeCh
}
