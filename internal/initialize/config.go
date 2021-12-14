package initialize

import (
	"bytes"
	"fmt"
	"os"
	"standard/internal/global"

	"github.com/gobuffalo/packr/v2"

	"github.com/spf13/viper"
)

const (
	configType        = "yml"
	configPath        = "../../conf"
	developmentConfig = "config.dev.yml"
	productionConfig  = "config.prod.yml"
	stagingConfig     = "config.stage.yml"
)

func InitConfig() {
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
	fmt.Println("初始化配置文件完成")
}

func readConfig(v *viper.Viper, configFile string) {
	conf, err := global.ConfBox.Find(configFile)
	if err != nil {
		panic(fmt.Sprintf("初始化配置文件失败:%v", err))
	}
	if err = v.ReadConfig(bytes.NewReader(conf)); err != nil {
		panic(fmt.Sprintf("初始化配置文件失败:%v", err))
	}
}
