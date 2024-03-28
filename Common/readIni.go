package Common

import (
	"fmt"

	"github.com/spf13/viper"
)

// viper 常用读取配置的方法
//Get(key string) : interface{}
//GetBool(key string) : bool
//GetFloat64(key string) : float64
//GetInt(key string) : int
//GetIntSlice(key string) : []int
//GetString(key string) : string
//GetStringMap(key string) : map[string]interface{}
//GetStringMapString(key string) : map[string]string
//GetStringSlice(key string) : []string
//GetTime(key string) : time.Time
//GetDuration(key string) : time.Duration
//IsSet(key string) : bool
//AllSettings() : map[string]interface{}

// ReadLin name 读取配置文件名
func ReadLin(name string) *viper.Viper {
	config := viper.New()
	config.AddConfigPath("./ini/") // 文件所在目录
	config.SetConfigName(name)     // 文件名
	config.SetConfigType("ini")    // 文件类型

	if err := config.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			fmt.Println("找不到配置文件..")
		} else {
			fmt.Println("配置文件出错..")
		}
	}
	return config
}

// ReadLinGet 返回interfaces
func ReadLinGet(name string, readName string) interface{} {
	config := ReadLin(name)
	host := config.Get(readName) // 读取配置
	return host

}

// ReadLinInt 返回int
func ReadLinInt(name string, readName string) int {
	config := ReadLin(name)
	host := config.GetInt(readName) // 读取配置
	return host
}

// ReadLinBool  返回Bool
func ReadLinBool(name string, readName string) bool {
	config := ReadLin(name)
	host := config.GetBool(readName) // 读取配置
	return host
}

// ReadLinString 返回string
func ReadLinString(name string, readName string) string {
	config := ReadLin(name)
	host := config.GetString(readName) // 读取配置
	return host
}
