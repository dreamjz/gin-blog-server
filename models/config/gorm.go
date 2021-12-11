package config

type GormCfg struct {
	TablePrefix  string `mapstructure:"tablePrefix"`
	MaxIdleConns int    `mapstructure:"maxIdleConns"`
	MaxOpenConns int    `mapstructure:"maxOpenConns"`
}
