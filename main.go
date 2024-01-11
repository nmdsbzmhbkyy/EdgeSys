package main

import (
	"EdgeSys/pkg/cache"
	"EdgeSys/pkg/config"
	"EdgeSys/pkg/global"
	"EdgeSys/pkg/initialize"
	"EdgeSys/pkg/middleware"
	"context"
	"github.com/spf13/cobra"
	"log"
	"mod.miligc.com/edge-common/CommonKit/logger"
	"mod.miligc.com/edge-common/CommonKit/rediscli"
	"mod.miligc.com/edge-common/CommonKit/restfulx"
	"mod.miligc.com/edge-common/CommonKit/starter"
	"os"
	"os/signal"
	"syscall"
)

var (
	configFile string
)

var rootCmd = &cobra.Command{
	Use:   "EdgeSys is the main component in the edge common lib.",
	Short: `EdgeSys is go go-restful frame`,
	PreRun: func(cmd *cobra.Command, args []string) {
		if configFile != "" {
			global.Conf = config.InitConfig(configFile)
			global.Log = logger.InitLog(global.Conf.Log.File.GetFilename(), global.Conf.Log.Level)
			dbGorm := starter.DbGorm{Type: global.Conf.Server.DbType}
			if global.Conf.Server.DbType == "mysql" {
				dbGorm.Dsn = global.Conf.Mysql.Dsn()
				dbGorm.MaxIdleConns = global.Conf.Mysql.MaxIdleConns
				dbGorm.MaxOpenConns = global.Conf.Mysql.MaxOpenConns
			} else {
				dbGorm.Dsn = global.Conf.Postgresql.PgDsn()
				dbGorm.MaxIdleConns = global.Conf.Postgresql.MaxIdleConns
				dbGorm.MaxOpenConns = global.Conf.Postgresql.MaxOpenConns
			}
			dbGorm.DBLog = global.Conf.Log.DBLog
			global.Db = dbGorm.GormInit()
			global.Log.Infof("%s连接成功", global.Conf.Server.DbType)
			client, err := rediscli.NewRedisClient(global.Conf.Redis.Host, global.Conf.Redis.Password, global.Conf.Redis.Port, global.Conf.Redis.Db)
			if err != nil {
				global.Log.Panic("Redis连接错误")
			} else {
				global.Log.Info("Redis连接成功")
			}
			cache.RedisDb = client
			initialize.InitTable()
			// 初始化事件监听
			go initialize.InitEvents()
		} else {
			global.Log.Panic("请配置config")
		}
	},
	Run: func(cmd *cobra.Command, args []string) {
		defer func() {
			if err := recover(); err != nil {
				global.Log.Error("recover 异常捕获:", err)
			}
		}()
		// 前置 函数
		restfulx.UseBeforeHandlerInterceptor(middleware.PermissionHandler)
		// 后置 函数
		restfulx.UseAfterHandlerInterceptor(middleware.LogHandler)
		restfulx.UseAfterHandlerInterceptor(middleware.OperationHandler)

		app := initialize.InitRouter()
		global.Log.Info("路由初始化完成")
		app.Start(context.TODO())

		stop := make(chan os.Signal, 1)
		signal.Notify(stop, syscall.SIGTERM, os.Interrupt)
		<-stop
		if err := app.Stop(context.TODO()); err != nil {
			log.Fatalf("fatal app stop: %s", err)
			os.Exit(-3)
		}
	},
}

func init() {
	rootCmd.Flags().StringVar(&configFile, "config", getEnvStr("PANDA_CONFIG", "./config.yml"), "edge config file path.")
}

func getEnvStr(env string, defaultValue string) string {
	v := os.Getenv(env)
	if v == "" {
		return defaultValue
	}
	return v
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		rootCmd.PrintErrf("root cmd execute: %s", err)
		os.Exit(1)
	}
}
