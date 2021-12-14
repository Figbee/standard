package initialize

import (
	"fmt"
	"os"
	"standard/internal/global"
	"testing"

	"github.com/gobuffalo/packr/v2"
	"github.com/spf13/viper"
)

func TestInitConfig(t *testing.T) {
	//1.packr 打包 2.viper.new 3.viper.setConfigType 设置类型 4 box.find  查找打包的文件内容(默认配置)
	//5. v.readconfig 读取box.find的配置内容 //将
	global.ConfBox = packr.New("mybox", configPath)
	//获取实例(可创建多实例读取多个配置文件)
	v := viper.New()
	v.SetConfigType(configType)
	readConfig(v, developmentConfig)
	settings := v.AllSettings()
	for key, setting := range settings {
		v.SetDefault(key, setting)
	}
	//读取当前go运行环境变量
	env := os.Getenv("GO_ENV")
	configName := ""
	if env == "staging" {
		configName = stagingConfig
	} else if env == "production" {
		configName = productionConfig
	}
	if configName != "" {
		readConfig(v, configName)
	}

	//viper.unmarshal将viper解构到一个结构体体
	if err := v.Unmarshal(&global.Conf); err != nil {
		panic(fmt.Sprintf("初始化配置文件失败:%v", err))
	}
	t.Log("初始化配置文件完成")
	t.Logf("struct is %+v", global.Conf)
}
