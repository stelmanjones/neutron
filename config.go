package neutron

import (
	"github.com/spf13/viper"
)

type NeutronConfig struct {
	Name   string `yaml:"name"`
	Width  int    `yaml:"width"`
	Height int    `yaml:"height"`
	Debug  bool   `yaml:"debug"`
}

func InitConfig() *NeutronConfig {
	c := &NeutronConfig{}
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./neutron.yaml")
	viper.SetConfigName("neutron")
	viper.SetDefault("name", "Neutron")
	viper.SetDefault("debug", false)
	viper.SetDefault("width", 800)
	viper.SetDefault("height", 600)

	viper.ReadInConfig()
	viper.Unmarshal(c)
	return c
}
