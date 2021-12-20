package collector

import (
	"exporter/script"
	"github.com/prometheus/client_golang/prometheus"
)

type RessourcesCollector struct {
	UnmonitoredMemory   *prometheus.Desc
	UnmonitoredCPU      *prometheus.Desc
	ApacheMemory        *prometheus.Desc
	ApacheCPU           *prometheus.Desc
	JavaMemory          *prometheus.Desc
	JavaCPU             *prometheus.Desc
	NodeMemory          *prometheus.Desc
	NodeCPU             *prometheus.Desc
	NPMMemory           *prometheus.Desc
	NPMCPU              *prometheus.Desc
	PHPMemory           *prometheus.Desc
	PHPCPU              *prometheus.Desc
	PHPCGIMemory        *prometheus.Desc
	PHPCGICPU           *prometheus.Desc
	PythonMemory        *prometheus.Desc
	PythonCPU           *prometheus.Desc
	RedisMemory         *prometheus.Desc
	RedisCPU            *prometheus.Desc
	GunicornMemory         *prometheus.Desc
	GunicornCPU            *prometheus.Desc
}

var (
	programs = map[string]string{
		"apache": "bin/apache ",
		"java": "bin/java ",
		"node": "node ",
		"npm": "npm ",
		"php": "bin/php ",
		"phpcgi": "bin/php-cgi",
		"python": "bin/python ",
		"redis": "redis-server ",
		"gunicorn": "gunicorn"}
)

func NewRessourcesCollector() *RessourcesCollector {
	return &RessourcesCollector{
		UnmonitoredMemory: prometheus.NewDesc("ethersys_pod_unmonitored_memstats",
			"Shows the Memory used by unmonitored process",
			nil, nil,
		),
		UnmonitoredCPU: prometheus.NewDesc("ethersys_pod_unmonitored_cpustats",
			"Shows CPU usage by unmonitored process",
			nil, nil,
		),
		ApacheMemory: prometheus.NewDesc("ethersys_pod_apache_memstats",
			"Shows the Memory used by Apache ",
			nil, nil,
		),
		ApacheCPU: prometheus.NewDesc("ethersys_pod_apache_cpustats",
			"Shows CPU usage by Apache ",
			nil, nil,
		),
		JavaMemory: prometheus.NewDesc("ethersys_pod_java_memstats",
			"Shows the Memory used by Java",
			nil, nil,
		),
		JavaCPU: prometheus.NewDesc("ethersys_pod_java_cpustats",
			"Shows CPU usage by Java",
			nil, nil,
		),
		NodeMemory: prometheus.NewDesc("ethersys_pod_node_memstats",
			"Shows the Memory used by Node",
			nil, nil,
		),
		NodeCPU: prometheus.NewDesc("ethersys_pod_node_cpustats",
			"Shows CPU usage by Node",
			nil, nil,
		),
		NPMMemory: prometheus.NewDesc("ethersys_pod_npm_memstats",
			"Shows the Memory used by NPM",
			nil, nil,
		),
		NPMCPU: prometheus.NewDesc("ethersys_pod_npm_cpustats",
			"Shows CPU usage by NPM",
			nil, nil,
		),
		PHPMemory: prometheus.NewDesc("ethersys_pod_php_memstats",
			"Shows the Memory used by PHP",
			nil, nil,
		),
		PHPCPU: prometheus.NewDesc("ethersys_pod_php_cpustats",
			"Shows CPU usage by PHP",
			nil, nil,
		),
		PHPCGIMemory: prometheus.NewDesc("ethersys_pod_phpcgi_memstats",
			"Shows the Memory used by PHP-CGI",
			nil, nil,
		),
		PHPCGICPU: prometheus.NewDesc("ethersys_pod_phpcgi_cpustats",
			"Shows CPU usage by PHP-CGI",
			nil, nil,
		),
		PythonMemory: prometheus.NewDesc("ethersys_pod_python_memstats",
			"Shows the Memory used by Python",
			nil, nil,
		),
		PythonCPU: prometheus.NewDesc("ethersys_pod_python_cpustats",
			"Shows CPU usage by Python",
			nil, nil,
		),
		RedisMemory: prometheus.NewDesc("ethersys_pod_redis_memstats",
			"Shows the Memory used by Redis",
			nil, nil,
		),
		RedisCPU: prometheus.NewDesc("ethersys_pod_redis_cpustats",
			"Shows CPU usage by Redis",
			nil, nil,
		),
		GunicornMemory: prometheus.NewDesc("ethersys_pod_gunicorn_memstats",
			"Shows the Memory used by Gunicorn",
			nil, nil,
		),
		GunicornCPU: prometheus.NewDesc("ethersys_pod_gunicorn_cpustats",
			"Shows CPU usage by Gunicorn",
			nil, nil,
		),
	}
}

