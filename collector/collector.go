package collector

import (
	"exporter/script"

	"github.com/prometheus/client_golang/prometheus"
)

type RamCollector struct {
	PHPRam    *prometheus.Desc
	ApacheRam *prometheus.Desc
	NPMRam    *prometheus.Desc
	JavaRam   *prometheus.Desc
	RedisRam  *prometheus.Desc
	NodeRam   *prometheus.Desc
}

var (
	programmeList = [6]string{"bin/php",
		"node ",
		"npm ",
		"bin/java",
		"redis-server",
		"bin/apache"}
)

func NewwRamCollector() *RamCollector {
	return &RamCollector{
		PHPRam: prometheus.NewDesc("PHP_ram",
			"Shows the RAM used by PHP",
			nil, nil,
		),
		ApacheRam: prometheus.NewDesc("Apache_ram",
			"Shows the RAM used by Apache ",
			nil, nil,
		),
		NPMRam: prometheus.NewDesc("NPM_ram",
			"Shows the RAM used by NPM",
			nil, nil,
		),
		JavaRam: prometheus.NewDesc("Java_ram",
			"Shows the RAM used by Java",
			nil, nil,
		),
		RedisRam: prometheus.NewDesc("Redis_ram",
			"Shows the RAM used by Redis",
			nil, nil,
		),
		NodeRam: prometheus.NewDesc("Node__ram",
			"Shows the RAM used by Node",
			nil, nil,
		),
	}
}

func (collector *RamCollector) Describe(ch chan<- *prometheus.Desc) {

	ch <- collector.PHPRam
	ch <- collector.ApacheRam
	ch <- collector.NPMRam
	ch <- collector.JavaRam
	ch <- collector.NodeRam
	ch <- collector.RedisRam
}

func (collector *RamCollector) Collect(ch chan<- prometheus.Metric) {

	ch <- prometheus.MustNewConstMetric(collector.PHPRam, prometheus.CounterValue, script.UsedRam(programmeList[0]))
	ch <- prometheus.MustNewConstMetric(collector.ApacheRam, prometheus.CounterValue, script.UsedRam(programmeList[1]))
	ch <- prometheus.MustNewConstMetric(collector.NPMRam, prometheus.CounterValue, script.UsedRam(programmeList[2]))
	ch <- prometheus.MustNewConstMetric(collector.JavaRam, prometheus.CounterValue, script.UsedRam(programmeList[3]))
	ch <- prometheus.MustNewConstMetric(collector.RedisRam, prometheus.CounterValue, script.UsedRam(programmeList[4]))
	ch <- prometheus.MustNewConstMetric(collector.NodeRam, prometheus.CounterValue, script.UsedRam(programmeList[5]))
}
