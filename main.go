package main

import (
	"EdgeSys/apps/business"
	"EdgeSys/pkg/cache"
	"EdgeSys/pkg/config"
	"EdgeSys/pkg/global"
	"EdgeSys/pkg/initialize"
	"EdgeSys/pkg/middleware"
	"context"
	"github.com/spf13/cobra"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	gormlog "gorm.io/gorm/logger"
	"log"
	"mod.miligc.com/edge-common/CommonKit/logger"
	"mod.miligc.com/edge-common/CommonKit/rediscli"
	"mod.miligc.com/edge-common/CommonKit/restfulx"
	"mod.miligc.com/edge-common/CommonKit/starter"
	public "mod.miligc.com/edge-common/edgesys-common/pkg"
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
			public.Log = logger.InitLog(global.Conf.Log.File.GetFilename(), global.Conf.Log.Level)
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
			public.Log.Infof("%s连接成功", global.Conf.Server.DbType)
			//初始化业务系统数据库
			initBusinessDb()
			client, err := rediscli.NewRedisClient(global.Conf.Redis.Host, global.Conf.Redis.Password, global.Conf.Redis.Port, global.Conf.Redis.Db)
			if err != nil {
				public.Log.Panic("Redis连接错误")
			} else {
				public.Log.Info("Redis连接成功")
			}
			cache.RedisDb = client
			initialize.InitTable()
			// 初始化事件监听
			go initialize.InitEvents()
		} else {
			public.Log.Panic("请配置config")
		}
	},
	Run: func(cmd *cobra.Command, args []string) {
		// 前置 函数
		restfulx.UseBeforeHandlerInterceptor(middleware.PermissionHandler)
		// 后置 函数
		restfulx.UseAfterHandlerInterceptor(middleware.LogHandler)
		restfulx.UseAfterHandlerInterceptor(middleware.OperationHandler)

		app := initialize.InitRouter()
		if global.Conf.Server.LoadMethod == "debug" {
			business.InitModule{}.InitPluginDebug(app.Container)
		} else {
			business.InitModule{}.InitPlugin(app.Container)
		}
		public.Log.Info("路由初始化完成")
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

func initBusinessDb() {
	//连接业务系统数据库
	mysqlConfig := mysql.Config{
		DSN:                       global.Conf.Mysql.Msn(), // DSN data source name
		DefaultStringSize:         191,                     // string 类型字段的默认长度
		DisableDatetimePrecision:  true,                    // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
		DontSupportRenameIndex:    true,                    // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
		DontSupportRenameColumn:   true,                    // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
		SkipInitializeWithVersion: false,                   // 根据版本自动配置
	}
	ormConfig := &gorm.Config{Logger: gormlog.Default.LogMode(gormlog.Info)}
	milidbGorm, _ := gorm.Open(mysql.New(mysqlConfig), ormConfig)
	public.Log.Infof("连接mysql [%s]", global.Conf.Mysql.Msn())
	global.MiliDb = milidbGorm
	public.Log.Infof("%s连接成功", global.Conf.Server.DbType)

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
