package config

import (
	"fmt"
	"github.com/spf13/viper"
	"os"
)

func InitConfig() {
	wordDir, _ := os.Getwd()
	fmt.Println(wordDir)          //获取当前工作目录的路径
	viper.SetConfigName("config") //这一行代码设置了配置文件的名称为 "yml"，这意味着 Viper 将尝试查找名为 "yml" 的配置文件 ////这一行代码设置了配置文件的搜索路径。
	viper.SetConfigType("yml")
	viper.AddConfigPath(wordDir + "/../config")
	err := viper.ReadInConfig() //用于从配置文件中读取配置信息并加载到 Viper 实例中。这个函数通常在使用 Viper 时非常有用，
	if err != nil {
		fmt.Println("Failed to read config.yml file:", err)
		// 可以添加其他日志信息以便调试
		panic(err)
	}
	//fmt.Println("配置文件搜索路径:", viper.GetViper().ConfigFileUsed())
}
