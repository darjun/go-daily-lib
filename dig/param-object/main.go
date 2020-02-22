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
	dig.In

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
		return nil, err
	}

	db, err := cfg.Section("redis").Key("db").Int()
	if err != nil {
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

func PrintInfo(config Config) {
	fmt.Println("=========== redis section ===========")
	fmt.Println("redis ip:", config.Redis.IP)
	fmt.Println("redis port:", config.Redis.Port)
	fmt.Println("redis db:", config.Redis.DB)

	fmt.Println("=========== mysql section ===========")
	fmt.Println("mysql ip:", config.MySQL.IP)
	fmt.Println("mysql port:", config.MySQL.Port)
	fmt.Println("mysql user:", config.MySQL.User)
	fmt.Println("mysql password:", config.MySQL.Password)
	fmt.Println("mysql db:", config.MySQL.Database)
}

func main() {
	container := dig.New()

	container.Provide(InitOption)
	container.Provide(InitConfig)
	container.Provide(InitRedisConfig)
	container.Provide(InitMySQLConfig)

	err := container.Invoke(PrintInfo)
	if err != nil {
		log.Fatal(err)
	}
}
