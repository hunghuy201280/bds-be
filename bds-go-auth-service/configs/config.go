package configs

import (
	"bds-go-auth-service/common/l"
	"github.com/spf13/viper"
	"strings"
)

var (
	ll = l.New()
)

type Config struct {
	Port int `yaml:"port"`
}

func Load() *Config {
	var cfg = &Config{}

	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./configs/")

	err := viper.ReadInConfig()
	if err != nil {
		ll.Fatal("Failed to read viper config", l.Error(err))
	}

	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "__"))
	viper.AutomaticEnv()

	err = viper.Unmarshal(cfg)
	if err != nil {
		ll.Fatal("Failed to unmarshal config", l.Error(err))
	}

	ll.Info("Config loaded", l.Object("config", cfg))
	return cfg
}
