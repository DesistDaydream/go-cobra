package config

import (
	"fmt"
	"os"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type Config struct {
	Username string
	Password string
}

func NewConfig(configFile string) *Config {
	var config Config
	if configFile != "" {
		fmt.Println("使用配置文件: ", configFile)
		viper.SetConfigFile(configFile)
	} else {
		// 设置一下 Viper 读取配置文件时的搜索路径、文件名、文件类型等信息
		viper.AddConfigPath("./config") // 首先在当前目录查找配置
		viper.SetConfigName("config")   // 配置文件名 (不带后缀)
		viper.SetConfigType("yaml")     // 如果配置文件的名称中没有扩展名，则需要配置此项
	}

	// 上述配置会让 Viper 从 ./config/ 目录中搜索 config.yaml 文件
	// 搜索到文件后，读取配置文件
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println("读取配置文件失败: ", err)
		os.Exit(1)
	}
	logrus.Info("配置文件绝对路径: ", viper.ConfigFileUsed())

	// 解析配置文件
	err = viper.Unmarshal(&config)
	if err != nil {
		fmt.Println("解析配置文件失败: ", err)
		os.Exit(1)
	}
	return &config
}
