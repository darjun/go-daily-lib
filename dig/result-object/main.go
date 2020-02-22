package main

import (
	"fmt"
	"log"

	"github.com/jessevdk/go-flags"
	"go.uber.org/dig"
	"gopkg.in/ini.v1"
)

type Option struct {
	ConfigFile string `short:"c" long:"config" description:"Name of config file."`
}

type RedisConfig struct {
	IP   string
	Port int
	DB   int
}

type MySQLConfig struct {
	IP       string
	Port     int
	User     string
	Password string
	Database string
}

type Config struct {
	dig.Out

	Redis *RedisConfig
	MySQL *MySQLConfig
}

func InitOption() (*Option, error) {
	var opt Option
	_, err := flags.Parse(&opt)

	return &opt, err
}

func InitConfig(opt *Option) (*ini.File, error) {
	cfg, err := ini.Load(opt.ConfigFile)
	return cfg, err
}

func InitRedisConfig(cfg *ini.File) (*RedisConfig, error) {
	port, err := cfg.Section("redis").Key("port").Int()
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	db, err := cfg.Section("redis").Key("db").Int()
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	return &RedisConfig{
		IP:   cfg.Section("redis").Key("ip").String(),
		Port: port,
		DB:   db,
	}, nil
}

func InitMySQLConfig(cfg *ini.File) (*MySQLConfig, error) {
	port, err := cfg.Section("mysql").Key("port").Int()
	if err != nil {
		return nil, err
	}

	return &MySQLConfig{
		IP:       cfg.Section("mysql").Key("ip").String(),
		Port:     port,
		User:     cfg.Section("mysql").Key("user").String(),
		Password: cfg.Section("mysql").Key("password").String(),
		Database: cfg.Section("mysql").Key("database").String(),
	}, nil
}

func InitRedisAndMySQLConfig(cfg *ini.File) (Config, error) {
	var config Config

	redis, err := InitRedisConfig(cfg)
	if err != nil {
		return config, err
	}

	mysql, err := InitMySQLConfig(cfg)
	if err != nil {
		return config, err
	}

	config.Redis = redis
	config.MySQL = mysql
	return config, nil
}

func PrintInfo(redis *RedisConfig, mysql *MySQLConfig) {
	fmt.Println("=========== redis section ===========")
	fmt.Println("redis ip:", redis.IP)
	fmt.Println("redis port:", redis.Port)
	fmt.Println("redis db:", redis.DB)

	fmt.Println("=========== mysql section ===========")
	fmt.Println("mysql ip:", mysql.IP)
	fmt.Println("mysql port:", mysql.Port)
	fmt.Println("mysql user:", mysql.User)
	fmt.Println("mysql password:", mysql.Password)
	fmt.Println("mysql db:", mysql.Database)
}

func main() {
	container := dig.New()

	container.Provide(InitOption)
	container.Provide(InitConfig)
	container.Provide(InitRedisAndMySQLConfig)

	err := container.Invoke(PrintInfo)
	if err != nil {
		log.Fatal(err)
	}
}
