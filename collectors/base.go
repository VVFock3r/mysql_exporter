package collectors

import (
	"database/sql"
	"github.com/sirupsen/logrus"
)

type mysqlCollector struct {
	db *sql.DB
}

func (c *mysqlCollector) status(name string) float64 {
	var (
		vname string
		value float64
	)

	// 如果发生错误，会在查一次
	// 比如数据库发生重启，会产生无效的连接
	for i := 1; i <= 2; i++ {
		sql := `show global status where variable_name=?`
		err := c.db.QueryRow(sql, name).Scan(&vname, &value)
		if err != nil {
			logrus.Info("query mysql status ", name, " error: ", err)
		} else {
			break
		}
	}

	return value
}

func (c *mysqlCollector) variables(name string) float64 {
	var (
		vname string
		value float64
	)

	sql := `show global variables where variable_name=?`
	err := c.db.QueryRow(sql, name).Scan(&vname, &value)
	if err != nil {
		logrus.Info("query mysql variables error: ", err)
		return 0
	}
	return value
}
