package mysql

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

var db *sqlx.DB

func Init() (err error) {
	dsn := fmt.Sprintf("%s:%d@tcp(%s:%d)/%s",
		viper.GetString("mysql.username"),
		viper.GetInt("mysql.password"),
		viper.GetString("mysql.host"),
		viper.GetInt("mysql.port"),
		viper.GetString("mysql.dbname"),
	)
	db, err = sqlx.Connect("mysql", dsn)
	if err != nil {
		zap.L().Fatal("连接数据库失败", zap.Error(err))
		return
	}
	db.SetMaxOpenConns(viper.GetInt("mysql.SetMaxOpenConns"))
	db.SetMaxIdleConns(viper.GetInt("mysql.SetMaxIdleConns"))
	return
}
