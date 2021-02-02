package collectors

import (
	"database/sql"
	"fmt"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/sirupsen/logrus"
)

// 慢查询收集器
// 慢查询信息：
//	show variables like 'slow_query%';
//	show variables like 'long_query_time';
// 开启慢查询：set global slow_query_log='ON';
// 构造慢查询：select sleep(11);
type SlowQueriesCollector struct {
	mysqlCollector
	desc *prometheus.Desc
}

// 写入指标
func (c *SlowQueriesCollector) Describe(descs chan<- *prometheus.Desc) {
	descs <- c.desc
}

// 写入数据
func (c *SlowQueriesCollector) Collect(metrics chan<- prometheus.Metric) {
	count := c.status("Slow_queries")
	logrus.Debug(fmt.Sprintf("%s %f", "Metric SlowQueries", count))
	metrics <- prometheus.MustNewConstMetric(c.desc, prometheus.CounterValue, count)
}

// 构造函数
func NewSlowQueriesCollector(db *sql.DB) *SlowQueriesCollector {
	return &SlowQueriesCollector{
		mysqlCollector: mysqlCollector{db},
		desc: prometheus.NewDesc(
			"mysql_global_status_slow_queries",
			"MySQL global status slow queries",
			nil,
			nil,
		),
	}
}

// QPS收集器
type QpsCollector struct {
	mysqlCollector
	desc *prometheus.Desc
}

func (c *QpsCollector) Describe(descs chan<- *prometheus.Desc) {
	descs <- c.desc
}

func (c *QpsCollector) Collect(metrics chan<- prometheus.Metric) {
	count := c.status("queries")
	logrus.Debug(fmt.Sprintf("%s %f", "Metric QPS", count))
	metrics <- prometheus.MustNewConstMetric(c.desc, prometheus.CounterValue, count)
}

func NewQpsCollector(db *sql.DB) *QpsCollector {
	return &QpsCollector{
		mysqlCollector: mysqlCollector{db},
		desc: prometheus.NewDesc(
			"mysql_global_status_queries",
			"MySQL global status Queries",
			nil,
			nil,
		),
	}
}
