package initialize

import (
	"fmt"
	"gin-blog-server/global"
	"log"
	"os"

	"github.com/fsnotify/fsnotify"

	"github.com/spf13/viper"
)

// InitViper 初始化 Viper 读取配置文件
func InitViper() *viper.Viper {
	v := viper.New()
	// Get APP_ENV, default dev
	env := os.Getenv("APP_ENV")
	if env == "" {
		env = "dev"
	}
	cfgName := fmt.Sprintf("config.%s", env)
	v.SetConfigName(cfgName)
	v.SetConfigType("yaml")
	v.AddConfigPath("./config")

	if err := v.ReadInConfig(); err != nil {
		log.Fatal(fmt.Sprintf("Read config: %s failed, %s", cfgName, err.Error()))
	}

	// Unmarshal config
	if err := v.Unmarshal(&global.AppCfg); err != nil {
		log.Fatal("Unmarshal config failed: ", err.Error())
	}

	// Watching and re-reading
	v.WatchConfig()
	v.OnConfigChange(func(e fsnotify.Event) {
		log.Printf("Config file: %s changed, Operation: %s", e.Name, e.Op)
		// re-reading
		if err := v.Unmarshal(&global.AppCfg); err != nil {
			log.Print("Reload config failed")
			return
		}
		log.Print("Reloaded config")
	})

	return v
}
