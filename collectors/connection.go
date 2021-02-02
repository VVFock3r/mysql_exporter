package collectors

import (
	"database/sql"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/sirupsen/logrus"
	"fmt"
)

type ConnectionCollector struct {
	mysqlCollector
	maxConnectionDesc    *prometheus.Desc
	threadsConnectedDesc *prometheus.Desc
}

func (c *ConnectionCollector) Describe(descs chan<- *prometheus.Desc) {
	descs <- c.maxConnectionDesc
	descs <- c.threadsConnectedDesc
}

func (c *ConnectionCollector) Collect(metrics chan<- prometheus.Metric) {
	maxConnections := c.variables("max_connections")

	metrics <- prometheus.MustNewConstMetric(
		c.maxConnectionDesc,
		prometheus.CounterValue,
		maxConnections,
	)
	logrus.Debug(fmt.Sprintf("%s %f", "Metric max_connections", maxConnections))

	threadsConnected := c.status("threads_connected")

	metrics <- prometheus.MustNewConstMetric(
		c.threadsConnectedDesc,
		prometheus.CounterValue,
		threadsConnected,
	)
	logrus.Debug(fmt.Sprintf("%s %f", "Metric threads_connected", threadsConnected))

}

func NewConnectionCollector(db *sql.DB) *ConnectionCollector {
	return &ConnectionCollector{
		mysqlCollector: mysqlCollector{db},
		maxConnectionDesc: prometheus.NewDesc(
			"mysql_global_variables_max_connections",
			"MySQL global variables max connections",
			nil,
			nil,
		),
		threadsConnectedDesc: prometheus.NewDesc(
			"mysql_global_status_threads_connected",
			"MySQL global status threads connected",
			nil,
			nil,
		),
	}
}
