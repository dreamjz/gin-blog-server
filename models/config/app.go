package config

type AppCfg struct {
	Mysql  MysqlCfg  `mapstructure:"mysql"`
	Server ServerCfg `mapstructure:"server"`
}
