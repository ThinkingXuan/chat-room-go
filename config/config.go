package config

import (
	"github.com/fsnotify/fsnotify"
	"github.com/golang/glog"
	"github.com/spf13/viper"
)

// Config struct
type Config struct {
	Path string
}

// InitConfig init config
func InitConfig(cfgPath string) error {
	c := Config{
		Path: cfgPath,
	}

	if err := c.setConfig(); err != nil {
		return err
	}

	//c.watchConfig()
	return nil
}

// setConfig set config params
func (c *Config) setConfig() error {
	if c.Path != "" {
		viper.SetConfigFile(c.Path)
	} else {
		viper.AddConfigPath("server/conf")
		viper.SetConfigName("config")
	}
	viper.SetConfigType("yaml")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		return err
	}
	return nil
}

// watchConifg watch conifg file change
func (c *Config) watchConfig() {
	viper.WatchConfig()
	viper.OnConfigChange(func(in fsnotify.Event) {
		glog.Infof("Config file changed: %s", in.Name)
	})
}
