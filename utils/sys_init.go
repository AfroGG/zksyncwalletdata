package utils

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/go-redis/redis"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func InitConfig() {
	viper.SetConfigName("web")
	viper.AddConfigPath("conf")
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Printf("配置错误err = %v", err)
	}
	fmt.Printf("web inited")
	fmt.Printf("mysql inited")
}

var (
	DB  *gorm.DB
	Red *redis.Client
)

func InitMysql() {
	newLogger := logger.New(
		log.New(os.Stdout, "\n\r", log.LstdFlags),
		logger.Config{
			SlowThreshold: time.Second, //慢sql阈值
			LogLevel:      logger.Info, //级别
			Colorful:      true,        //彩色
		},
	)
	DB, _ = gorm.Open(mysql.Open(viper.GetString("mysql.dns")), &gorm.Config{})

	fmt.Println(&gorm.Config{Logger: newLogger})

	//user := models.User{}
	//DB.Find(&user)
	//fmt.Println(user)
}

func InitRedis() {
	Red = redis.NewClient(&redis.Options{
		Addr:         viper.GetString("redis.addr"),
		Password:     viper.GetString("redis.password"),
		DB:           viper.GetInt("redis.DB"),
		PoolSize:     viper.GetInt("redis.poolSize"),
		MinIdleConns: viper.GetInt("redis.minIdleConn"),
	})
	pong, err := Red.Ping().Result()
	if err != nil {
		fmt.Println("init redis err", err)
	} else {
		fmt.Println(pong)
	}

}
