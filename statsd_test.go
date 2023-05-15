package statsd

import (
	"testing"
	"time"
)

func initConfig() {
	config := &Config{
		Project: "stats",
		Host:    "192.168.120.100",
		Port:    1675,
		Enable:  true,
	}

	Setup(config)
}

func Test_Gauge(t *testing.T) {
	initConfig()

	var v int64

	ticker := time.NewTicker(100 * time.Millisecond)
	for {
		<-ticker.C

		v++
		Gauge("statsd.gauge", v)

		if v == 10 {
			break
		}
	}
}

func Test_Incr(t *testing.T) {
	initConfig()

	ticker := time.NewTicker(100 * time.Millisecond)

	cnt := 0
	for {
		<-ticker.C

		Incr("statsd.incr")

		if cnt == 20 {
			break
		}
	}
}
