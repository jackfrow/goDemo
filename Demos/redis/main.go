package main

import (
	"fmt"
	"github.com/go-redis/redis"
)

func main() {
	// 创建一个新的 Redis 客户端
	client := redis.NewClient(&redis.Options{
		Addr: "localhost:6379", // Redis 服务器地址
	})
	// 检查是否成功连接到 Redis
	_, err := client.Ping().Result()
	if err != nil {
		fmt.Println("无法连接到 Redis:", err)
		return
	}

	// 设置键值对
	err = client.Set("example_key", "example_value", 0).Err()
	if err != nil {
		fmt.Println("设置键值对时出错:", err)
		return
	}

	// 获取值
	val, err := client.Get("example_key").Result()
	if err != nil {
		fmt.Println("获取值时出错:", err)
		return
	}
	fmt.Println("example_key 的值为:", val)

	// 删除键
	err = client.Del("example_key").Err()
	if err != nil {
		fmt.Println("删除键时出错:", err)
		return
	}
	fmt.Println("example_key 已成功删除")
}
