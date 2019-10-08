package yaml

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

const fileType = "yaml"

type Config struct {
	Name string
}

func Init(cfg string) error {
	c := Config{
		Name: cfg,
	}
	// 初始化配置文件
	if err := c.initConfig(); err != nil {
		return err
	}
	c.watchConfig()
	return nil
}

func (c *Config) initConfig() error {
	if c.Name != "" {
		// 如果指定了配置文件，则解析指定的配置文件
		viper.SetConfigFile(c.Name)
	} else {
		// 如果没有指定配置文件，则解析默认的配置文件
		viper.AddConfigPath("conf")
		viper.SetConfigName("config")
	}
	// 设置配置文件格式为YAML
	viper.SetConfigType(fileType)
	// viper解析配置文件
	if err := viper.ReadInConfig(); err != nil {
		return err
	}
	return nil
}

// 监听配置文件是否改变,用于热更新
func (c *Config) watchConfig() {
	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		fmt.Printf("Config file changed: %s\n", e.Name)
		fmt.Printf("Config file changed operation: %s\n", e.Op)
	})
}

func Check(filename string) {
	if err := Init(filename); err != nil {
		panic(err)
	}
}

func Get(name string) string {
	v := viper.GetString(name)
	return v
}