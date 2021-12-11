package config

type MysqlCfg struct {
	Host     string `mapstructure:"host"`
	Port     int    `mapstructrue:"port"`
	Username string `mapstructure:"username"`
	Password string `mapstructure:"password"`
	Database string `mapstructure:"database"`
}
