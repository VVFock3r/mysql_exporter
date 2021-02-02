package collectors

import (
	"database/sql"
	"fmt"
	"github.com/prometheus/client_golang/prometheus"
)

// （查询、插入、更新、删除）执行次数收集器
type CommandCollector struct {
	mysqlCollector
	desc *prometheus.Desc
}

// 写入指标
func (c *CommandCollector) Describe(descs chan<- *prometheus.Desc) {
	descs <- c.desc
}

// 写入数据
func (c *CommandCollector) Collect(metrics chan<- prometheus.Metric) {
	names := []string{
		"insert",
		"update",
		"delete",
		"select",
	}

	// 这里有个严重的性能问题
	// 假设连接超时时间为5s，那么这里就要阻塞20s才能返回
	// 暂时解决方法为 先检查一下连接，再查询
	if c.db.Ping() != nil {
		for _, name := range names {
			metrics <- prometheus.MustNewConstMetric(c.desc, prometheus.CounterValue, 0, name)
		}
	} else {
		for _, name := range names {
			count := c.status(fmt.Sprintf("Com_%s", name))
			metrics <- prometheus.MustNewConstMetric(c.desc, prometheus.CounterValue, count, name)
		}
	}
}

// 构造函数
func NewCommandCollector(db *sql.DB) *CommandCollector {
	return &CommandCollector{
		mysqlCollector: mysqlCollector{db},
		desc: prometheus.NewDesc(
			"mysql_global_status_command",
			"MySQL global status command",
			[]string{"command"},
			nil,
		),
	}
}
