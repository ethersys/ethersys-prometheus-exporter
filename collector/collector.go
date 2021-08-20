package collector

import (
	"exporter/script"

	"github.com/prometheus/client_golang/prometheus"
)

type MemoryCollector struct {
	ApacheMemory *prometheus.Desc
	JavaMemory   *prometheus.Desc
	NodeMemory   *prometheus.Desc
	NPMMemory    *prometheus.Desc
	PHPMemory    *prometheus.Desc
	RedisMemory  *prometheus.Desc
}

var (
	programmeList = [6]string{"bin/php",
		"node ",
		"npm ",
		"bin/java",
		"redis-server",
		"bin/apache"}
)

func NewMemoryCollector() *MemoryCollector {
	return &MemoryCollector{
		ApacheMemory: prometheus.NewDesc("ethersys_pod_apache_memstats",
			"Shows the Memory used by Apache ",
			nil, nil,
		),
		JavaMemory: prometheus.NewDesc("ethersys_pod_java_memstats",
			"Shows the Memory used by Java",
			nil, nil,
		),
		NodeMemory: prometheus.NewDesc("ethersys_pod_node_memstats",
			"Shows the Memory used by Node",
			nil, nil,
		),
		NPMMemory: prometheus.NewDesc("ethersys_pod_npm_memstats",
			"Shows the Memory used by NPM",
			nil, nil,
		),
		PHPMemory: prometheus.NewDesc("ethersys_pod_php_memstats",
			"Shows the Memory used by PHP",
			nil, nil,
		),
		RedisMemory: prometheus.NewDesc("ethersys_pod_redis_memstats",
			"Shows the Memory used by Redis",
			nil, nil,
		),
	}
}

func (collector *MemoryCollector) Describe(ch chan<- *prometheus.Desc) {

	ch <- collector.ApacheMemory
	ch <- collector.JavaMemory
	ch <- collector.NodeMemory
	ch <- collector.NPMMemory
	ch <- collector.PHPMemory
	ch <- collector.RedisMemory
}

func (collector *MemoryCollector) Collect(ch chan<- prometheus.Metric) {

	ch <- prometheus.MustNewConstMetric(collector.ApacheMemory, prometheus.CounterValue, script.UsedMemory(programmeList[0]))
	ch <- prometheus.MustNewConstMetric(collector.JavaMemory, prometheus.CounterValue, script.UsedMemory(programmeList[1]))
	ch <- prometheus.MustNewConstMetric(collector.NodeMemory, prometheus.CounterValue, script.UsedMemory(programmeList[2]))
	ch <- prometheus.MustNewConstMetric(collector.NPMMemory, prometheus.CounterValue, script.UsedMemory(programmeList[3]))
	ch <- prometheus.MustNewConstMetric(collector.PHPMemory, prometheus.CounterValue, script.UsedMemory(programmeList[4]))
	ch <- prometheus.MustNewConstMetric(collector.RedisMemory, prometheus.CounterValue, script.UsedMemory(programmeList[5]))
}