func (collector *RessourcesCollector) Describe(ch chan<- *prometheus.Desc) {
	ch <- collector.UnmonitoredMemory
	ch <- collector.UnmonitoredCPU
	ch <- collector.ApacheMemory
	ch <- collector.ApacheCPU
	ch <- collector.JavaMemory
	ch <- collector.JavaCPU
	ch <- collector.NodeMemory
	ch <- collector.NodeCPU
	ch <- collector.NPMMemory
	ch <- collector.NPMCPU
	ch <- collector.PHPMemory
	ch <- collector.PHPCPU
	ch <- collector.PHPCGIMemory
	ch <- collector.PHPCGICPU
	ch <- collector.PythonMemory
	ch <- collector.PythonCPU
	ch <- collector.RedisMemory
	ch <- collector.RedisCPU
	ch <- collector.GunicornMemory
	ch <- collector.GunicornCPU
}

func (collector *RessourcesCollector) Collect(ch chan<- prometheus.Metric) {
	ch <- prometheus.MustNewConstMetric(collector.UnmonitoredMemory, prometheus.CounterValue, script.UnmonitoredMemory(programs))
	ch <- prometheus.MustNewConstMetric(collector.UnmonitoredCPU, prometheus.CounterValue, script.UnmonitoredCPU(programs))
	ch <- prometheus.MustNewConstMetric(collector.ApacheMemory, prometheus.CounterValue, script.UsedMemory(programs["apache"]))
	ch <- prometheus.MustNewConstMetric(collector.ApacheCPU, prometheus.CounterValue, script.UsedCPU(programs["apache"]))
	ch <- prometheus.MustNewConstMetric(collector.JavaMemory, prometheus.CounterValue, script.UsedMemory(programs["java"]))
	ch <- prometheus.MustNewConstMetric(collector.JavaCPU, prometheus.CounterValue, script.UsedCPU(programs["java"]))
	ch <- prometheus.MustNewConstMetric(collector.NodeMemory, prometheus.CounterValue, script.UsedMemory(programs["node"]))
	ch <- prometheus.MustNewConstMetric(collector.NodeCPU, prometheus.CounterValue, script.UsedCPU(programs["node"]))
	ch <- prometheus.MustNewConstMetric(collector.NPMMemory, prometheus.CounterValue, script.UsedMemory(programs["npm"]))
	ch <- prometheus.MustNewConstMetric(collector.NPMCPU, prometheus.CounterValue, script.UsedCPU(programs["npm"]))
	ch <- prometheus.MustNewConstMetric(collector.PHPMemory, prometheus.CounterValue, script.UsedMemory(programs["php"]))
	ch <- prometheus.MustNewConstMetric(collector.PHPCPU, prometheus.CounterValue, script.UsedCPU(programs["php"]))
	ch <- prometheus.MustNewConstMetric(collector.PHPCGIMemory, prometheus.CounterValue, script.UsedMemory(programs["phpcgi"]))
	ch <- prometheus.MustNewConstMetric(collector.PHPCGICPU, prometheus.CounterValue, script.UsedCPU(programs["phpcgi"]))
	ch <- prometheus.MustNewConstMetric(collector.PythonMemory, prometheus.CounterValue, script.UsedMemory(programs["python"]))
	ch <- prometheus.MustNewConstMetric(collector.PythonCPU, prometheus.CounterValue, script.UsedCPU(programs["python"]))
	ch <- prometheus.MustNewConstMetric(collector.RedisMemory, prometheus.CounterValue, script.UsedMemory(programs["redis"]))
	ch <- prometheus.MustNewConstMetric(collector.RedisCPU, prometheus.CounterValue, script.UsedCPU(programs["redis"]))
	ch <- prometheus.MustNewConstMetric(collector.GunicornMemory, prometheus.CounterValue, script.UsedMemory(programs["gunicorn"]))
	ch <- prometheus.MustNewConstMetric(collector.GunicornCPU, prometheus.CounterValue, script.UsedCPU(programs["gunicorn"]))
}
