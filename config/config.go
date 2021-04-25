package config

import "gopkg.in/ini.v1"
import "fmt"
import "os"

type Config struct {
	AppPort   string
	RedisUrl  string
	UserConns string
	IPConns   string
	KeyPrefix string
	DB        int
	Topic     string
	Topic2    string
	KafkaUrl  string
	LogPath string

	RedisWriteUrl  string
	RedisWriteDB  int
	RedisWritePrefix  string

}

var AppConfig = &Config{}

func init() {
	//init config
	cfg, err := ini.Load("app.ini")
	if err != nil {
		fmt.Printf("Fail to read file: %v", err)
		os.Exit(1)
	}

	port := cfg.Section("").Key("app_port").String()
	logPath := cfg.Section("").Key("log_path").String()
	redisUrl := cfg.Section("redis").Key("url").String()
	userConns := cfg.Section("app").Key("userconns").String()
	ipConns := cfg.Section("app").Key("ipconns").String()
	keyPrefix := cfg.Section("redis").Key("prefix").String()
	redisDB, _ := cfg.Section("redis").Key("db").Int()
	kafkaTopic := cfg.Section("kafka").Key("topic").String()
	kafkaUrl := cfg.Section("kafka").Key("url").String()
	redisWriteDB, _ := cfg.Section("rediswrite").Key("db").Int()
	redisWritePrefix := cfg.Section("rediswrite").Key("prefix").String()
	redisWriteUrl := cfg.Section("rediswrite").Key("url").String()
	topic2 := cfg.Section("kafka").Key("topic2").String()

	AppConfig.RedisWriteUrl  = redisWriteUrl
	AppConfig.RedisWriteDB   = redisWriteDB
	AppConfig.RedisWritePrefix = redisWritePrefix
	AppConfig.AppPort = port
	AppConfig.RedisUrl = redisUrl
	AppConfig.UserConns = userConns
	AppConfig.IPConns = ipConns
	AppConfig.KeyPrefix = keyPrefix
	AppConfig.DB = redisDB
	AppConfig.Topic = kafkaTopic
	AppConfig.KafkaUrl = kafkaUrl
	AppConfig.LogPath = logPath
	AppConfig.Topic2 = topic2
}
