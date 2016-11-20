package g

import (
	"fmt"
	"os"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/cache"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	log "github.com/ninefive/nf-goutils/logtool"
)

var Cache cache.Cache
var blogCacheExpire int64
var catalogCacheExpire int64
var RunMode string
var Cfg = beego.AppConfig

func InitEnv() {
	var err error

	// log
	loglevel := Cfg.String("log_level")
	log.SetLevelWithDefault(loglevel, "info")

	// cache
	Cache, err := cache.NewCache("memory", `{"interval":60}`)
	if err != nil {
		log.Fetal("cache init fail :(")
		os.Exit(1)
	}
	blogCacheExpire, _ = Cfg.Int64("blog_cache_expire")
	catalogCacheExpire, _ = Cfg.Int64("catalog_cache_expire")

	// database
	dbUser := Cfg.String("db_user")
	dbPass := Cfg.String("db_pass")
	dbHost := Cfg.String("db_host")
	dbPort := Cfg.String("db_port")
	dbName := Cfg.String("db_name")
	maxIdleConn, _ := Cfg.Int("db_max_idle_conn")
	maxOpenConn, _ := Cfg.Int("db_max_open_conn")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8", dbUser, dbPass, dbHost, dbPort, dbName)

	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", dsn, maxIdleConn, maxOpenConn)

	RunMode = Cfg.String("runmode")
	if RunMode == "dev" {
		orm.Debug = true
	}
	initCfg()
}
