package main

import (
	"database/sql"
	"fmt"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/sirupsen/logrus"
	"gopkg.in/natefinch/lumberjack.v2"
	"io"
	"log"
	"mysql_exporter/collectors"
	"mysql_exporter/config"
	"mysql_exporter/handler"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

func initLogger(options config.Logger) ([]io.Closer, error) {
	// 设置日志级别
	level, err := logrus.ParseLevel(options.Global.Level)
	if err != nil {
		return nil, err
	}
	logrus.SetLevel(level)

	// 设置日志格式
	logrus.SetFormatter(&config.LogFormatter{})

	// 设置输出终端
	var writers []io.Writer
	var closes []io.Closer // 抛出需要defer的

	if options.Stdout.Enabled {
		writers = append(writers, os.Stdout)
	}
	if options.Stderr.Enabled {
		writers = append(writers, os.Stderr)
	}

	if options.File.Enabled {
		logger := &lumberjack.Logger{
			Filename:   options.File.FileName,
			MaxSize:    options.File.MaxSize,
			MaxAge:     options.File.MaxAge,
			MaxBackups: options.File.MaxBackups,
			Compress:   options.File.Compress,
		}
		writers = append(writers, logger)
		closes = append(closes, logger)
	}

	if len(writers) == 0 {
		logrus.SetOutput(os.Stderr)
	} else {
		logrus.SetOutput(io.MultiWriter(writers...))
	}

	return closes, nil
}

func initMySQL(options config.MySQL) (*sql.DB, error) {
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?charset=%s&loc=%s&timeout=%s&readTimeout=%s",
		options.Username,
		options.Password,
		options.Host,
		options.Port,
		options.Db,
		options.CharSet,
		options.TimeZone,
		options.ConnTimeout,
		options.ReadTimeout,
	)

	return sql.Open("mysql", dsn)
}

func initMetrics(options *config.Options, db *sql.DB) {
	prometheus.MustRegister(
		prometheus.NewGaugeFunc(
			prometheus.GaugeOpts{
				Name: "mysql_up",
				Help: "Mysql Up Info",
				ConstLabels: prometheus.Labels{
					"addr": options.MySQL.Host,
				},
			},
			func() float64 {
				if err := db.Ping(); err != nil {
					logrus.Error("failed to connect to mysql: ", fmt.Sprintf("%s:%d", options.MySQL.Host, options.MySQL.Port))
					return 0
				}
				return 1
			},
		),
	)

	// 慢查询数
	prometheus.MustRegister(collectors.NewSlowQueriesCollector(db))
	logrus.Debug("register collector: SlowQueriesCollector")

	// 总QPS
	prometheus.MustRegister(collectors.NewQpsCollector(db))
	logrus.Debug("register collector: QpsCollector")

	// 执行次数收集器
	prometheus.MustRegister(collectors.NewCommandCollector(db))
	logrus.Debug("register collector: CommandCollector")

	// 连接数
	prometheus.MustRegister(collectors.NewConnectionCollector(db))
	logrus.Debug("register collector: ConnectionCollector")

	// 流量
	prometheus.MustRegister(collectors.NewTrafficCollector(db))
	logrus.Debug("register collector: TrafficCollector")
}

func main() {
	// 加载配置
	conf := "etc/conf/mysql_exporter.yml"
	options, err := config.ParseConfig(conf)
	if err != nil {
		log.Fatal(err)
	}

	// 日志
	if closes, err := initLogger(options.Logger); err != nil {
		log.Fatal(err)
	} else {
		logrus.Debug("read config success: ", conf)
		logrus.Debug("init log success")

		for _, c := range closes {
			defer c.Close()
		}
	}

	//数据库
	mysql_addr := fmt.Sprintf("%s:%d", options.MySQL.Host, options.MySQL.Port)
	db, err := initMySQL(options.MySQL)
	if err != nil {
		logrus.Fatal(err)
	}
	if err := db.Ping(); err != nil {
		logrus.Error("failed to connect to mysql: ", mysql_addr)
		logrus.Error(err)
	} else {
		logrus.Debug("connected to mysql success: ", mysql_addr)
	}

	// Metrics
	initMetrics(options, db)

	// 落地页
	http.Handle("/", handler.Auth(
		handler.IndexHandler(),
		options.Web.Auth,
	))

	// 暴露http api
	http.Handle(options.Web.Addr.Path, handler.Auth(
		promhttp.Handler(),
		options.Web.Auth,
	))

	// 启动web服务
	httpAddr := fmt.Sprintf("%s:%d", options.Web.Addr.Host, options.Web.Addr.Port)

	if options.Web.SSL.CertFile == "" {
		logrus.Debug(fmt.Sprintf(" * Running on http://%s (Press CTRL+C to quit)\n", httpAddr))
		logrus.Fatal(http.ListenAndServe(httpAddr, nil))
	} else {
		logrus.Debug(fmt.Sprintf(" * Running on https://%s (Press CTRL+C to quit)\n", httpAddr))
		logrus.Fatal(http.ListenAndServeTLS(httpAddr, options.Web.SSL.CertFile, options.Web.SSL.KeyFile, nil))
	}
}
