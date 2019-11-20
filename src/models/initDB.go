package models

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/astaxie/beego/logs"
	"github.com/go-ini/ini"
	_ "github.com/lib/pq"
)

var db *sql.DB

func InitDB() error {
	cfg := LoadConf("conf.ini")
	host := cfg.Section("db").Key("host").MustString("localhost")
	port := cfg.Section("db").Key("port").MustInt(5432)
	user := cfg.Section("db").Key("user").MustString("postgres")
	password := cfg.Section("db").Key("password").MustString("postgres")
	dbname := cfg.Section("db").Key("dbname").MustString("postgres")

	dbInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	var err error
	db, err = sql.Open("postgres", dbInfo)
	if err != nil {
		logs.Error(err)
		return err
	}
	err = db.Ping()
	if err != nil {
		logs.Error(err)
		return err
	}
	db.SetMaxIdleConns(20)
	db.SetMaxOpenConns(40)
	logs.Info("DB connected!")
	return nil
}

func LoadConf(confPath string) *ini.File {
	cfg, err := ini.Load(confPath)
	if err != nil {
		logs.Error(err)
		os.Exit(1)
	}
	return cfg
}
