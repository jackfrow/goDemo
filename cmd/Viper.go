package main

import (
	"fmt"
	"github.com/spf13/viper"
	"log"
)

func main() {
	// 设置 Viper 读取配置文件
	viper.SetConfigName("config") // 配置文件的名称（不包含扩展名）
	viper.AddConfigPath(".")      // 配置文件所在的路径（相对路径）
	viper.SetConfigType("json")   // 配置文件类型（json）

	// 读取配置文件
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file, %s", err)
	}

	// 获取配置值
	appName := viper.GetString("app_name")
	version := viper.GetString("version")
	debug := viper.GetBool("debug")
	port := viper.GetInt("port")

	// 打印配置
	fmt.Printf("App Name: %s\n", appName)
	fmt.Printf("Version: %s\n", version)
	fmt.Printf("Debug: %v\n", debug)
	fmt.Printf("Port: %d\n", port)
}
